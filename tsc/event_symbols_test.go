package tsc_test

import (
	"testing"

	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/tsc"
)

func TestGetEventSymbols(t *testing.T) {
	const text = `#0090fadein_from_left
<MNA<CMU0000<FAI0000<END
#0091fadein_from_up
<MNA<CMU0000<FAI0001<END
#0092fadein_from_right
<MNA<CMU0000<FAI0002<END
#0093fadein_from_down
<MNA<CMU0000<FAI0003<END
#0094fadein_from_center
<MNA<CMU0000<FAI0004<END`

	document := lsp.TextDocumentItem{
		Text: text,
	}

	conf := config.New()

	symbols := tsc.GetEventSymbols(text, document, &conf)

	t.Run("should return correct number of symbols", func(t *testing.T) {
		if len(symbols) != 5 {
			t.Errorf(
				"tsc.GetEventSymbols(text, document) length got %v, want %v",
				len(symbols),
				5,
			)
		}
	})

	t.Run("should return correct event names and locations", func(t *testing.T) {
		data := []lsp.DocumentSymbol{
			{
				Name: "#0090fadein_from_left",
				Range: &lsp.Range{
					Start: lsp.Position{
						Line:      0,
						Character: 0,
					},
					End: lsp.Position{
						Line:      0,
						Character: 21,
					},
				},
			},

			{
				Name: "#0091fadein_from_up",
				Range: &lsp.Range{
					Start: lsp.Position{
						Line:      2,
						Character: 0,
					},
					End: lsp.Position{
						Line:      2,
						Character: 19,
					},
				},
			},

			{
				Name: "#0092fadein_from_right",
				Range: &lsp.Range{
					Start: lsp.Position{
						Line:      4,
						Character: 0,
					},
					End: lsp.Position{
						Line:      4,
						Character: 22,
					},
				},
			},

			{
				Name: "#0093fadein_from_down",
				Range: &lsp.Range{
					Start: lsp.Position{
						Line:      6,
						Character: 0,
					},
					End: lsp.Position{
						Line:      6,
						Character: 21,
					},
				},
			},

			{
				Name: "#0094fadein_from_center",
				Range: &lsp.Range{
					Start: lsp.Position{
						Line:      8,
						Character: 0,
					},
					End: lsp.Position{
						Line:      8,
						Character: 23,
					},
				},
			},
		}

		for idx, symbol := range symbols {
			expected := data[idx]

			if symbol.Name != expected.Name {
				t.Errorf(
					"tsc.GetEventSymbols(text, document) @ %d -> name: got %v, want %v",
					idx,
					symbol.Name,
					expected.Name,
				)
			}

			if symbol.Range.Start.Line != expected.Range.Start.Line {
				t.Errorf(
					"tsc.GetEventSymbols(text, document) @ %d -> range start line: got %v, want %v",
					idx,
					symbol.Range.Start.Line,
					expected.Range.Start.Line,
				)
			}

			if symbol.Range.Start.Character != expected.Range.Start.Character {
				t.Errorf(
					"tsc.GetEventSymbols(text, document) @ %d -> range start character: got %v, want %v",
					idx,
					symbol.Range.Start.Character,
					expected.Range.Start.Character,
				)
			}

			if symbol.Range.End.Line != expected.Range.End.Line {
				t.Errorf(
					"tsc.GetEventSymbols(text, document) @ %d -> range end line: got %v, want %v",
					idx,
					symbol.Range.End.Line,
					expected.Range.End.Line,
				)
			}

			if symbol.Range.End.Character != expected.Range.End.Character {
				t.Errorf(
					"tsc.GetEventSymbols(text, document) @ %d -> range end character: got %v, want %v",
					idx,
					symbol.Range.End.Character,
					expected.Range.End.Character,
				)
			}
		}
	})
}
