package w3cdtf_test

import (
	"fmt"
	"log"
	"time"

	w3cdtf "github.com/ak10m/go-w3cdtf"
)

var (
	jst = time.FixedZone("JST", 9*60*60)
)

func ExampleLayout() {
	t, err := time.Parse(time.RFC3339, "2005-04-23T12:34:56.789+09:00")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(t.Format(w3cdtf.Year))
	fmt.Println(t.Format(w3cdtf.YearMonth))
	fmt.Println(t.Format(w3cdtf.Date))
	fmt.Println(t.Format(w3cdtf.DateHoursMinutes))
	fmt.Println(t.Format(w3cdtf.DateTime))
	fmt.Println(t.Format(w3cdtf.DateTimeNano))

	// Output:
	// 2005
	// 2005-04
	// 2005-04-23
	// 2005-04-23T12:34+09:00
	// 2005-04-23T12:34:56+09:00
	// 2005-04-23T12:34:56.789+09:00
}

func ExampleParse() {
	fmt.Println(w3cdtf.MustParse("2005").Format(time.RFC3339))
	fmt.Println(w3cdtf.MustParse("2005-04").Format(time.RFC3339))
	fmt.Println(w3cdtf.MustParse("2005-04-23").Format(time.RFC3339))
	fmt.Println(w3cdtf.MustParse("2005-04-23T12:34Z").Format(time.RFC3339))
	fmt.Println(w3cdtf.MustParse("2005-04-23T12:34+09:00").Format(time.RFC3339))
	fmt.Println(w3cdtf.MustParse("2005-04-23T12:34:56Z").Format(time.RFC3339))
	fmt.Println(w3cdtf.MustParse("2005-04-23T12:34:56+09:00").Format(time.RFC3339))
	fmt.Println(w3cdtf.MustParse("2005-04-23T12:34:56.789Z").Format(time.RFC3339Nano))
	fmt.Println(w3cdtf.MustParse("2005-04-23T12:34:56.789+09:00").Format(time.RFC3339Nano))

	// Output:
	// 2005-01-01T00:00:00Z
	// 2005-04-01T00:00:00Z
	// 2005-04-23T00:00:00Z
	// 2005-04-23T12:34:00Z
	// 2005-04-23T12:34:00+09:00
	// 2005-04-23T12:34:56Z
	// 2005-04-23T12:34:56+09:00
	// 2005-04-23T12:34:56.789Z
	// 2005-04-23T12:34:56.789+09:00
}

func ExampleFormat() {
	fmt.Println(w3cdtf.Format(time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)))
	fmt.Println(w3cdtf.Format(time.Date(2005, 4, 1, 0, 0, 0, 0, time.UTC)))
	fmt.Println(w3cdtf.Format(time.Date(2005, 4, 23, 0, 0, 0, 0, time.UTC)))
	fmt.Println(w3cdtf.Format(time.Date(2005, 4, 23, 12, 34, 0, 0, time.UTC)))
	fmt.Println(w3cdtf.Format(time.Date(2005, 4, 23, 12, 34, 0, 0, jst)))
	fmt.Println(w3cdtf.Format(time.Date(2005, 4, 23, 12, 34, 56, 0, time.UTC)))
	fmt.Println(w3cdtf.Format(time.Date(2005, 4, 23, 12, 34, 56, 0, jst)))
	fmt.Println(w3cdtf.Format(time.Date(2005, 4, 23, 12, 34, 56, 789000000, time.UTC)))
	fmt.Println(w3cdtf.Format(time.Date(2005, 4, 23, 12, 34, 56, 789000000, jst)))

	// Output:
	// 2005
	// 2005-04
	// 2005-04-23
	// 2005-04-23T12:34Z
	// 2005-04-23T12:34+09:00
	// 2005-04-23T12:34:56Z
	// 2005-04-23T12:34:56+09:00
	// 2005-04-23T12:34:56.789Z
	// 2005-04-23T12:34:56.789+09:00
}
