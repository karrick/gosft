package gosft

import (
	"errors"
	"strings"
	"testing"
)

func ensureError(tb testing.TB, got, want error) {
	tb.Helper()
	if got == nil {
		if want != nil {
			tb.Fatalf("GOT: %v; WANT: %T(%q)", got, want, want.Error())
		}
	} else if want == nil {
		tb.Fatalf("GOT: %T(%q); WANT: %v", got, got.Error(), want)
	} else {
		var target error
		if ok := errors.As(got, &target); !ok {
			tb.Fatalf("GOT: %T(%q); WANT: %T(%q)", got, got.Error(), want, want.Error())
		}
		if g, w := got.Error(), want.Error(); !strings.Contains(g, w) {
			tb.Fatalf("GOT: %v; WANT: %v", g, w)
		}
	}
}
