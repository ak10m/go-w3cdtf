package w3cdtf

import "time"

// Date and Time Formats
// See: https://www.w3.org/TR/NOTE-datetime
const (
	Year             = "2006"                      // Year
	YearMonth        = "2006-01"                   // Year and month
	Date             = "2006-01-02"                // Complete date
	DateHoursMinutes = "2006-01-02T15:04Z07:00"    // Complete date plus hours and minutes
	DateTime         = "2006-01-02T15:04:05Z07:00" // Complete date plus hours, minutes and seconds
	DateTimeNano     = time.RFC3339Nano            // Complete date plus hours, minutes, seconds and a decimal fraction of a second
)

// Format returns a textual representation of the time value formatted shortest W3C-DTF.
func Format(t time.Time) string {
	return t.Format(matchFormat(t))
}

func matchFormat(t time.Time) string {
	if t.Nanosecond() > 0 {
		return DateTimeNano
	}

	if t.Second() > 0 {
		return DateTime
	}

	if t.Minute() > 0 || t.Hour() > 0 {
		return DateHoursMinutes
	}

	if t.Day() != 1 {
		return Date
	}

	if t.Month() != 1 {
		return YearMonth
	}

	return Year
}
