package handlers

import (
	"context"

	"github.com/sourcegraph/go-lsp"
)

// Initialize is the callback that runs on the "initialize" method
func (mh *MethodHandler) Initialize(ctx context.Context, params lsp.InitializeParams) (lsp.InitializeResult, error) {
	result := lsp.InitializeResult{
		Capabilities: lsp.ServerCapabilities{
			TextDocumentSync: &lsp.TextDocumentSyncOptionsOrKind{
				Options: &lsp.TextDocumentSyncOptions{
					OpenClose: true,
					Change:    lsp.TDSKIncremental,
				},
			},
			CompletionProvider: &lsp.CompletionOptions{
				ResolveProvider: false,
			},
			HoverProvider: true,
		},
	}

	return result, nil
}
