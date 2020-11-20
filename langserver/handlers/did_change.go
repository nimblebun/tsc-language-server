package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/sourcegraph/go-lsp"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem/filehandler"
	"pkg.nimblebun.works/tsc-language-server/tsc/validators"
)

// TextDocumentDidChange is the callback that runs on the
// "textDocument/didChange" method
func (mh *MethodHandler) TextDocumentDidChange(ctx context.Context, req *jrpc2.Request) error {
	var params lsp.DidChangeTextDocumentParams
	err := req.UnmarshalParams(jrpc2.NonStrict(&params))
	if err != nil {
		return err
	}

	fs, err := lsctx.FileSystem(ctx)
	if err != nil {
		return err
	}

	handler := filehandler.VersionedFileHandlerFromDocument(params.TextDocument)
	path, err := handler.FullPath()
	if err != nil {
		return err
	}

	changes := params.ContentChanges
	err = fs.Change(handler, changes)
	if err != nil {
		return err
	}

	contents, err := fs.ReadFile(path)

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

	doc := lsp.TextDocumentItem{
		URI:        params.TextDocument.URI,
		LanguageID: "tsc",
		Version:    params.TextDocument.Version,
		Text:       string(contents),
	}

	results := validators.Validate(doc, conf)
	diags.Diagnose(ctx, params.TextDocument.URI, results)

	return nil
}
