package validators_test

import "github.com/sourcegraph/go-lsp"

func dummyTextDocument(input string) lsp.TextDocumentItem {
	return lsp.TextDocumentItem{
		Text: input,
	}
}
