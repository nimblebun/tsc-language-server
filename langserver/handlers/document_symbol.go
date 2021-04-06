package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"pkg.nimblebun.works/go-lsp"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem/filehandler"
	"pkg.nimblebun.works/tsc-language-server/tsc"
)

// TextDocumentSymbol is the callback that runs on the
// "textDocument/documentSymbol" method.
func (mh *MethodHandler) TextDocumentSymbol(ctx context.Context, req *jrpc2.Request) ([]lsp.DocumentSymbol, error) {
	var symbols []lsp.DocumentSymbol

	var params lsp.DocumentSymbolParams
	err := req.UnmarshalParams(jrpc2.NonStrict(&params))
	if err != nil {
		return symbols, err
	}

	fs, err := lsctx.FileSystem(ctx)
	if err != nil {
		return symbols, err
	}

	config, err := lsctx.Config(ctx)

	handler := filehandler.FromDocumentURI(params.TextDocument.URI)

	path, err := handler.FullPath()
	if err != nil {
		return symbols, err
	}

	data, err := fs.ReadFile(path)
	if err != nil {
		return symbols, err
	}

	contents := string(data)

	doc := lsp.TextDocumentItem{
		URI:        params.TextDocument.URI,
		LanguageID: "tsc",
		Text:       contents,
	}

	symbols = tsc.GetEventSymbols(contents, doc, config)

	return symbols, nil
}
