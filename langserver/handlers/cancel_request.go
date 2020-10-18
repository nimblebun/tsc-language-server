package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/sourcegraph/go-lsp"
)

// CancelRequest will be called on "$/cancelRequest"
func CancelRequest(ctx context.Context, params lsp.CancelParams) error {
	id := params.ID.Str
	jrpc2.CancelRequest(ctx, id)
	return nil
}
