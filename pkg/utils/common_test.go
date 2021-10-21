package utils

import (
	"regexp"
	"testing"
)

func TestGoDotEnvVariable(t *testing.T) {
	item := "jdnfksdmfksd"
	want := regexp.MustCompile(`\b` + item + `\b`)
	msg := GoDotEnvVariable("ACCESS_SECRET")

	if !want.MatchString(msg) {
		t.Fatalf(`GoDotEnvVariable("ACCESS_SECRET") = %q, want match for %#q`, msg, want)
	}
}

func TestGoDotEnvVariableEmpty(t *testing.T) {

	item := "Key env required"
	want := regexp.MustCompile(`\b` + item + `\b`)
	msg := GoDotEnvVariable("")

	if !want.MatchString(msg) {
		t.Fatalf(`GoDotEnvVariable("") = %q, want match for %#q`, msg, want)
	}
}
