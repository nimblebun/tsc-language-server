package validators_test

import "pkg.nimblebun.works/go-lsp"

func dummyTextDocument(input string) lsp.TextDocumentItem {
	return lsp.TextDocumentItem{
		Text: input,
	}
}
