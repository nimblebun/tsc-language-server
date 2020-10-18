package handlers

import (
	"context"

	"github.com/sourcegraph/go-lsp"
)

// Initialized is the callback that runs on the "initialized" method
func Initialized(ctx context.Context, params lsp.None) error {
	return nil
}
