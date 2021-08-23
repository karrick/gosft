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
