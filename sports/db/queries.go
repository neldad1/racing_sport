package db

const (
	eventsList = "list"
)

func getEventsQueries() map[string]string {
	return map[string]string{
		eventsList: `
			SELECT 
				id,
				name, 
				venue_id, 
				sport_id,
				participants_id,
				advertised_start_time,
				advertised_end_time 
			FROM events
		`,
	}
}
