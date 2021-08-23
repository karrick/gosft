# gostf

Small Go library to format times in accordance with preconfigured time
format string.

Documentation is available via
[![GoDoc](https://godoc.org/github.com/karrick/gosft?status.svg)](https://godoc.org/github.com/karrick/gosft)
and
[https://pkg.go.dev/github.com/karrick/gosft?tab=doc](https://pkg.go.dev/github.com/karrick/gosft?tab=doc).

## Description

Provides an alternative to using Go's standard library's time
formatting that is similar to using time formatting from other
programming languages.

## Example

```Go
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"

    "github.com/karrick/gosft"
)

func main() {
    tf1, err := gosft.New("%F %T")
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: %s\n", filepath.Base(os.Args[0]), err)
        os.Exit(1)
    }

    when := time.Date(2009, time.February, 5, 5, 0, 57, 12345600, time.UTC)
    fmt.Println(tf1.Format(when))
    // Output: 2009-02-05 05:00:57

    tf2, err := gosft.NewCompat(time.RFC1123Z)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s: %s\n", filepath.Base(os.Args[0]), err)
        os.Exit(1)
    }

    fmt.Println(tf2.Format(when))
    // Output: Thu, 05 Feb 2009 05:00:57 +0000
}
```

## Performance

The primary goal is to be more easy to use when creating code to
format a time in a specific format. A secondary goal is to be about
the same speed as the standard library or better. In the below chart,
the `stdlib` line is the amount of time it takes to format a
particular time using all of the standard library's predefined
formats. The `Append` line is the amount of time it takes for this
library to format the same time using the same formats from the
standard library when appending to a pre-allocated byte slice. The
`Format` line is the amount of time it takes for this library to
allocate space for and format the same time using the same formats
from the standard library.

```Bash
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/karrick/gosft
cpu: AMD Ryzen Threadripper 3960X 24-Core Processor 
Benchmark/stdlib-48         	  337138	      5442 ns/op
Benchmark/Append-48         	  396704	      2576 ns/op
Benchmark/Format-48         	  232258	      5368 ns/op
PASS
ok  	github.com/karrick/gosft	4.997s
```

## Supportd format specifiers

Supports most of the format specifiers from `man 3 strftime`.

|Specifier | Supported | Purpose |
|--|---|---|
| `%a` | Yes | The abbreviated name of the day of the week. |
| `%A` | Yes | The full name of the day of the week. |
| `%b` | Yes | The abbreviated month name. |
| `%B` | Yes | Thee full name of the month. |
| `%c` | Yes | Time and date. Equivalent to `%a %b %e %H:%M:%S %Y` |
| `%C` | Yes | The century number (year/100) as a 2-digit integer. |
| `%d` | Yes | The day of the month as a decimal number (range 01 to 31). |
| `%D` | Yes | Equivalent to %m/%d/%y. |
| `%e` | Yes | Like %d, the ay of the month as a decimal number, but a leading space rather than zero. |
| `%E` | No  | Modifier: use alternative ("era-based") format. |
| `%F` | Yes | Equivalent to %Y-%m-%d (the ISO 8601 date format. |
| `%g` | Yes | Like %G, but without century, that is, with a 2-digit year (00-99). |
| `%G` | Yes | The ISO 8601 week-based year with century as a 4-digit decimal number. |
| `%h` | Yes | Equivalent to %b. |
| `%H` | Yes | The hour as a decimal number using a 24-hour clock (range 00 to 23). |
| `%I` | Yes | The hour as a decimal number using a 12-hour clock (range 01 to 12). |
| `%j` | Yes | The day of the year as a decimal number (range 001 to 366). |
| `%k` | Yes | The hour (24-hour clock) as a decimal number (range 0 to 23). |
| `%l` | Yes | The hour (12-hour clock) as a decimal number (range 1 to 12). |
| `%m` | Yes | The month as a decimal number (range 01 to 12). |
| `%M` | Yes | The minute as a decimal number (range 00 to 59). |
| `%n` | Yes | A newline character. |
| `%O` | No  | Modifier: use alternative numeric symbols. |
| `%p` | Yes | Either "AM" or "PM" according to the given time value. |
| `%P` | Yes | Either "am" or "pm" according to the given time value. |
| `%r` | Yes | The time in a.m. or p.m. notation. Equivalent to %I:%M:%S %p |
| `%R` | Yes | The time in 24-hour notation. Equivalent to %H:%M |
| `%s` | Yes | The number of seconds since the Epoch, 1970-01-01 00:00:00 +0000 UTC. |
| `%S` | Yes | The second as a decimal number (range 00 to 60). |
| `%t` | Yes | A tab character. |
| `%T` | Yes | The time in 24-hour notation (%H:%M:%S). |
| `%u` | Yes | The day of the week as a decomal, range 1 to 7, Monday being 1. |
| `%U` | No  | The week number of the current year as a decimal number. |
| `%V` | No  | The ISO 8601 week number of the current year as a decimal number. |
| `%w` | Yes | The day of the week as a decimal, range 0 to 6, Sunday being 0. |
| `%W` | No  | The week number of the current year as a decimal number. |
| `%x` | Yes | Equivalent to %m/%d/%y |
| `%X` | Yes | Equivalent to %H:%M:%S |
| `%y` | Yes | The year as a decimal number without a century (range 00 to 99). |
| `%Y` | Yes | The year as a decimal number including the century. |
| `%z` | Yes | The ++hhmm or -hhmm numeric timezone. |
| `%Z` | Yes | The timezone name or abbreviation. |
| `%+` | No  | The date and time in date(1) format. |
| `%%` | Yes | A % character. |
