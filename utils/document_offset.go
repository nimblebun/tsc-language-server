package utils

import "pkg.nimblebun.works/go-lsp"

// DocumentOffset is a utility function used to turn an LSP cursor position into
// an offset from the current file buffer.
func DocumentOffset(text []byte, position lsp.Position) int {
	column, line := 0, 0

	for offset, character := range text {
		if line == position.Line && column == position.Character {
			return offset
		}

		if character == '\r' && offset+1 < len(text) && text[offset+1] == '\n' {
			continue
		}

		if character == '\n' {
			line++
			column = 0
		} else {
			column++
		}
	}

	return -1
}
