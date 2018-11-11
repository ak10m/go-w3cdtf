# W3C-DTF for golang

## Installation

Use go get.

```bash
go get github.com/ak10m/go-w3cdtf
```

Then import the w3cdtf package into your own code.

```go
import w3cdtf "github.com/ak10m/go-w3cdtf"
```

## Usage

```go
import (
	"log"
	"time"

	w3cdtf "github.com/ak10m/go-w3cdtf"
)

func main() {
	t1, err := time.Parse(time.RFC3339, "2005-04-23T12:34:56.789+09:00")
	if err != nil {
		log.Fatal(err)
	}

	// W3C-DTF layout
	log.Println(t1.Format(w3cdtf.Year))             // 2005
	log.Println(t1.Format(w3cdtf.YearMonth))        // 2005-04
	log.Println(t1.Format(w3cdtf.Date))             // 2005-04-23
	log.Println(t1.Format(w3cdtf.DateHoursMinutes)) // 2005-04-23T12:34+09:00
	log.Println(t1.Format(w3cdtf.DateTime))         // 2005-04-23T12:34:56+09:00
	log.Println(t1.Format(w3cdtf.DateTimeNano))     // 2005-04-23T12:34:56.789+09:00

	// w3cdtf.Format
	t2 := time.Date(2005, 4, 1, 12, 34, 0, 0, time.UTC)
	// returns shortest W3C-DTF format
	log.Println(w3cdtf.Format(t2)) // 2005-04-01T12:34Z

	// w3cdtf.Parse
	t3, err := w3cdtf.Parse("2005-04")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(t3) // 2005-04-01 00:00:00 +0000 UTC

  // w3cdtf.MustParse
	t4 := w3cdtf.MustParse("2005-23")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(t4) // 2005-04-23 00:00:00 +0000 UTC
}
```
