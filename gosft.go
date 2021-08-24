package gosft

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"
)

// Formatter will format time.Time values in accordance with their
// configured format specification. A single Formatter may safely be
// used by multiple Go routines simultaneously.
type Formatter struct {
	formatters []func(*[]byte, time.Time)
	size       int
}

var formatMap map[string]string

func init() {
	formatMap = map[string]string{
		time.ANSIC:       "%c",
		time.UnixDate:    "%a %b %e %T %Z %Y",  // "Mon Jan _2 15:04:05 MST 2006",
		time.RubyDate:    "%a %b %d %T %z %Y",  // "Mon Jan 02 15:04:05 -0700 2006",
		time.RFC822:      "%d %b %y %R %Z",     // "02 Jan 06 15:04 MST",
		time.RFC822Z:     "%d %b %y %R %z",     // "02 Jan 06 15:04 -0700",
		time.RFC850:      "%A, %d-%b-%y %T %Z", // "Monday, 02-Jan-06 15:04:05 MST",
		time.RFC1123:     "%a, %d %b %Y %T %Z", // "Mon, 02 Jan 2006 15:04:05 MST",
		time.RFC1123Z:    "%a, %d %b %Y %T %z", // "Mon, 02 Jan 2006 15:04:05 -0700",
		time.RFC3339:     "%Y-%m-%dT%T%1",      // "2006-01-02T15:04:05Z07:00", // TODO: %1 not standard
		time.RFC3339Nano: "%Y-%m-%dT%T.%N%1",   // "2006-01-02T15:04:05.999999999Z07:00", // TODO: %1 not standard
		time.Kitchen:     "%2:%M%p",            // "3:04PM", // TODO: %2 not standard
		time.Stamp:       "%b %e %T",           // "Jan _2 15:04:05"
		time.StampMilli:  "%b %e %T.%3",        // "Jan _2 15:04:05.000" // TODO: %3 not standard
		time.StampMicro:  "%b %e %T.%4",        // "Jan _2 15:04:05.000000" // TODO: %4 not standard
		time.StampNano:   "%b %e %T.%N",        // "Jan _2 15:04:05.000000000"
	}
}

// New returns a formatter that formats times according to the
// provided format string.
func New(format string) (*Formatter, error) {
	return create(format, false)
}

// NewCompat returns a formatter that formats times according to the
// provided Go standard library compatible time string format.
func NewCompat(format string) (*Formatter, error) {
	value, ok := formatMap[format]
	if !ok {
		return nil, fmt.Errorf("cannot find equivalent for time format string: %q", format)
	}
	return create(value, true)
}

func create(format string, special bool) (*Formatter, error) {
	// Build slice of formatting functions, each will emit the
	// requested information.
	var formatters []func(*[]byte, time.Time)

	var buf []byte
	var foundPercent bool

	for ri, rune := range format {
		if !foundPercent {
			if rune == '%' {
				foundPercent = true
				if len(buf) > 0 {
					formatters = append(formatters, makeStringFormatter(buf))
					buf = nil
				}
			} else {
				appendRune(&buf, rune)
			}
			continue
		}
		switch rune {
		case 'a':
			formatters = append(formatters, appendWeekdayShort)
		case 'A':
			formatters = append(formatters, appendWeekdayLong)
		case 'b':
			formatters = append(formatters, appendMonthShort)
		case 'B':
			formatters = append(formatters, appendMonthLong)
		case 'c':
			formatters = append(formatters, appendC)
		case 'C':
			formatters = append(formatters, appendCC)
		case 'd':
			formatters = append(formatters, appendD)
		case 'D':
			formatters = append(formatters, appendDC)
		case 'e':
			formatters = append(formatters, appendE)
		case 'F':
			formatters = append(formatters, appendFC)
		case 'g':
			formatters = append(formatters, appendG)
		case 'G':
			formatters = append(formatters, appendGC)
		case 'h':
			formatters = append(formatters, appendMonthShort)
		case 'H':
			formatters = append(formatters, appendHC)
		case 'I':
			formatters = append(formatters, appendIC)
		case 'j':
			formatters = append(formatters, appendJ)
		case 'k':
			formatters = append(formatters, appendK)
		case 'l':
			formatters = append(formatters, appendL)
		case 'm':
			formatters = append(formatters, appendM)
		case 'M':
			formatters = append(formatters, appendMC)
		case 'n':
			formatters = append(formatters, appendN)
		case 'N':
			formatters = append(formatters, appendNC)
		case 'p':
			formatters = append(formatters, appendP)
		case 'P':
			formatters = append(formatters, appendPC)
		case 'r':
			formatters = append(formatters, appendR)
		case 'R':
			formatters = append(formatters, appendRC)
		case 's':
			formatters = append(formatters, appendS)
		case 'S':
			formatters = append(formatters, appendSC)
		case 't':
			formatters = append(formatters, appendT)
		case 'T':
			formatters = append(formatters, appendTC)
		case 'u':
			formatters = append(formatters, appendU)
		case 'w':
			formatters = append(formatters, appendW)
		case 'x':
			formatters = append(formatters, appendX)
		case 'X':
			formatters = append(formatters, appendXC)
		case 'y':
			formatters = append(formatters, appendY)
		case 'Y':
			formatters = append(formatters, appendYC)
		case 'z':
			formatters = append(formatters, appendZ)
		case 'Z':
			formatters = append(formatters, appendZC)
		case '%':
			formatters = append(formatters, appendPercent)
		case '+':
			formatters = append(formatters, appendPlus)
		case '1':
			if !special {
				return nil, fmt.Errorf("cannot recognize format verb %q at index %d", rune, ri)
			}
			formatters = append(formatters, appendTZ)
		case '2':
			if !special {
				return nil, fmt.Errorf("cannot recognize format verb %q at index %d", rune, ri)
			}
			formatters = append(formatters, appendLMin)
		case '3':
			if !special {
				return nil, fmt.Errorf("cannot recognize format verb %q at index %d", rune, ri)
			}
			formatters = append(formatters, appendMilli)
		case '4':
			if !special {
				return nil, fmt.Errorf("cannot recognize format verb %q at index %d", rune, ri)
			}
			formatters = append(formatters, appendMicro)
		default:
			return nil, fmt.Errorf("cannot recognize format verb %q at index %d", rune, ri)
		}
		foundPercent = false
	}

	if foundPercent {
		return nil, errors.New("cannot find closing format verb")
	}

	if len(buf) > 0 {
		formatters = append(formatters, makeStringFormatter(buf))
	}

	// When instantiating a formatter, want to calculate and store the
	// longest slice of bytes that are needed to format any time using
	// the specified time format string. For this reason, create a
	// time.Time instance that has the longest month name, September,
	// and the longest weekday name, Thursday, and format it using the
	// provided time format specification, count the length of that
	// string, and store that in the time formatter's size field. Then
	// whenever need to format a new time, pre-allocate a byte slice
	// with that number of bytes, and format the provided time into
	// that byte slice.
	when := time.Date(2021, time.September, 30, 23, 59, 59, 123456789, time.UTC)

	tf := &Formatter{formatters: formatters}
	tf.size = len(tf.Format(when))

	return tf, nil
}

// Append will format t in accordance with its preconfigured format
// specification and append the formatted bytes to buf.
func (tf *Formatter) Append(buf []byte, t time.Time) []byte {
	for _, f := range tf.formatters {
		f(&buf, t)
	}
	return buf
}

// Format will format t and return a string in accordance with its
// preconfigured format specification.
func (tf *Formatter) Format(t time.Time) string {
	return string(tf.Append(make([]byte, 0, tf.size), t))
}

func makeStringFormatter(value []byte) func(*[]byte, time.Time) {
	return func(buf *[]byte, _ time.Time) {
		*buf = append(*buf, value...)
	}
}

func appendRune(buf *[]byte, r rune) {
	if r < utf8.RuneSelf {
		*buf = append(*buf, byte(r))
		return
	}
	olen := len(*buf)
	*buf = append(*buf, 0, 0, 0, 0)              // grow buf large enough to accommodate largest possible UTF8 sequence
	n := utf8.EncodeRune((*buf)[olen:olen+4], r) // encode rune into newly allocated buf space
	*buf = (*buf)[:olen+n]                       // trim buf to actual size used by rune addition
}

// digits is two concatenated string slices that allow using an offset
// of 10 to index into space padded values when necessary.
const digits = "0123456789 123456789"

// Dividend ÷ Divisor = Quotient

func append2DigitsMin(buf *[]byte, i int) {
	quotient := i / 10
	remainder := i % 10
	if quotient > 0 {
		*buf = append(*buf, digits[quotient])
	}
	*buf = append(*buf, digits[remainder])
}

func append2DigitsZero(buf *[]byte, i int) {
	quotient := i / 10
	remainder := i % 10
	*buf = append(*buf, digits[quotient])
	*buf = append(*buf, digits[remainder])
}

func append2DigitsSpace(buf *[]byte, i int) {
	quotient := i / 10
	remainder := i % 10
	*buf = append(*buf, digits[10+quotient])
	*buf = append(*buf, digits[remainder])
}

func append3DigitsZero(buf *[]byte, i int) {
	// hundreds
	quotient := i / 100
	remainder := i % 100
	*buf = append(*buf, digits[quotient])

	// tens
	quotient = remainder / 10
	remainder %= 10
	*buf = append(*buf, digits[quotient])

	// ones
	*buf = append(*buf, digits[remainder])
}

func append4DigitsZero(buf *[]byte, i int) {
	// thousands
	quotient := i / 1000
	remainder := i % 1000
	*buf = append(*buf, digits[quotient])

	// hundreds
	quotient = remainder / 100
	remainder %= 100
	*buf = append(*buf, digits[quotient])

	// tens
	quotient = remainder / 10
	remainder %= 10
	*buf = append(*buf, digits[quotient])

	// ones
	*buf = append(*buf, digits[remainder])
}

func append6DigitsZero(buf *[]byte, i int) {
	// hundred-thousands
	quotient := i / 100000
	remainder := i % 100000
	*buf = append(*buf, digits[quotient])

	// ten-thousands
	quotient = remainder / 10000
	remainder = i % 10000
	*buf = append(*buf, digits[quotient])

	// thousands
	quotient = remainder / 1000
	remainder = i % 1000
	*buf = append(*buf, digits[quotient])

	// hundreds
	quotient = remainder / 100
	remainder %= 100
	*buf = append(*buf, digits[quotient])

	// tens
	quotient = remainder / 10
	remainder %= 10
	*buf = append(*buf, digits[quotient])

	// ones
	*buf = append(*buf, digits[remainder])
}

func append9DigitsZero(buf *[]byte, i int) {
	// hundred-millions
	//              123456789
	quotient := i / 100000000
	remainder := i % 100000000
	*buf = append(*buf, digits[quotient])

	// ten-millions
	quotient = remainder / 10000000
	remainder = i % 10000000
	*buf = append(*buf, digits[quotient])

	// millions
	quotient = remainder / 1000000
	remainder = i % 1000000
	*buf = append(*buf, digits[quotient])

	// hundred-thousands
	quotient = remainder / 100000
	remainder = i % 100000
	*buf = append(*buf, digits[quotient])

	// ten-thousands
	quotient = remainder / 10000
	remainder = i % 10000
	*buf = append(*buf, digits[quotient])

	// thousands
	quotient = remainder / 1000
	remainder = i % 1000
	*buf = append(*buf, digits[quotient])

	// hundreds
	quotient = remainder / 100
	remainder %= 100
	*buf = append(*buf, digits[quotient])

	// tens
	quotient = remainder / 10
	remainder %= 10
	*buf = append(*buf, digits[quotient])

	// ones
	*buf = append(*buf, digits[remainder])
}

//                    0         1         2         3         4         5
//                    012345678901234567890123456789012345678901234567890
const weekdaysLong = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"

var weekdaysLongIndices = []int{0, 6, 12, 19, 28, 36, 42, 50} // 50 is extra index to cover final entry in list

func appendWeekdayShort(buf *[]byte, t time.Time) {
	// %a     The  abbreviated  name  of  the day of the week according to the
	//        current locale.  (Calculated from tm_wday.)  (The specific names
	//        used  in  the current locale can be obtained by calling nl_lang‐
	//        info(3) with ABDAY_{1–7} as an argument.)
	index := weekdaysLongIndices[t.Weekday()]
	*buf = append(*buf, weekdaysLong[index:index+3]...)
}

func appendWeekdayLong(buf *[]byte, t time.Time) {
	// %A     The full name of the day of the week according  to  the  current
	//        locale.  (Calculated from tm_wday.)  (The specific names used in
	//        the current locale can be  obtained  by  calling  nl_langinfo(3)
	//        with DAY_{1–7} as an argument.)
	index := t.Weekday()
	*buf = append(*buf, weekdaysLong[weekdaysLongIndices[index]:weekdaysLongIndices[index+1]]...)
}

//                  0         1         2         3         4         5         6         7
//                  012345678901234567890123456789012345678901234567890123456789012345678901234
const monthsLong = "JanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNovemberDecember"

var monthsLongIndices = []int{0, 7, 15, 20, 25, 28, 32, 36, 42, 51, 58, 66, 74} // 74 is extra index to cover final entry in list

func appendMonthShort(buf *[]byte, t time.Time) {
	// %b     The  abbreviated  month  name  according  to the current locale.
	//        (Calculated from tm_mon.)  (The specific names used in the  cur‐
	//        rent  locale  can be obtained by calling nl_langinfo(3) with AB‐
	//        MON_{1–12} as an argument.)
	index := monthsLongIndices[t.Month()-1]
	*buf = append(*buf, monthsLong[index:index+3]...)
}

func appendMonthLong(buf *[]byte, t time.Time) {
	// %B     The full month name according to the  current  locale.   (Calcu‐
	//        lated from tm_mon.)  (The specific names used in the current lo‐
	//        cale can be obtained by calling nl_langinfo(3)  with  MON_{1–12}
	//        as an argument.)
	index := t.Month() - 1
	*buf = append(*buf, monthsLong[monthsLongIndices[index]:monthsLongIndices[index+1]]...)
}

func appendC(buf *[]byte, t time.Time) {
	// %c     The  preferred  date and time representation for the current lo‐
	//        cale.  (The specific format used in the current  locale  can  be
	//        obtained  by  calling nl_langinfo(3) with D_T_FMT as an argument
	//        for the %c conversion specification, and  with  ERA_D_T_FMT  for
	//        the %Ec conversion specification.)  (In the POSIX locale this is
	//        equivalent to %a %b %e %H:%M:%S %Y.)
	appendWeekdayShort(buf, t)
	*buf = append(*buf, ' ')
	appendMonthShort(buf, t)
	*buf = append(*buf, ' ')

	append2DigitsSpace(buf, t.Day()) // appendE(buf, t)

	*buf = append(*buf, ' ')
	appendTC(buf, t)

	*buf = append(*buf, ' ')
	appendYC(buf, t)
}

func appendCC(buf *[]byte, t time.Time) {
	// %C     The century number (year/100) as a 2-digit  integer.  (SU)  (The
	//        %EC  conversion  specification  corresponds  to  the name of the
	//        era.)  (Calculated from tm_year.)
	year := t.Year()
	append2DigitsZero(buf, year/100)
}

func appendD(buf *[]byte, t time.Time) {
	// %d     The day of the month as a  decimal  number  (range  01  to  31).
	//        (Calculated from tm_mday.)
	append2DigitsZero(buf, t.Day())
}

func appendDC(buf *[]byte, t time.Time) {
	// %D     Equivalent  to  %m/%d/%y.  (Yecch—for Americans only.  Americans
	//        should note that in other countries %d/%m/%y is  rather  common.
	//        This  means that in international context this format is ambigu‐
	//        ous and should not be used.) (SU)
	year, month, day := t.Date()

	append2DigitsZero(buf, int(month))
	*buf = append(*buf, '/')

	append2DigitsZero(buf, day)
	*buf = append(*buf, '/')

	append2DigitsZero(buf, year%100)
}

func appendE(buf *[]byte, t time.Time) {
	// %e     Like %d, the day of the month as a decimal number, but a leading
	//        zero is replaced by a space. (SU) (Calculated from tm_mday.)
	append2DigitsSpace(buf, t.Day())
}

// func appendEC(buf *[]byte, t time.Time) {
// 	// %E     Modifier: use alternative ("era-based") format, see below. (SU)
// }

func appendFC(buf *[]byte, t time.Time) {
	// %F     Equivalent to %Y-%m-%d (the ISO 8601 date format). (C99)
	year, month, day := t.Date()

	append4DigitsZero(buf, year)
	*buf = append(*buf, '-')

	append2DigitsZero(buf, int(month))
	*buf = append(*buf, '-')

	append2DigitsZero(buf, day)
}

func appendG(buf *[]byte, t time.Time) {
	// %g     Like %G, but without century,  that  is,  with  a  2-digit  year
	//        (00–99). (TZ) (Calculated from tm_year, tm_yday, and tm_wday.)
	year, _ := t.ISOWeek()
	append2DigitsZero(buf, year%100)
}

func appendGC(buf *[]byte, t time.Time) {
	// %G     The ISO 8601 week-based year (see NOTES) with century as a deci‐
	//        mal number.  The 4-digit year corresponding to the ISO week num‐
	//        ber  (see %V).  This has the same format and value as %Y, except
	//        that if the ISO week number belongs  to  the  previous  or  next
	//        year,  that year is used instead. (TZ) (Calculated from tm_year,
	//        tm_yday, and tm_wday.)
	year, _ := t.ISOWeek()
	append4DigitsZero(buf, year)
}

func appendHC(buf *[]byte, t time.Time) {
	// %H     The  hour as a decimal number using a 24-hour clock (range 00 to
	//        23).  (Calculated from tm_hour.)
	append2DigitsZero(buf, t.Hour())
}

func appendIC(buf *[]byte, t time.Time) {
	// %I     The hour as a decimal number using a 12-hour clock (range 01  to
	//        12).  (Calculated from tm_hour.)
	hour := t.Hour()
	if hour > 12 {
		hour -= 12
	}
	append2DigitsZero(buf, hour)
}

func appendJ(buf *[]byte, t time.Time) {
	// %j     The  day  of  the  year  as a decimal number (range 001 to 366).
	//        (Calculated from tm_yday.)
	append3DigitsZero(buf, t.YearDay())
}

func appendK(buf *[]byte, t time.Time) {
	// %k     The hour (24-hour clock) as a decimal number (range  0  to  23);
	//        single  digits are preceded by a blank.  (See also %H.)  (Calcu‐
	//        lated from tm_hour.)  (TZ)
	append2DigitsSpace(buf, t.Hour())
}

func appendL(buf *[]byte, t time.Time) {
	// %l     The hour (12-hour clock) as a decimal number (range  1  to  12);
	//        single  digits are preceded by a blank.  (See also %I.)  (Calcu‐
	//        lated from tm_hour.)  (TZ)
	hour := t.Hour()
	if hour > 12 {
		hour -= 12
	}
	append2DigitsSpace(buf, hour)
}

func appendLMin(buf *[]byte, t time.Time) {
	// only used to support time.Kitchen.
	hour := t.Hour()
	if hour > 12 {
		hour -= 12
	}
	append2DigitsMin(buf, hour)
}

func appendM(buf *[]byte, t time.Time) {
	// %m     The month as a decimal number (range  01  to  12).   (Calculated
	//        from tm_mon.)
	append2DigitsZero(buf, int(t.Month()))
}

func appendMC(buf *[]byte, t time.Time) {
	// %M     The  minute  as  a decimal number (range 00 to 59).  (Calculated
	//        from tm_min.)
	append2DigitsZero(buf, t.Minute())
}

func appendN(buf *[]byte, t time.Time) {
	// %n     A newline character. (SU)
	*buf = append(*buf, '\n')
}

func appendNC(buf *[]byte, t time.Time) {
	append9DigitsZero(buf, t.Nanosecond())
}

func appendMicro(buf *[]byte, t time.Time) {
	append6DigitsZero(buf, t.Nanosecond()/1000)
}

func appendMilli(buf *[]byte, t time.Time) {
	append3DigitsZero(buf, t.Nanosecond()/1000000)
}

// func appendOC(buf *[]byte, t time.Time) {
// 	// %O     Modifier: use alternative numeric symbols, see below. (SU)
// }

func appendP(buf *[]byte, t time.Time) {
	// %p     Either "AM" or "PM" according to the given time  value,  or  the
	//        corresponding  strings  for the current locale.  Noon is treated
	//        as "PM" and midnight as "AM".  (Calculated from tm_hour.)   (The
	//        specific  string  representations  used for "AM" and "PM" in the
	//        current locale can be obtained by  calling  nl_langinfo(3)  with
	//        AM_STR and PM_STR, respectively.)
	hour := t.Hour()
	if hour < 12 {
		*buf = append(*buf, []byte("AM")...)
	} else {
		*buf = append(*buf, []byte("PM")...)
	}
}

func appendPC(buf *[]byte, t time.Time) {
	// %P     Like %p but in lowercase: "am" or "pm" or a corresponding string
	//        for the current locale.  (Calculated from tm_hour.)  (GNU)
	hour := t.Hour()
	if hour < 12 {
		*buf = append(*buf, []byte("am")...)
	} else {
		*buf = append(*buf, []byte("pm")...)
	}
}

func appendR(buf *[]byte, t time.Time) {
	// %r     The time in a.m. or p.m. notation.  (SU)  (The  specific  format
	//        used  in  the current locale can be obtained by calling nl_lang‐
	//        info(3) with T_FMT_AMPM as an argument.)  (In the  POSIX  locale
	//        this is equivalent to %I:%M:%S %p.)
	// 09:24:14 PM
	var pm bool
	hour, minute, second := t.Clock()

	if hour >= 12 {
		pm = true
		if hour > 12 {
			hour -= 12
		}
	}

	append2DigitsZero(buf, hour)
	*buf = append(*buf, ':')

	append2DigitsZero(buf, minute)
	*buf = append(*buf, ':')

	append2DigitsZero(buf, second)
	*buf = append(*buf, ' ')

	if pm {
		*buf = append(*buf, []byte("PM")...)
	} else {
		*buf = append(*buf, []byte("AM")...)
	}
}

func appendRC(buf *[]byte, t time.Time) {
	// %R     The  time  in  24-hour notation (%H:%M).  (SU) For a version in‐
	//        cluding the seconds, see %T below.
	hour, minute, _ := t.Clock()

	append2DigitsZero(buf, hour)
	*buf = append(*buf, ':')
	append2DigitsZero(buf, minute)
}

func appendS(buf *[]byte, t time.Time) {
	// %s     The number of seconds since the Epoch, 1970-01-01 00:00:00 +0000
	//        (UTC). (TZ) (Calculated from mktime(tm).)
	*buf = strconv.AppendInt(*buf, t.Unix(), 10)
}

func appendSC(buf *[]byte, t time.Time) {
	// %S     The  second as a decimal number (range 00 to 60).  (The range is
	//        up to 60 to allow for  occasional  leap  seconds.)   (Calculated
	//        from tm_sec.)
	append2DigitsZero(buf, t.Second())
}

func appendT(buf *[]byte, t time.Time) {
	// %t     A tab character. (SU)
	*buf = append(*buf, '\t')
}

func appendTC(buf *[]byte, t time.Time) {
	// %T     The time in 24-hour notation (%H:%M:%S).  (SU)
	hour, minute, second := t.Clock()

	append2DigitsZero(buf, hour)
	*buf = append(*buf, ':')

	append2DigitsZero(buf, minute)
	*buf = append(*buf, ':')

	append2DigitsZero(buf, second)
}

func appendU(buf *[]byte, t time.Time) {
	// %u     The  day of the week as a decimal, range 1 to 7, Monday being 1.
	//        See also %w.  (Calculated from tm_wday.)  (SU)
	if wd := t.Weekday(); wd > 0 {
		*buf = append(*buf, byte(wd+'0'))
	} else {
		*buf = append(*buf, '7')
	}
}

// func appendUC(buf *[]byte, t time.Time) {
// 	// %U     The week number of the current year as a decimal  number,  range
// 	//        00  to  53,  starting  with the first Sunday as the first day of
// 	//        week 01.  See also %V and  %W.   (Calculated  from  tm_yday  and
// 	//        tm_wday.)
// 	_, week := t.ISOWeek()
// 	if week < 10 {
// 		*buf = append(*buf, '0')
// 	}
// 	*buf = strconv.AppendInt(*buf, int64(week), 10)
// }

// func appendVC(buf *[]byte, t time.Time) {
// 	// %V     The  ISO 8601  week  number (see NOTES) of the current year as a
// 	//        decimal number, range 01 to 53, where week 1 is the  first  week
// 	//        that  has  at least 4 days in the new year.  See also %U and %W.
// 	//        (Calculated from tm_year, tm_yday, and tm_wday.)  (SU)
// 	_, week := t.ISOWeek()
// 	if week < 10 {
// 		*buf = append(*buf, '0')
// 	}
// 	*buf = strconv.AppendInt(*buf, int64(week), 10)
// }

func appendW(buf *[]byte, t time.Time) {
	// %w     The day of the week as a decimal, range 0 to 6, Sunday being  0.
	//        See also %u.  (Calculated from tm_wday.)
	*buf = append(*buf, byte(t.Weekday()+'0'))
}

// func appendWC(buf *[]byte, t time.Time) {
// 	// %W     The  week  number of the current year as a decimal number, range
// 	//        00 to 53, starting with the first Monday as  the  first  day  of
// 	//        week 01.  (Calculated from tm_yday and tm_wday.)
// }

func appendX(buf *[]byte, t time.Time) {
	// %x     The preferred date representation for the current locale without
	//        the time.  (The specific format used in the current  locale  can
	//        be  obtained by calling nl_langinfo(3) with D_FMT as an argument
	//        for the %x conversion specification, and with ERA_D_FMT for  the
	//        %Ex  conversion  specification.)   (In  the POSIX locale this is
	//        equivalent to %m/%d/%y.)
	// 08/20/2021
	appendDC(buf, t)
}

func appendXC(buf *[]byte, t time.Time) {
	// %X     The preferred time representation for the current locale without
	//        the  date.   (The specific format used in the current locale can
	//        be obtained by calling nl_langinfo(3) with T_FMT as an  argument
	//        for  the %X conversion specification, and with ERA_T_FMT for the
	//        %EX conversion specification.)  (In the  POSIX  locale  this  is
	//        equivalent to %H:%M:%S.)
	appendTC(buf, t)
}

func appendY(buf *[]byte, t time.Time) {
	// %y     The year as a decimal number without a century (range 00 to 99).
	//        (The %Ey conversion specification corresponds to the year  since
	//        the  beginning of the era denoted by the %EC conversion specifi‐
	//        cation.)  (Calculated from tm_year)
	append2DigitsZero(buf, t.Year()%100)
}

func appendYC(buf *[]byte, t time.Time) {
	// %Y     The year as a decimal number including the  century.   (The  %EY
	//        conversion  specification  corresponds  to  the full alternative
	//        year representation.)  (Calculated from tm_year)
	append4DigitsZero(buf, t.Year())
}

func appendZ(buf *[]byte, t time.Time) {
	// %z     The +hhmm or -hhmm numeric  timezone  (that  is,  the  hour  and
	//        minute offset from UTC). (SU)
	_, offset := t.Zone()
	hour := offset / 60
	minute := offset % 60
	if offset >= 0 {
		*buf = append(*buf, '+')
	} else {
		*buf = append(*buf, '-')
	}
	append2DigitsZero(buf, hour)
	append2DigitsZero(buf, minute)
}

func appendZC(buf *[]byte, t time.Time) {
	// %Z     The timezone name or abbreviation.
	name, _ := t.Zone()
	*buf = append(*buf, name...)
}

func appendTZ(buf *[]byte, t time.Time) {
	_, offset := t.Zone()
	if offset == 0 {
		*buf = append(*buf, 'Z')
	} else {
		hour := offset / 60
		minute := offset % 60
		if offset < 0 {
			*buf = append(*buf, '-')
		}
		append2DigitsZero(buf, hour)
		*buf = append(*buf, ':')
		append2DigitsZero(buf, minute)
	}
}

func appendPercent(buf *[]byte, t time.Time) {
	// %%     A literal '%' character.
	*buf = append(*buf, '%')
}

func appendPlus(buf *[]byte, t time.Time) {
	// %+     The  date  and  time  in  date(1) format. (TZ) (Not supported in
	//        glibc2.)
	// "%a %b %e %T %p %Z %Y"
	appendWeekdayShort(buf, t)
	*buf = append(*buf, ' ')
	appendMonthShort(buf, t)
	*buf = append(*buf, ' ')
	appendE(buf, t)
	*buf = append(*buf, ' ')
	appendTC(buf, t)
	*buf = append(*buf, ' ')
	appendP(buf, t)
	*buf = append(*buf, ' ')
	appendZC(buf, t)
	*buf = append(*buf, ' ')
	appendYC(buf, t)
}
