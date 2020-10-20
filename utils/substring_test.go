package utils

import "testing"

func TestSubstring(t *testing.T) {
	var tests = []struct {
		name   string
		from   int
		length int
		want   string
	}{
		{"should return empty string when `from` is gt len(str)", 300, 1, ""},
		{"should return empty string when `from` or `length` are negative", -3, -1, ""},
		{"should return to the end of string when `length` is 0", 2, 0, "llo world"},
		{"should return basic substring", 2, 3, "llo"},
		{"should return to the end of string in case of overflow", 2, 30, "llo world"},
	}

	input := "hello world"

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := Substring(input, testcase.from, testcase.length)
			if got != testcase.want {
				t.Errorf(
					"Substring(\"hello world\", %d, %d) got %v, want %v",
					testcase.from,
					testcase.length,
					got,
					testcase.want,
				)
			}
		})
	}
}
