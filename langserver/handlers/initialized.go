package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
)

// Initialized is the callback that runs on the "initialized" method
func Initialized(_ context.Context, _ *jrpc2.Request) error {
	return nil
}
