package db

import (
	"time"

	"syreclabs.com/go/faker"
)

// seed add dummy data to sports database.
func (s *sportsRepo) seed() error {
	return s.seedEvents()
}

// seedEvents add dummy data to events table in sports database.
func (s *sportsRepo) seedEvents() error {
	statement, err := s.db.Prepare(`CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY, name TEXT, venue_id INTEGER, sport_id INTEGER, participants_id INTEGER, advertised_start_time DATETIME, advertised_end_time DATETIME)`)
	if err != nil {
		return err
	}

	if _, err = statement.Exec(); err != nil {
		return err
	}

	for i := 1; i <= 100; i++ {
		statement, err = s.db.Prepare(`INSERT OR IGNORE INTO events(id, name, venue_id, sport_id, participants_id, advertised_start_time, advertised_end_time) VALUES (?,?,?,?,?,?,?)`)
		if err != nil {
			return err
		}
		_, err = statement.Exec(
			i,
			faker.Team().Name(),
			faker.Number().Between(1, 30),
			faker.Number().Between(1, 20),
			faker.Number().Between(1, 10),
			faker.Time().Between(time.Now().AddDate(0, 0, -5), time.Now().AddDate(0, 0, 2)).Format(time.RFC3339),
			faker.Time().Between(time.Now().AddDate(0, 0, 3), time.Now().AddDate(0, 0, 5)).Format(time.RFC3339),
		)
		if err != nil {
			return err
		}
	}

	return nil
}
