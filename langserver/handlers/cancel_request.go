package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	lsp "pkg.nimblebun.works/go-lsp"
)

// CancelRequest will be called on "$/cancelRequest"
func CancelRequest(ctx context.Context, req *jrpc2.Request) error {
	var params lsp.CancelParams
	if err := req.UnmarshalParams(&params); err != nil {
		return err
	}

	jrpc2.ServerFromContext(ctx).CancelRequest(params.ID.String())
	return nil
}
