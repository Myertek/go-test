package db

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Craig"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Craig")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Craig") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
