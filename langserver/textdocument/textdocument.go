// Package textdocument is a slim wrapper around lsp.TextDocumentItem that
// contains useful methods such as PositionAt
//
// See: https://github.com/microsoft/vscode-languageserver-node/blob/master/textDocument/src/main.ts
package textdocument

import (
	"math"

	"pkg.nimblebun.works/go-lsp"
)

// TextDocument is a struct that extends the lsp.TextDocumentItem struct with
// useful helper methods
type TextDocument struct {
	lsp.TextDocumentItem
}

func (doc *TextDocument) computeLineOffsets(text string, isAtLineStart bool, textOffset int) []int {
	result := []int{}

	if isAtLineStart {
		result = append(result, textOffset)
	}

	for i := 0; i < len(text); i++ {
		ch := text[i]

		if ch == '\r' || ch == '\n' {
			if ch == '\r' && i+1 < len(text) && text[i+1] == '\n' {
				i++
			}
			result = append(result, textOffset+i+1)
		}
	}

	return result
}

// PositionAt will return a valid position from a given offset by calculating it
// relative to the line offsets of the document
func (doc *TextDocument) PositionAt(offset int) lsp.Position {
	offset = int(math.Max(math.Min(float64(offset), float64(len(doc.Text))), 0))

	lineOffsets := doc.computeLineOffsets(doc.Text, true, 0)
	low := 0
	high := len(lineOffsets)

	if high == 0 {
		return lsp.Position{
			Line:      0,
			Character: offset,
		}
	}

	for low < high {
		mid := (low + high) / 2
		if lineOffsets[mid] > offset {
			high = mid
		} else {
			low = mid + 1
		}
	}

	line := low - 1
	return lsp.Position{
		Line:      line,
		Character: offset - lineOffsets[line],
	}
}

// From will turn an lsp.TextDocumentItem into a textdocument.TextDocument
func From(document lsp.TextDocumentItem) TextDocument {
	return TextDocument{
		document,
	}
}
