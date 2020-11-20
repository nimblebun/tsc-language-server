package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
)

// TSCSetConfig is the callback that runs on the "tsc/setConfig" method.
func (mh *MethodHandler) TSCSetConfig(ctx context.Context, req *jrpc2.Request) error {
	conf, err := lsctx.Config(ctx)
	if err != nil {
		return err
	}

	userConfig := []byte(req.ParamString())
	conf.Update(userConfig)

	return nil
}
