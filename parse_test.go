package w3cdtf

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	jst, _ := time.LoadLocation("Asia/Tokyo")

	for _, tc := range []struct {
		input    string
		expected time.Time
	}{
		{"2018", time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)},
		{"2018-04", time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)},
		{"2018-04-20", time.Date(2018, 4, 20, 0, 0, 0, 0, time.UTC)},
		{"2018-04-20T12:34Z", time.Date(2018, 4, 20, 12, 34, 0, 0, time.UTC)},
		{"2018-04-20T12:34+09:00", time.Date(2018, 4, 20, 12, 34, 0, 0, jst)},
		{"2018-04-20T12:34:56Z", time.Date(2018, 4, 20, 12, 34, 56, 0, time.UTC)},
		{"2018-04-20T12:34:56+09:00", time.Date(2018, 4, 20, 12, 34, 56, 0, jst)},
		{"2018-04-20T12:34:56.789Z", time.Date(2018, 4, 20, 12, 34, 56, 789000000, time.UTC)},
		{"2018-04-20T12:34:56.789+09:00", time.Date(2018, 4, 20, 12, 34, 56, 789000000, jst)},
	} {
		actual, err := Parse(tc.input)
		if !tc.expected.Equal(actual) {
			t.Errorf(" got: %v; want: %v", actual, tc.expected)
		}
		if err != nil {
			t.Errorf(" got: %v; want: %v", err, nil)
		}
	}
}

func didPanic(f func()) bool {
	didPanic := false

	func() {
		defer func() {
			if e := recover(); e != nil {
				didPanic = true
			}
		}()

		f()
	}()

	return didPanic
}

func TestMustParse(t *testing.T) {
	if didPanic(func() { MustParse("2005-04-23") }) {
		t.Errorf(`MustParse("2005-04-23") should not panic`)
	}

	if !didPanic(func() { MustParse("invalid format") }) {
		t.Errorf(`MustParse("invalid format") should panic`)
	}
}
