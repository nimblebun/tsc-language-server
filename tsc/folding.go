package tsc

import (
	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/langserver/textdocument"
	"pkg.nimblebun.works/tsc-language-server/utils"
)

func createFoldingRange(start int, end int, document textdocument.TextDocument) lsp.FoldingRange {
	startPos := document.PositionAt(start)
	endPos := document.PositionAt(end)

	return lsp.FoldingRange{
		StartLine:      startPos.Line,
		StartCharacter: startPos.Character,

		EndLine:      endPos.Line,
		EndCharacter: endPos.Character,

		Kind: lsp.FoldingRangeKindRegion,
	}
}

// GetFoldingRanges will return all foldable ranges from a given document.
func GetFoldingRanges(textDocumentItem lsp.TextDocumentItem, conf *config.Config) []lsp.FoldingRange {
	text := textDocumentItem.Text

	document := textdocument.From(textDocumentItem)
	ranges := make([]lsp.FoldingRange, 0)

	start := -1
	end := -1

	for idx, letter := range text {
		if letter == '#' && IsEvent(utils.Substring(text, idx, 5), conf.Setup.LooseChecking.Events) {
			if end-start > 4 {
				foldingRange := createFoldingRange(start, end, document)
				ranges = append(ranges, foldingRange)
			}

			start = idx
		}

		if letter != '\n' && letter != '\r' {
			end = idx
		}
	}

	if end-start > 4 {
		foldingRange := createFoldingRange(start, end, document)
		ranges = append(ranges, foldingRange)
	}

	return ranges
}
