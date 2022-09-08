package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"git.neds.sh/matty/entain/sports/proto/sports"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type sportsRepo struct {
	db   *sql.DB
	init sync.Once
}

// SportsRepo provides repository access to sports.
type SportsRepo interface {
	// Init will initialise our sports repository.
	Init() error

	// EventsList will return a list of sport events.
	EventsList(filter *sports.ListEventsRequestFilter, orderBy string) ([]*sports.Event, error)
	// GetEventById will return the information of a given event id.
	GetEventByID(id int64) (*sports.Event, error)
}

// NewSportsRepo creates a new sport repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

// Init prepares the sports repository dummy data.
func (s *sportsRepo) Init() error {
	var err error

	s.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy sport events.
		err = s.seed()
	})

	return err
}

// GetEventByID will return the information of a given event id.
func (s *sportsRepo) GetEventByID(id int64) (*sports.Event, error) {
	var event sports.Event
	var advertisedStart, advertisedEnd time.Time

	row := s.db.QueryRow(`SELECT id, 
	name, 
	venue_id, 
	sport_id, 
	participants_id, 
	advertised_start_time,
	advertised_end_time
	FROM events where id=?`, id)

	if err := row.Scan(&event.Id, &event.Name, &event.VenueId, &event.SportId, &event.ParticipantsId, &advertisedStart, &advertisedEnd); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("event not found")
		}

		return nil, err
	}

	// Convert time to proto timestamp.
	event.AdvertisedStartTime = timestamppb.New(advertisedStart)
	event.AdvertisedEndTime = timestamppb.New(advertisedEnd)

	event.Status = getEventStatus(advertisedStart, advertisedEnd)

	return &event, nil
}

// EventsList will return a list of sport events.
func (s *sportsRepo) EventsList(filter *sports.ListEventsRequestFilter, orderBy string) ([]*sports.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getEventsQueries()[eventsList]

	query, args = s.applyFilter(query, filter)

	query, err = s.applyOrderBy(query, orderBy)
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return s.scanEvents(rows)
}

// applyFilter builds the requested filter for sports events.
func (s *sportsRepo) applyFilter(query string, filter *sports.ListEventsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	// Add filter for sport_id
	if filter.SportId != nil {
		clauses = append(clauses, "sport_id=?")
		args = append(args, *filter.SportId)
	}

	if len(filter.Status) > 0 {
		status := strings.TrimSpace(filter.Status)
		status = strings.ToUpper(status)
		now := time.Now().Format(time.RFC3339)

		if status == "CLOSED" {
			clauses = append(clauses, "advertised_end_time < ?")
			args = append(args, now)
		} else if status == "OPEN" {
			clauses = append(clauses, "advertised_start_time > ?")
			args = append(args, now)
		} else if status == "ONGOING" {
			clauses = append(clauses, "advertised_start_time < ?")
			args = append(args, now)

			clauses = append(clauses, "advertised_end_time > ?")
			args = append(args, now)
		}
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

// applyOrderBy adds the ORDER BY clause in the query based on the requested sorting option.
func (s *sportsRepo) applyOrderBy(query string, orderBy string) (string, error) {
	// Return immediately when orderBy is empty.
	if len(orderBy) == 0 {
		return query, nil
	}

	orderBy = strings.TrimSpace(orderBy)

	// TODO: Implement sorting using status as it is not a column in the event table.
	if orderBy == "status" {
		return query, nil
	}

	// Check for desc/asc option.
	params := strings.Split(orderBy, " ")
	paramsLength := len(params)
	fieldName := params[0]

	// Return invalid value error if the length of params is more than 2. Field name and asc/desc are the only values accepted.
	if paramsLength > 2 {
		return query, fmt.Errorf("invalid orderBy value: %s Format is `fieldName desc`", orderBy)
	}

	// Get the corresponding column name.
	clause := getEventColumnName(fieldName)

	if len(clause) == 0 {
		return query, fmt.Errorf("unable to find the field name: %s", fieldName)
	}

	// Check if asc or desc is provided.
	if paramsLength == 2 {
		order := strings.ToUpper(params[1])
		if order != "ASC" && order != "DESC" {
			return query, fmt.Errorf("invalid sort order: %s. Choose either `asc` or `desc`", params[1])
		}
		clause += " " + order
	}

	query += " ORDER BY " + clause

	return query, nil
}

// scanEvents copies the data from the database into the values of each event.
func (s *sportsRepo) scanEvents(rows *sql.Rows) ([]*sports.Event, error) {
	var events []*sports.Event

	for rows.Next() {
		var event sports.Event
		var advertisedStart, advertisedEnd time.Time

		if err := rows.Scan(&event.Id, &event.Name, &event.VenueId, &event.SportId, &event.ParticipantsId, &advertisedStart, &advertisedEnd); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		// Convert time to proto timestamp.
		event.AdvertisedStartTime = timestamppb.New(advertisedStart)
		event.AdvertisedEndTime = timestamppb.New(advertisedEnd)

		event.Status = getEventStatus(advertisedStart, advertisedEnd)

		events = append(events, &event)
	}

	return events, nil
}

// getEventStatus gets the correct status of the event: OPEN for future event, ONGOING for event that has started and CLOSED for past event.
func getEventStatus(advertisedStart, advertisedEnd time.Time) string {
	status := "CLOSED"

	now := time.Now()

	if advertisedStart.After(now) {
		status = "OPEN"
	} else if advertisedStart.Before(now) && advertisedEnd.After(now) {
		status = "ONGOING"
	}

	return status
}

// getEventColumnName checks if the value is equivalent to any of the column names in sports.Event.
func getEventColumnName(value string) string {
	var columnName string

	v := reflect.ValueOf(sports.Event{})

	for i := 0; i < v.NumField(); i++ {
		structField := v.Type().Field(i)
		protoBufTag := structField.Tag.Get("protobuf")
		if strings.Contains(protoBufTag, value) {
			values := strings.Split(structField.Tag.Get("json"), ",")
			columnName = values[0]
			break
		}
	}

	return columnName
}
