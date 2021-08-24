package gosft

import (
	"testing"
	"time"
)

func TestAppend2Digits(t *testing.T) {
	t.Run("min", func(t *testing.T) {
		tests := []struct {
			value int
			want  string
		}{
			{0, "0"},
			{1, "1"},
			{9, "9"},
			{10, "10"},
			{99, "99"},
		}
		for _, c := range tests {
			t.Run(c.want, func(t *testing.T) {
				var buf []byte
				append2DigitsMin(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
	t.Run("pad with zero", func(t *testing.T) {
		tests := []struct {
			value int
			want  string
		}{
			{0, "00"},
			{1, "01"},
			{9, "09"},
			{10, "10"},
			{99, "99"},
		}
		for _, c := range tests {
			t.Run(c.want, func(t *testing.T) {
				var buf []byte
				append2DigitsZero(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
	t.Run("pad with space", func(t *testing.T) {
		tests := []struct {
			value int
			want  string
		}{
			{0, " 0"},
			{1, " 1"},
			{9, " 9"},
			{10, "10"},
			{99, "99"},
		}
		for _, c := range tests {
			t.Run(c.want, func(t *testing.T) {
				var buf []byte
				append2DigitsSpace(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
}

func TestAppend3Digits(t *testing.T) {
	t.Run("pad with zero", func(t *testing.T) {
		tests := []struct {
			value int
			want  string
		}{
			{0, "000"},
			{1, "001"},
			{9, "009"},
			{10, "010"},
			{99, "099"},
			{100, "100"},
			{789, "789"},
		}
		for _, c := range tests {
			t.Run(c.want, func(t *testing.T) {
				var buf []byte
				append3DigitsZero(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
}

func append4DigitsMin(buf *[]byte, i int) {
	var foo bool
	if i >= 1000 {
		*buf = append(*buf, digits[i/1000])
		i %= 1000
		foo = true
	}

	if i >= 100 {
		*buf = append(*buf, digits[i/100])
		i %= 100
		foo = true
	} else if foo {
		*buf = append(*buf, '0')
	}

	if i >= 10 {
		*buf = append(*buf, digits[i/10])
		i %= 10
		foo = true
	} else if foo {
		*buf = append(*buf, '0')
	}

	*buf = append(*buf, digits[i])
}

func append4DigitsSpace(buf *[]byte, i int) {
	offset := 10

	// thousands
	*buf = append(*buf, digits[offset+i/1000])
	if i >= 1000 {
		offset = 0
	}
	i %= 1000

	// hundreds
	*buf = append(*buf, digits[offset+i/100])
	if i >= 100 {
		offset = 0
	}
	i %= 100

	// tens
	*buf = append(*buf, digits[offset+i/10])
	i %= 10

	// ones
	*buf = append(*buf, digits[i])
}

func TestAppend4Digits(t *testing.T) {
	t.Run("min", func(t *testing.T) {
		tests := []struct {
			value int
			want  string
		}{
			{0, "0"},
			{1, "1"},
			{9, "9"},
			{10, "10"},
			{99, "99"},
			{100, "100"},
			{789, "789"},
			{1234, "1234"},
		}
		for _, c := range tests {
			t.Run(c.want, func(t *testing.T) {
				var buf []byte
				append4DigitsMin(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
	t.Run("pad with zero", func(t *testing.T) {
		tests := []struct {
			value int
			want  string
		}{
			{0, "0000"},
			{1, "0001"},
			{9, "0009"},
			{10, "0010"},
			{99, "0099"},
			{100, "0100"},
			{789, "0789"},
			{1234, "1234"},
		}
		for _, c := range tests {
			t.Run(c.want, func(t *testing.T) {
				var buf []byte
				append4DigitsZero(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
	t.Run("pad with space", func(t *testing.T) {
		tests := []struct {
			value int
			want  string
		}{
			{0, "   0"},
			{1, "   1"},
			{9, "   9"},
			{10, "  10"},
			{99, "  99"},
			{100, " 100"},
			{789, " 789"},
			{1234, "1234"},
		}
		for _, c := range tests {
			t.Run(c.want, func(t *testing.T) {
				var buf []byte
				append4DigitsSpace(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
}

func TestAppend6Digits(t *testing.T) {
	tests := []struct {
		value int
		want  string
	}{
		{0, "000000"},
		{1, "000001"},
		{9, "000009"},
		{10, "000010"},
		{99, "000099"},
		{123456, "123456"},
	}
	for _, c := range tests {
		t.Run(c.want, func(t *testing.T) {
			var buf []byte
			append6DigitsZero(&buf, c.value)
			if got, want := string(buf), c.want; got != want {
				t.Errorf("GOT: %q; WANT: %q", got, want)
			}
		})
	}
}

func TestAppend9Digits(t *testing.T) {
	tests := []struct {
		value int
		want  string
	}{
		{0, "000000000"},
		{1, "000000001"},
		{9, "000000009"},
		{10, "000000010"},
		{99, "000000099"},
		{123456789, "123456789"},
	}
	for _, c := range tests {
		t.Run(c.want, func(t *testing.T) {
			var buf []byte
			append9DigitsZero(&buf, c.value)
			if got, want := string(buf), c.want; got != want {
				t.Errorf("GOT: %q; WANT: %q", got, want)
			}
		})
	}
}

func TestFormatter(t *testing.T) {
	// Use the same date-time stamp that Go standard library uses,
	// namely 2006-01-02T15:04:05Z07:00
	when := time.Date(2006, time.January, 2, 3, 4, 5, 12345678, time.UTC)

	tests := []struct {
		format, want string
	}{
		{"%a", "Mon"},                      // The abbreviated name of the day of the week.
		{"%A", "Monday"},                   // The full name of the day of the week.
		{"%b", "Jan"},                      // The abbreviated month name.
		{"%B", "January"},                  // Thee full name of the month.
		{"%c", "Mon Jan  2 03:04:05 2006"}, // Time and date. Equivalent to `%a %b %e %H:%M:%S %Y`.
		{"%C", "20"},                       // The century number (year/100) as a 2-digit integer (00..99).
		{"%d", "02"},                       // The day of the month (01..31).
		{"%D", "01/02/06"},                 // Equivalent to `%m/%d/%y`.
		{"%e", " 2"},                       // Like `%d`, the day of the month ( 1..31).
		// {"%E", "TODO"}, // Modifier: use alternative ("era-based") format.
		{"%F", "2006-01-02"}, // Equivalent to `%Y-%m-%d`, the ISO 8601 date format.
		{"%g", "06"},         // Like `%G`, but without century, that is, with a 2-digit year (00..99).
		{"%G", "2006"},       // The ISO 8601 week-based year with century as a 4-digit decimal number.
		{"%h", "Jan"},        // Equivalent to `%b`.
		{"%H", "03"},         // The hour as a decimal number using a 24-hour clock (00..23).
		{"%I", "03"},         // The hour as a decimal number using a 12-hour clock (01..12).
		{"%j", "002"},        // The day of the year as a decimal number (001..366).
		{"%k", " 3"},         // The hour (24-hour clock) as a decimal number ( 0..23).
		{"%l", " 3"},         // The hour (12-hour clock) as a decimal number ( 1..12).
		{"%m", "01"},         // The month as a decimal number (01..12).
		{"%M", "04"},         // The minute as a decimal number (00..59).
		{"%n", "\n"},         // A newline character.
		// {"%O", "TODO"}, // Modifier: use alternative numeric symbols.
		{"%p", "AM"},          // Either "AM" or "PM" according to the given time value.
		{"%P", "am"},          // Either "am" or "pm" according to the given time value.
		{"%r", "03:04:05 AM"}, // The time in a.m. or p.m. notation. Equivalent to `%I:%M:%S %p`.
		{"%R", "03:04"},       // The time in 24-hour notation. Equivalent to `%H:%M`
		{"%s", "1136171045"},  // The number of seconds since the Epoch, 1970-01-01 00:00:00 +0000 UTC.
		{"%S", "05"},          // The second as a decimal number (00..60).
		{"%t", "\t"},          // A tab character.
		{"%T", "03:04:05"},    // The time in 24-hour notation. Equivalent to `%H:%M:%S`.
		{"%u", "1"},           // The day of the week, (1..7); 1 is Monday.
		// {"%U", "TODO"}, // The week number of the current year as a decimal number.
		// {"%V", "TODO"}, // The ISO 8601 week number of the current year as a decimal number.
		{"%w", "1"}, // The day of the week as a decimal, (0..6); 0 is Sunday.
		// {"%W", "TODO"}, // The week number of the current year as a decimal number.
		{"%x", "01/02/06"}, // Equivalent to `%m/%d/%y`
		{"%X", "03:04:05"}, // Equivalent to `%H:%M:%S`
		{"%y", "06"},       // The year as a decimal number without a century (00..99).
		{"%Y", "2006"},     // The year as a decimal number including the century.
		{"%z", "+0000"},    // The ++hhmm or -hhmm numeric timezone.
		{"%Z", "UTC"},      // The timezone name or abbreviation.
		{"%+", "Mon Jan  2 03:04:05 AM UTC 2006"}, // The date and time in date(1) format.
		{"%%", "%"}, // A % character.

		// Make sure embedded substrings are included.
		{"abc %F def %T ghi", "abc 2006-01-02 def 03:04:05 ghi"},
	}

	for _, c := range tests {
		tf, err := New(c.format)
		ensureError(t, err, nil)

		t.Run("Append", func(t *testing.T) {
			t.Run(c.format, func(t *testing.T) {
				buf := make([]byte, 0, 64)
				if got, want := string(tf.Append(buf, when)), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		})

		t.Run("Format", func(t *testing.T) {
			t.Run(c.format, func(t *testing.T) {
				if got, want := tf.Format(when), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		})
	}
}

func TestWeekdays(t *testing.T) {
	tests := []struct {
		day         int
		short, long string
	}{
		{2, "Mon", "Monday"}, // The second day of this year-month happens to be a Monday.
		{3, "Tue", "Tuesday"},
		{4, "Wed", "Wednesday"},
		{5, "Thu", "Thursday"},
		{6, "Fri", "Friday"},
		{7, "Sat", "Saturday"},
		{8, "Sun", "Sunday"},
	}

	sf, err := New("%a")
	ensureError(t, err, nil)
	lf, err := New("%A")
	ensureError(t, err, nil)

	for _, c := range tests {
		when := time.Date(2006, time.January, c.day, 3, 4, 5, 12345678, time.UTC)
		t.Run("short", func(t *testing.T) {
			if got, want := sf.Format(when), c.short; got != want {
				t.Errorf("GOT: %q; WANT: %q", got, want)
			}
		})
		t.Run("long", func(t *testing.T) {
			if got, want := lf.Format(when), c.long; got != want {
				t.Errorf("GOT: %q; WANT: %q", got, want)
			}
		})
	}
}

func TestMonths(t *testing.T) {
	tests := []struct {
		month       time.Month
		short, long string
	}{
		{time.January, "Jan", "January"},
		{time.February, "Feb", "February"},
		{time.March, "Mar", "March"},
		{time.April, "Apr", "April"},
		{time.May, "May", "May"},
		{time.June, "Jun", "June"},
		{time.July, "Jul", "July"},
		{time.August, "Aug", "August"},
		{time.September, "Sep", "September"},
		{time.October, "Oct", "October"},
		{time.November, "Nov", "November"},
		{time.December, "Dec", "December"},
	}

	sf, err := New("%b")
	ensureError(t, err, nil)
	lf, err := New("%B")
	ensureError(t, err, nil)

	for _, c := range tests {
		when := time.Date(2006, c.month, 2, 3, 4, 5, 12345678, time.UTC)
		t.Run("short", func(t *testing.T) {
			if got, want := sf.Format(when), c.short; got != want {
				t.Errorf("GOT: %q; WANT: %q", got, want)
			}
		})
		t.Run("long", func(t *testing.T) {
			if got, want := lf.Format(when), c.long; got != want {
				t.Errorf("GOT: %q; WANT: %q", got, want)
			}
		})
	}
}

func TestCompatibility(t *testing.T) {
	// Use the same date-time stamp that Go standard library uses,
	// namely 2006-01-02T15:04:05Z07:00
	when := time.Date(2006, time.January, 2, 3, 4, 5, 12345678, time.UTC)

	tests := []string{
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}

	for _, c := range tests {
		t.Run(c, func(t *testing.T) {
			tf, err := NewCompat(c)
			ensureError(t, err, nil)

			if got, want := tf.Format(when), when.Format(c); got != want {
				t.Errorf("GOT: %q; WANT: %q", got, want)
			}
		})
	}
}

func BenchmarkCompatibility(b *testing.B) {
	var err error
	var foo string

	buf := make([]byte, 0, 128)

	// Use the same date-time stamp that Go standard library uses,
	// namely 2006-01-02T15:04:05Z07:00
	when := time.Date(2006, time.January, 2, 3, 4, 5, 12345678, time.UTC)

	tests := []struct {
		format string
		tf     *Formatter
	}{
		{time.ANSIC, nil},
		{time.UnixDate, nil},
		{time.RubyDate, nil},
		{time.RFC822, nil},
		{time.RFC822Z, nil},
		{time.RFC850, nil},
		{time.RFC1123, nil},
		{time.RFC1123Z, nil},
		{time.RFC3339, nil},
		{time.RFC3339Nano, nil},
		{time.Kitchen, nil},
		{time.Kitchen, nil},
		{time.Stamp, nil},
		{time.StampMilli, nil},
		{time.StampMicro, nil},
		{time.StampNano, nil},
	}

	for i, c := range tests {
		tests[i].tf, err = NewCompat(c.format)
		ensureError(b, err, nil)
	}

	b.ResetTimer()

	b.Run("Append", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, c := range tests {
				buf = c.tf.Append(buf, when)
				buf = buf[:0] // reset buffer to append to first byte in pre-allocated slice
			}
		}
	})
	_ = buf

	b.Run("Format", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, c := range tests {
				foo = c.tf.Format(when)
			}
		}
	})
	_ = foo

	b.Run("stdlib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, c := range tests {
				foo = when.Format(c.format)
			}
		}
	})
	_ = foo
}
