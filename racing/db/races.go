package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"

	"git.neds.sh/matty/entain/racing/proto/racing"
)

// RacesRepo provides repository access to races.
type RacesRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *racing.ListRacesRequestFilter, orderBy string) ([]*racing.Race, error)
}

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewRacesRepo(db *sql.DB) RacesRepo {
	return &racesRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *racesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

// List performs the requested parameters and returns the final list of races.
func (r *racesRepo) List(filter *racing.ListRacesRequestFilter, orderBy string) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getRaceQueries()[racesList]

	query, args = r.applyFilter(query, filter)

	query, err = r.applyOrderBy(query, orderBy)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

// applyFilter processes the filters included in the request.
func (r *racesRepo) applyFilter(query string, filter *racing.ListRacesRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	if filter.Visible != nil {
		// Handle the boolean implementation for SQLite3
		condition := "NOT visible"
		if *filter.Visible {
			condition = "visible"
		}
		clauses = append(clauses, condition)
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

// applyOrderBy adds the ORDER BY clause in the query based on the requested sorting option.
func (r *racesRepo) applyOrderBy(query string, orderBy string) (string, error) {
	// Return immediately when orderBy is empty.
	if len(orderBy) == 0 {
		return query, nil
	}

	orderBy = strings.TrimSpace(orderBy)

	// Check for desc/asc option.
	params := strings.Split(orderBy, " ")
	paramsLength := len(params)
	fieldName := params[0]

	// Return invalid value error if the length of params is more than 2. Field name and asc/desc are the only values accepted.
	if paramsLength > 2 {
		return query, fmt.Errorf("invalid orderBy value: %s Format is `fieldName desc`", orderBy)
	}

	// Get the corresponding column name.
	clause := getRaceColumnName(fieldName)

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

func (r *racesRepo) scanRaces(
	rows *sql.Rows,
) ([]*racing.Race, error) {
	var races []*racing.Race

	for rows.Next() {
		var race racing.Race
		var advertisedStart time.Time

		if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		race.AdvertisedStartTime = ts

		races = append(races, &race)
	}

	return races, nil
}

// getRaceColumnName checks if the value is equivalent to any of the column names in racing.Race.
func getRaceColumnName(value string) string {
	var columnName string

	v := reflect.ValueOf(racing.Race{})

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
