package db

import (
	"time"

	"syreclabs.com/go/faker"
)

func (r *racesRepo) seed() error {
	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS races (id INTEGER PRIMARY KEY, meeting_id INTEGER, name TEXT, number INTEGER, visible INTEGER, advertised_start_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO races(id, meeting_id, name, number, visible, advertised_start_time) VALUES (?,?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.Number().Between(1, 10),
				faker.Team().Name(),
				faker.Number().Between(1, 12),
				faker.Number().Between(0, 1),
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			)
		}
	}

	err = r.setFutureStartTime(100)

	return err
}

// Set advertised_start_time of the specified id to a future time value to show the status as OPEN.
func (r *racesRepo) setFutureStartTime(id int) error {
	var additionalHour time.Duration = (60 * time.Duration(id))
	futureTime := time.Now().Add(additionalHour + time.Hour)

	statement, err := r.db.Prepare(`UPDATE races SET advertised_start_time=? where id=?`)

	if err == nil {
		_, err = statement.Exec(futureTime, id)
	}
	return err
}
