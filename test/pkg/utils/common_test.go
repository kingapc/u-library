package utils

import (
	"regexp"
	"testing"

	"github.com/rpinedafocus/u-library/pkg/utils"
)

func TestGoDotEnvVariable(t *testing.T) {
	item := "postgres"
	want := regexp.MustCompile(`\b` + item + `\b`)
	msg := utils.GoDotEnvVariable("DBUSER")

	if !want.MatchString(msg) {
		t.Fatalf(`GoDotEnvVariable("DBUSER") = %q, want match for %#q`, msg, want)
	}
}

func TestGoDotEnvVariableEmpty(t *testing.T) {

	item := "Key env required"
	want := regexp.MustCompile(`\b` + item + `\b`)
	msg := utils.GoDotEnvVariable("")

	if !want.MatchString(msg) {
		t.Fatalf(`GoDotEnvVariable("") = %q, want match for %#q`, msg, want)
	}
}
