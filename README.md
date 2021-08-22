# gostf

Small Go library to format times in accordance with preconfigured time
format string.

Documentation is available via
[![GoDoc](https://godoc.org/github.com/karrick/gosft?status.svg)](https://godoc.org/github.com/karrick/gosft)
and
[https://pkg.go.dev/github.com/karrick/gosft?tab=doc](https://pkg.go.dev/github.com/karrick/gosft?tab=doc).

## Description

Supports most of the format specifiers from `man 3 strftime`. Does not
currently support the following format specifiers:

|Specifier | Purpose |
|---|---|
|    %E | Modifier: use alternative ("era-based") format. |
|    %O | Modifier: use alternative numeric symbols. |
|    %U | The week number of the current year as a decimal number. |
|    %V | The ISO 8601 week number of the current year as a decimal number. |
|    %W | The week number of the current year as a decimal number. |
|    %z | The ++hhmm or -hhmm numeric timezone. |
|    %Z | The timezone name or abbreviation. |
|    %+ | The date and time in date(1) format. |

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
	tf, err := gosft.New("%F %T")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}

	when := time.Date(2009, time.February, 5, 5, 0, 57, 12345600, time.UTC)
	fmt.Println(tf.Format(when))
	// Output: 2009-02-05 05:00:57
}
```
