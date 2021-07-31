package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"pkg.nimblebun.works/go-lsp"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem"
	"pkg.nimblebun.works/tsc-language-server/tsc/validators"
)

// TextDocumentDidOpen is the callback that runs on the "textDocument/didOpen"
// method
func (mh *MethodHandler) TextDocumentDidOpen(ctx context.Context, req *jrpc2.Request) error {
	var params lsp.DidOpenTextDocumentParams
	err := req.UnmarshalParams(&params)
	if err != nil {
		return err
	}

	fs, err := lsctx.FileSystem(ctx)
	if err != nil {
		return err
	}

	fh := filesystem.FileFromDocumentItem(params.TextDocument)
	err = fs.Create(*fh.Handler, fh.Text)
	if err != nil {
		return err
	}

	diags, err := lsctx.Diagnostics(ctx)
	if err != nil {
		return err
	}

	conf, err := lsctx.Config(ctx)
	if err != nil {
		return err
	}

	results := validators.Validate(params.TextDocument, conf)
	diags.Diagnose(ctx, params.TextDocument.URI, results)

	return nil
}
