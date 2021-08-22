package gosft

import (
	"testing"
	"time"
)

func append2d(buf *[]byte, i int) {
	quotient := i / 10
	remainder := i % 10
	if quotient > 0 {
		*buf = append(*buf, digits[quotient])
	}
	*buf = append(*buf, digits[remainder])
}

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

func TestTimeFormatter(t *testing.T) {
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
		{"%F", "2009-02-05"},
		{"%g", "09"},
		{"%G", "2009"},
		{"%h", "Feb"},
		{"%H", "05"},
		{"%I", "05"},
		{"%j", "036"},
		{"%k", " 5"},
		{"%l", " 5"},
		{"%m", "02"},
		{"%M", "00"},
		{"%n", "\n"},
		{"%p", "AM"},
		{"%P", "am"},
		{"%r", "05:00:57 AM"},
		{"%R", "05:00"},
		{"%s", "1233810057"},
		{"%S", "57"},
		{"%t", "\t"},
		{"%T", "05:00:57"},
		{"%u", "4"},
		{"%w", "4"},
		{"%x", "02/05/09"},
		{"%X", "05:00:57"},
		{"%y", "09"},
		{"%Y", "2009"},
		{"%%", "%"},
		//
		{"abc %F def %T ghi", "abc 2009-02-05 def 05:00:57 ghi"},
	}
	for _, c := range tests {
		t.Run(c.format, func(t *testing.T) {
			tf, err := New(c.format)
			ensureError(t, err, nil)

			if got, want := tf.Format(when), c.want; got != want {
				t.Errorf("GOT: %q; WANT: %q", got, want)
			}
		})
	}
}
