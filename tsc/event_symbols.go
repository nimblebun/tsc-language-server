package tsc

import (
	"regexp"

	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/langserver/textdocument"
)

// GetEventSymbols will return a list of events as LSP-compatible symbols.
func GetEventSymbols(text string, textDocumentItem lsp.TextDocumentItem) []lsp.DocumentSymbol {
	document := textdocument.From(textDocumentItem)

	// this will match #0000 until the end of the line
	re := regexp.MustCompile("#(?:[0-9]{4}).*")

	symbols := make([]lsp.DocumentSymbol, 0)

	for _, match := range re.FindAllStringIndex(text, -1) {
		from, to := match[0], match[1]
		eventDeclaration := text[from:to]

		symbolRange := &lsp.Range{
			Start: document.PositionAt(from),
			End:   document.PositionAt(to),
		}

		symbol := lsp.DocumentSymbol{
			Name:           eventDeclaration,
			Kind:           lsp.SKEvent,
			Range:          symbolRange,
			SelectionRange: symbolRange,
		}

		symbols = append(symbols, symbol)
	}

	return symbols
}
