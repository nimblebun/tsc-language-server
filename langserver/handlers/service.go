package handlers

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/code"
	"github.com/creachadair/jrpc2/handler"
	"pkg.nimblebun.works/tsc-language-server/config"
	lsctx "pkg.nimblebun.works/tsc-language-server/langserver/context"
	"pkg.nimblebun.works/tsc-language-server/langserver/diagnostics"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem"
	"pkg.nimblebun.works/tsc-language-server/langserver/session"
)

const requestCancelled code.Code = -32800

var discardLogs = log.New(ioutil.Discard, "", 0)

// Service is a handler service for sessions
type Service struct {
	logger *log.Logger

	ctx context.Context

	sessionCtx  context.Context
	stopSession context.CancelFunc
}

func toHandlerMap(m map[string]handler.Func) handler.Map {
	handlerMap := make(handler.Map, len(m))

	for method, callback := range m {
		handlerMap[method] = handler.New(callback)
	}

	return handlerMap
}

func handle(ctx context.Context, req *jrpc2.Request, cb interface{}) (interface{}, error) {
	f := handler.New(cb)
	result, err := f.Handle(ctx, req)

	if ctx.Err() != nil && errors.Is(ctx.Err(), context.Canceled) {
		err = fmt.Errorf("%w: %s", requestCancelled.Err(), err)
	}

	return result, err
}

// SetLogger overwrites the current logger of the service
func (service *Service) SetLogger(logger *log.Logger) {
	service.logger = logger
}

// Assigner will create the session and set up the service method/handler map
func (service *Service) Assigner() (jrpc2.Assigner, error) {
	service.logger.Println("Preparing new session...")

	sess := session.New(service.stopSession)
	mh := NewMethodHandler(service.logger)

	err := sess.Prepare()
	if err != nil {
		return nil, fmt.Errorf("Unable to prepare session: %w", err)
	}

	conf := config.New()
	diags := diagnostics.NewNotifier(service.sessionCtx, service.logger)

	fs := filesystem.New()
	fs.SetLogger(service.logger)

	serviceMap := map[string]handler.Func{
		"initialize": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.Init(req)
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithFileSystem(ctx, fs)

			return handle(ctx, req, mh.Initialize)
		},

		"initialized": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.FinishInitialization(req)
			if err != nil {
				return nil, err
			}

			return handle(ctx, req, Initialized)
		},

		"textDocument/didOpen": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithConfig(ctx, &conf)
			ctx = lsctx.WithDiagnostics(ctx, diags)
			ctx = lsctx.WithFileSystem(ctx, fs)

			return handle(ctx, req, mh.TextDocumentDidOpen)
		},

		"textDocument/didClose": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithFileSystem(ctx, fs)

			return handle(ctx, req, mh.TextDocumentDidClose)
		},

		"textDocument/didChange": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithConfig(ctx, &conf)
			ctx = lsctx.WithDiagnostics(ctx, diags)
			ctx = lsctx.WithFileSystem(ctx, fs)

			return handle(ctx, req, mh.TextDocumentDidChange)
		},

		"textDocument/completion": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithConfig(ctx, &conf)

			return handle(ctx, req, TextDocumentCompletion)
		},

		"textDocument/documentSymbol": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithFileSystem(ctx, fs)

			return handle(ctx, req, mh.TextDocumentSymbol)
		},

		"textDocument/foldingRange": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithFileSystem(ctx, fs)
			ctx = lsctx.WithConfig(ctx, &conf)

			return handle(ctx, req, mh.TextDocumentFoldingRange)
		},

		"textDocument/hover": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithConfig(ctx, &conf)
			ctx = lsctx.WithFileSystem(ctx, fs)

			return handle(ctx, req, mh.TextDocumentHover)
		},

		"tsc/setConfig": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithConfig(ctx, &conf)

			return handle(ctx, req, mh.TSCSetConfig)
		},

		"tsc/resetConfig": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithConfig(ctx, &conf)

			return handle(ctx, req, mh.TSCResetConfig)
		},

		"shutdown": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.Shutdown(req)
			if err != nil {
				return nil, err
			}

			return nil, nil
		},

		"exit": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.Exit()
			if err != nil {
				return nil, err
			}

			service.stopSession()

			return nil, nil
		},

		"$/cancelRequest": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := sess.EnsureInitialized()
			if err != nil {
				return nil, err
			}

			return handle(ctx, req, CancelRequest)
		},
	}

	return toHandlerMap(serviceMap), nil
}

// Finish will terminate the service
func (service *Service) Finish(status jrpc2.ServerStatus) {
	if status.Closed() || status.Err != nil {
		service.logger.Printf("Session stopped (err: %v)", status.Err)
	}

	service.stopSession()
}

// NewSession instantiates a new session
func NewSession(ctx context.Context) session.ServiceSession {
	sessionCtx, stopSession := context.WithCancel(ctx)

	return &Service{
		logger: discardLogs,

		ctx: ctx,

		sessionCtx:  sessionCtx,
		stopSession: stopSession,
	}
}
