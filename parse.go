package w3cdtf

import (
	"fmt"
	"time"
)

// Parse parses a W3C-DTF formatted string and returns the time value it represents.
func Parse(s string) (time.Time, error) {
	for _, format := range []string{
		DateTimeNano,
		DateTime,
		DateHoursMinutes,
		Date,
		YearMonth,
		Year,
	} {
		t, err := time.Parse(format, s)
		if err != nil {
			continue
		}
		return t, nil
	}
	return time.Time{}, fmt.Errorf("invalid format")
}

// MustParse is like Parse but panics if the expression cannot be parsed.
func MustParse(s string) time.Time {
	t, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return t
}
