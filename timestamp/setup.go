package timestamp

import (
	"fmt"
	"time"
)

func NewTimestamp(params Config) (Timestamp, error) {
	local := "Africa/Freetown"
	format := "Monday, 02 January 2006, 03:04 PM"
	parse := time.RFC3339

	if params.Local != "" {
		local = params.Local
	}

	if params.Format != "" {
		format = params.Format
	}

	if params.Parse != "" {
		parse = params.Parse
	}

	location, err := time.LoadLocation(local)
	if err != nil {
		return Timestamp{}, fmt.Errorf("invalid default location: %w", err)
	}

	return Timestamp{location: location, format: format, parse: parse}, nil
}
