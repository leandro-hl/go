package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	// Arrange
	name := "Ramona"
	want := regexp.MustCompile(`\b` + name + `\b`)

	// Act
	msg, err := Hello(name)

	// Assert
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(
			`Hello(`+name+`) = %q, %v, want match for (expected) %#q, nil`,
			msg,
			err,
			want,
		)
	}
}

func TestHelloEmpty(t *testing.T) {
	// Act
	msg, err := Hello("")

	// Assert
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", err`, msg, err)
	}
}
