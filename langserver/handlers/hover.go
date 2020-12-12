package handlers

import (
	"context"

	"pkg.nimblebun.works/go-lsp"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem/filehandler"
	"pkg.nimblebun.works/tsc-language-server/tsc"
	"pkg.nimblebun.works/tsc-language-server/utils"
)

// TextDocumentHover is the callback that runs on the "textDocument/hover"
// method.
func (mh *MethodHandler) TextDocumentHover(ctx context.Context, params lsp.TextDocumentPositionParams) (lsp.Hover, error) {
	var result lsp.Hover

	fs, err := lsctx.FileSystem(ctx)
	if err != nil {
		return result, err
	}

	config, err := lsctx.Config(ctx)
	if err != nil {
		return result, err
	}

	handler := filehandler.FromDocumentURI(params.TextDocument.URI)

	path, err := handler.FullPath()
	if err != nil {
		return result, err
	}

	data, err := fs.ReadFile(path)
	if err != nil {
		return result, err
	}

	start := lsp.Position{
		Line:      params.Position.Line,
		Character: 0,
	}

	end := lsp.Position{
		Line:      params.Position.Line + 1,
		Character: 0,
	}

	contents := string(data)
	startOffset := utils.DocumentOffset(data, start)
	endOffset := utils.DocumentOffset(data, end)

	if endOffset == -1 {
		// FIXME: last line cannot be interpreted.
		endOffset = len(contents) - 1
	}

	text := utils.Substring(contents, startOffset, endOffset-startOffset)

	if len(text) == 0 {
		return result, nil
	}

	offset := utils.DocumentOffset(data, params.Position)

	info := tsc.GetHoverInfo(text, offset-startOffset, config)

	if len(info) > 0 {
		result.Contents = lsp.MarkupContent{
			Kind:  lsp.MKMarkdown,
			Value: info,
		}
	}

	return result, nil
}
