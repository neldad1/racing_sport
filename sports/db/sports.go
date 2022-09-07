package db

import (
	"database/sql"
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
	EventsList(filter *sports.ListEventsRequestFilter) ([]*sports.Event, error)
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

// EventsList will return a list of sport events.
func (s *sportsRepo) EventsList(filter *sports.ListEventsRequestFilter) ([]*sports.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getEventsQueries()[eventsList]

	query, args = s.applyFilter(query, filter)

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
