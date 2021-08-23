package gosft

import (
	"testing"
	"time"
)

func TestAppend2(t *testing.T) {
	t.Run("2d", func(t *testing.T) {
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
				append2d(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
	t.Run("02d", func(t *testing.T) {
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
				append02d(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
	t.Run("S2d", func(t *testing.T) {
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
				appendS2d(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
}

func TestAppend3(t *testing.T) {
	t.Run("03d", func(t *testing.T) {
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
				append03d(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
}

func append4d(buf *[]byte, i int) {
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

func TestAppend4(t *testing.T) {
	t.Run("4d", func(t *testing.T) {
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
				append4d(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
	t.Run("04d", func(t *testing.T) {
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
				append04d(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
	t.Run("S4d", func(t *testing.T) {
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
				appendS4d(&buf, c.value)
				if got, want := string(buf), c.want; got != want {
					t.Errorf("GOT: %q; WANT: %q", got, want)
				}
			})
		}
	})
}

func TestAppend9(t *testing.T) {
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
			append09d(&buf, c.value)
			if got, want := string(buf), c.want; got != want {
				t.Errorf("GOT: %q; WANT: %q", got, want)
			}
		})
	}
}

func TestFormatter(t *testing.T) {
	when := time.Date(2009, time.February, 5, 5, 0, 57, 12345600, time.UTC)

	tests := []struct {
		format, want string
	}{
		{"%a", "Thu"},
		{"%A", "Thursday"},
		{"%b", "Feb"},
		{"%B", "February"},
		{"%c", "Thu Feb  5 05:00:57 2009"},
		{"%C", "20"},
		{"%d", "05"},
		{"%D", "02/05/09"},
		{"%e", " 5"},
		// {"%E", "TODO"},
		{"%F", "2009-02-05"},
		{"%g", "09"},
		{"%G", "2009"},
		{"%h", "Feb"},
		{"%H", "05"},
		{"%I", "05"},
		{"%j", "036"},
		{"%k", " 5"},
		{"%l", "5"},
		{"%m", "02"},
		{"%M", "00"},
		{"%n", "\n"},
		// {"%O", "TODO"},
		{"%p", "AM"},
		{"%P", "am"},
		{"%r", "05:00:57 AM"},
		{"%R", "05:00"},
		{"%s", "1233810057"},
		{"%S", "57"},
		{"%t", "\t"},
		{"%T", "05:00:57"},
		{"%u", "4"},
		// {"%U", "TODO"},
		// {"%V", "TODO"},
		{"%w", "4"},
		// {"%W", "TODO"},
		{"%x", "02/05/09"},
		{"%X", "05:00:57"},
		{"%y", "09"},
		{"%Y", "2009"},
		{"%z", "+0000"},
		{"%Z", "UTC"},
		// {"%+", "TODO"},
		{"%%", "%"},

		// Make sure embedded substrings are included.
		{"abc %F def %T ghi", "abc 2009-02-05 def 05:00:57 ghi"},
	}
	for _, c := range tests {
		tf, err := New(c.format)
		ensureError(t, err, nil)
		buf := make([]byte, 0, 128)

		t.Run("Append", func(t *testing.T) {
			t.Run(c.format, func(t *testing.T) {
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

func TestCompatibility(t *testing.T) {
	when := time.Date(2009, time.February, 5, 5, 0, 57, 12345678, time.UTC)

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

func Benchmark(b *testing.B) {
	var err error
	var foo string

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

	when := time.Date(2009, time.February, 5, 5, 0, 57, 12345600, time.UTC)
	buf := make([]byte, 0, 100)

	b.ResetTimer()

	b.Run("stdlib", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, c := range tests {
				foo = when.Format(c.format)
			}
		}
	})
	_ = foo

	b.Run("Append", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, c := range tests {
				buf = c.tf.Append(buf, when)
			}
		}
	})
	_ = foo

	b.Run("Format", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, c := range tests {
				foo = c.tf.Format(when)
			}
		}
	})
	_ = foo
}
