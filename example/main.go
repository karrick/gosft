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
