package conn

import (
	"regexp"
	"testing"
)

func TestGetConnection(t *testing.T) {
	item := "jdnfksdmfksd"
	want := regexp.MustCompile(`\b` + item + `\b`)
	msg := GoDotEnvVariable("ACCESS_SECRET")

	if !want.MatchString(msg) {
		t.Fatalf(`GoDotEnvVariable("ACCESS_SECRET") = %q, want match for %#q`, msg, want)
	}
}
