package neodb_go_client

import "time"

func parseDate(date string) (*time.Time, error) {
	// parse date
	if date == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02T15:04:05Z", date)
	if err != nil {
		return nil, err
	}
	return &t, err
}
