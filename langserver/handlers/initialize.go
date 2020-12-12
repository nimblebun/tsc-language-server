package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"pkg.nimblebun.works/go-lsp"
)

// Initialize is the callback that runs on the "initialize" method
func (mh *MethodHandler) Initialize(ctx context.Context, _ *jrpc2.Request) (lsp.InitializeResult, error) {
	result := lsp.InitializeResult{
		Capabilities: lsp.ServerCapabilities{
			TextDocumentSync: &lsp.TextDocumentSyncOptions{
				OpenClose: true,
				Change:    lsp.TDSyncKindIncremental,
			},
			CompletionProvider: &lsp.CompletionOptions{
				ResolveProvider: false,
			},
			HoverProvider:        &lsp.HoverOptions{},
			FoldingRangeProvider: &lsp.FoldingRangeRegistrationOptions{},
		},
	}

	return result, nil
}
