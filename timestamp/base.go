package timestamp

import (
	"time"
)

type (
	Timestamp struct {
		location      *time.Location
		format, parse string
	}
	Config struct {
		Local  string `mapstructure:"local"`
		Format string `mapstructure:"format"`
		Parse  string `mapstructure:"parse"`
	}
)

func (t Timestamp) Now() time.Time {
	now := time.Now()
	return now.In(t.location)
}

func (t Timestamp) Parse(date string) (time.Time, error) {
	now, err := time.Parse(t.parse, date)
	if err != nil {
		return time.Time{}, err
	}

	return now.In(t.location), nil
}

func (t Timestamp) Format(date time.Time) string {
	return date.In(t.location).Format(t.format)
}
