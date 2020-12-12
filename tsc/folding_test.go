package tsc_test

import (
	"testing"

	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/tsc"
)

func TestGetFoldingRanges(t *testing.T) {
	dummyData := `#0098
<CNP0306:0117:0000<ANP0306:0032:0002<END

#0100
<KEY<FLJ0839:0101
<SOU0011<ANP0100:0000:0002
<FAO0000<TRA0046:0090:0017:0009

#0101
<KEY<MSGIt won't open...<NOD<END

#0200
<KEY
<FLJ0832:0204
<FLJ0824:0203
<FLJ0823:0202
<FLJ0821:0201<MSG
OPEN SHUTTER?<YNJ0000<CLR
OPENING SHUTTER<NOD<CLO<MYD0000
<WAI0030<ANP0250:0010:0001<WAI0010<ANP0300:0001:0002
<WAI0022<ANP0251:0010:0001<ANP0300:0003:0002
<WAI0032<ANP0252:0010:0001
<WAI0032<ANP0253:0010:0001
<WAI0032<ANP0254:0010:0001<DNP0250
<WAI0032<DNP0251
<WAI0032<DNP0252
<ANP0253:0001:0000<WAI0032<DNP0300
<CNP0301:0117:0000
<ANP0301:0021:0002
<FL-0820<FL+0821<FL+0822<MSGABNORMALITY DETECTED IN
SHUTTER NO. 4<NOD<END`

	dummyDocument := lsp.TextDocumentItem{
		Text: dummyData,
	}

	t.Run("should return correct number of folding ranges", func(t *testing.T) {
		ranges := tsc.GetFoldingRanges(dummyDocument)
		actual := len(ranges)
		expected := 4

		if expected != actual {
			t.Errorf("GetFoldingRanges(document) length, got %d, want %d", actual, expected)
		}
	})

	t.Run("should return correct ranges", func(t *testing.T) {
		ranges := tsc.GetFoldingRanges(dummyDocument)
		expectedRanges := []lsp.FoldingRange{
			{
				StartLine:      0,
				StartCharacter: 0,
				EndLine:        1,
				EndCharacter:   39,
			},
			{
				StartLine:      3,
				StartCharacter: 0,
				EndLine:        6,
				EndCharacter:   30,
			},
			{
				StartLine:      8,
				StartCharacter: 0,
				EndLine:        9,
				EndCharacter:   31,
			},
			{
				StartLine:      11,
				StartCharacter: 0,
				EndLine:        30,
				EndCharacter:   20,
			},
		}

		for idx := range ranges {
			actual := ranges[idx]
			expected := expectedRanges[idx]

			if actual.StartLine != expected.StartLine {
				t.Errorf("GetFoldingRanges(doc) @ %d -> StartLine, got %v, want %v", idx, actual.StartLine, expected.StartLine)
			}

			if actual.StartCharacter != expected.StartCharacter {
				t.Errorf("GetFoldingRanges(doc) @ %d -> StartCharacter, got %v, want %v", idx, actual.StartCharacter, expected.StartCharacter)
			}

			if actual.EndLine != expected.EndLine {
				t.Errorf("GetFoldingRanges(doc) @ %d -> EndLine, got %v, want %v", idx, actual.EndLine, expected.EndLine)
			}

			if actual.EndCharacter != expected.EndCharacter {
				t.Errorf("GetFoldingRanges(doc) @ %d -> EndCharacter, got %v, want %v", idx, actual.EndCharacter, expected.EndCharacter)
			}
		}
	})
}
