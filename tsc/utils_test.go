package tsc

import "testing"

func TestIsEvent(t *testing.T) {
	var tests = []struct {
		name string
		str  string
		want bool
	}{
		{"should return true on valid event", "#1234", true},
		{"should return false on invalid length", "#20", false},
		{"should return false if string doesn't start with #", "a", false},
		{"should return false if event # is invalid", "#3ab1", false},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := IsEvent(testcase.str)
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
		name string
		str  string
		want bool
	}{
		{"should return true on valid argument", "0012", true},
		{"should return true on valid argument with variable", "V008", true},
		{"should return false on invalid length", "012", false},
		{"should return false on invalid argument", "3a1b", false},
		{"should return false on invalid argument with variable", "V0AB", false},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := IsValidArgument(testcase.str)
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
