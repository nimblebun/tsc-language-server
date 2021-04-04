package tsc_test

import (
	"testing"

	"pkg.nimblebun.works/tsc-language-server/tsc"
)

func TestIsEvent(t *testing.T) {
	var tests = []struct {
		name  string
		str   string
		want  bool
		loose bool
	}{
		{"should return true on valid event", "#1234", true, false},
		{"should return false on invalid length", "#20", false, false},
		{"should return false if string doesn't start with #", "a", false, false},
		{"should return false if event # is invalid", "#3ab1", false, false},
		{"should return true on non-numerical event # with loose setting", "#3ab1", true, true},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := tsc.IsEvent(testcase.str, testcase.loose)
			if got != testcase.want {
				t.Errorf(
					"IsEvent(%s) got %v, want %v",
					testcase.str,
					got,
					testcase.want,
				)
			}
		})
	}
}

func TestIsValidArgument(t *testing.T) {
	var tests = []struct {
		name  string
		str   string
		want  bool
		loose bool
	}{
		{"should return true on valid argument", "0012", true, false},
		{"should return true on valid argument with variable", "V008", true, false},
		{"should return false on invalid length", "012", false, false},
		{"should return false on invalid argument", "3a1b", false, false},
		{"should return false on invalid argument with variable", "V0AB", false, false},
		{"should return true on non-numerical argument with loose setting", "0-12", true, true},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := tsc.IsValidArgument(testcase.str, testcase.loose)
			if got != testcase.want {
				t.Errorf(
					"IsValidArgument(%s) got %v, want %v",
					testcase.str,
					got,
					testcase.want,
				)
			}
		})
	}
}
