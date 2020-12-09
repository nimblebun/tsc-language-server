package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/sourcegraph/go-lsp"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem/filehandler"
	"pkg.nimblebun.works/tsc-language-server/tsc"
)

// FoldingRangeParams is a (loosely reproduced) structure that contains the
// fields sent on a textDocument/foldingRange request.
type FoldingRangeParams struct {
	TextDocument lsp.TextDocumentIdentifier `json:"textDocument"`
}

// TextDocumentFoldingRange is the callback that runs on the
// "textDocument/foldingRange" method.
func (mh *MethodHandler) TextDocumentFoldingRange(ctx context.Context, req *jrpc2.Request) ([]tsc.FoldingRange, error) {
	var params FoldingRangeParams
	err := req.UnmarshalParams(jrpc2.NonStrict(&params))
	if err != nil {
		return []tsc.FoldingRange{}, err
	}

	fs, err := lsctx.FileSystem(ctx)
	if err != nil {
		return []tsc.FoldingRange{}, err
	}

	handler := filehandler.FromDocumentURI(params.TextDocument.URI)
	path, err := handler.FullPath()
	if err != nil {
		return []tsc.FoldingRange{}, err
	}

	contents, err := fs.ReadFile(path)
	if err != nil {
		return []tsc.FoldingRange{}, err
	}

	doc := lsp.TextDocumentItem{
		URI:        params.TextDocument.URI,
		LanguageID: "tsc",
		Text:       string(contents),
	}

	ranges := tsc.GetFoldingRanges(doc)
	return ranges, nil
}
