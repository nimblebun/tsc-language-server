package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/sourcegraph/go-lsp"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem/filehandler"
)

// TextDocumentDidClose is the callback that runs on the "textDocument/didClose"
// method
func (mh *MethodHandler) TextDocumentDidClose(ctx context.Context, req *jrpc2.Request) error {
	var params lsp.DidCloseTextDocumentParams
	err := req.UnmarshalParams(jrpc2.NonStrict(&params))
	if err != nil {
		return err
	}

	fs, err := lsctx.FileSystem(ctx)
	if err != nil {
		return err
	}

	handler := filehandler.FromDocumentURI(params.TextDocument.URI)
	return fs.Remove(handler)
}
