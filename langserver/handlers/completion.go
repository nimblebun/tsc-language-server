package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
)

// TextDocumentCompletion is the callback that runs on the
// "textDocument/completion" method
func TextDocumentCompletion(_ context.Context, _ *jrpc2.Request) error {
	// TODO: implement autocomplete. This is just a temporary workaround so it
	// doesn't yell all the time :)
	return nil
}
