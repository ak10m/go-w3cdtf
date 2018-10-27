package w3cdtf

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	for _, tc := range []struct {
		input    time.Time
		expected string
	}{
		{time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC), "2018"},
		{time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC), "2018-04"},
		{time.Date(2018, 4, 23, 0, 0, 0, 0, time.UTC), "2018-04-23"},
		{time.Date(2018, 4, 23, 12, 0, 0, 0, time.UTC), "2018-04-23T12:00Z"},
		{time.Date(2018, 4, 23, 12, 0, 0, 0, jst), "2018-04-23T12:00+09:00"},
		{time.Date(2018, 4, 23, 12, 34, 0, 0, time.UTC), "2018-04-23T12:34Z"},
		{time.Date(2018, 4, 23, 12, 34, 0, 0, jst), "2018-04-23T12:34+09:00"},
		{time.Date(2018, 4, 23, 12, 34, 56, 0, time.UTC), "2018-04-23T12:34:56Z"},
		{time.Date(2018, 4, 23, 12, 34, 56, 0, jst), "2018-04-23T12:34:56+09:00"},
		{time.Date(2018, 4, 23, 12, 34, 56, 789000000, time.UTC), "2018-04-23T12:34:56.789Z"},
		{time.Date(2018, 4, 23, 12, 34, 56, 789000000, jst), "2018-04-23T12:34:56.789+09:00"},
	} {
		actual := Format(tc.input)
		if actual != tc.expected {
			t.Errorf(" got: %v; want: %v", actual, tc.expected)
		}
	}
}
