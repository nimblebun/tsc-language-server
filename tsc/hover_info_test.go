package tsc_test

import (
	"testing"

	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/tsc"
)

func TestGetHoverInfo(t *testing.T) {
	var tests = []struct {
		name string
		str  string
		at   int
		want string
	}{
		{"should return event", "#0069test", 0, "Event **#0069**"},
		{"should return empty string on nothing", "sue", 1, ""},
		{"should return empty string on unknown command", "<SUE0001", 1, ""},
		{
			"should return command name and description",
			"<MNA",
			1,
			"Command **<MNA**\n\nDisplay name of current map.",
		},
		{
			"should properly return argument definitions",
			"<TRA0021:0122:V010:0001",
			1,
			`Command **<TRAWWWW:XXXX:YYYY:ZZZZ**

Travel to map W, run event X, and move the PC to coordinates Y:Z.

* 0021: EgEnd1 - Side Room
* 0122
* V010
* 0001`,
		},
	}

	c := config.New()

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got := tsc.GetHoverInfo(testcase.str, testcase.at, &c)
			if got != testcase.want {
				t.Errorf(
					"GetHoverInfo(%s, %d, c) got %v, want %v",
					testcase.str,
					testcase.at,
					got,
					testcase.want,
				)
			}
		})
	}
}
