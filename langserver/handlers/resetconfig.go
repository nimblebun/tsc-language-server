package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
)

// TSCResetConfig is the callback that runs on the "tsc/setConfig" method.
func (mh *MethodHandler) TSCResetConfig(ctx context.Context, _ *jrpc2.Request) error {
	conf, err := lsctx.Config(ctx)
	if err != nil {
		return err
	}

	conf.Reset()

	return nil
}
