package tsc

import (
	"github.com/sourcegraph/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/langserver/textdocument"
	"pkg.nimblebun.works/tsc-language-server/utils"
)

const (
	// FRKComment resolves to the "comment" folding range kind.
	FRKComment = "comment"

	// FRKImports resolves to the "imports" folding range kind.
	FRKImports = "imports"

	// FRKRegion resolves to the "region" folding range kind.
	FRKRegion = "region"
)

// FoldingRange defines a LSP-compatible folding range.
type FoldingRange struct {
	StartLine      int `json:"startLine"`
	StartCharacter int `json:"startCharacter,omitempty"`
	EndLine        int `json:"endLine"`
	EndCharacter   int `json:"endCharacter"`

	Kind string `json:"kind,omitempty"`
}

func createFoldingRange(start int, end int, document textdocument.TextDocument) FoldingRange {
	startPos := document.PositionAt(start)
	endPos := document.PositionAt(end)

	return FoldingRange{
		StartLine:      startPos.Line,
		StartCharacter: startPos.Character,

		EndLine:      endPos.Line,
		EndCharacter: endPos.Character,

		Kind: FRKRegion,
	}
}

// GetFoldingRanges will return all foldable ranges from a given document.
func GetFoldingRanges(textDocumentItem lsp.TextDocumentItem) []FoldingRange {
	text := textDocumentItem.Text

	document := textdocument.From(textDocumentItem)
	ranges := make([]FoldingRange, 0)

	start := -1
	end := -1

	for idx, letter := range text {
		if letter == '#' && IsEvent(utils.Substring(text, idx, 5)) {
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
