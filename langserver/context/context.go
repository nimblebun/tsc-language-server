package context

import (
	"context"
	"fmt"

	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/langserver/diagnostics"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem"
)

type contextKey struct {
	Name string
}

func (ck *contextKey) String() string {
	return ck.Name
}

var (
	ctxConfig      = &contextKey{"config"}
	ctxDiagnostics = &contextKey{"diagnostics"}
	ctxFs          = &contextKey{"filesystem"}
)

// WithFileSystem will create a context that contains a file system for working
// with files.
func WithFileSystem(ctx context.Context, fs *filesystem.FileSystem) context.Context {
	return context.WithValue(ctx, ctxFs, fs)
}

// FileSystem will retrieve the FileSystem object from the provided context.
func FileSystem(ctx context.Context) (*filesystem.FileSystem, error) {
	fs, ok := ctx.Value(ctxFs).(*filesystem.FileSystem)
	if !ok {
		return nil, fmt.Errorf("missing context: %s", ctxFs)
	}

	return fs, nil
}

// WithConfig will create a context that contains a .tscrc.json configuration.
func WithConfig(ctx context.Context, conf *config.Config) context.Context {
	return context.WithValue(ctx, ctxConfig, conf)
}

// Config will retrieve the .tscrc.json configuration from the provided context.
func Config(ctx context.Context) (*config.Config, error) {
	conf, ok := ctx.Value(ctxConfig).(*config.Config)
	if !ok {
		return nil, fmt.Errorf("missing context: %s", ctxConfig)
	}

	return conf, nil
}

// WithDiagnostics will create a context that contains a diagnostics notifier.
func WithDiagnostics(ctx context.Context, diags *diagnostics.Notifier) context.Context {
	return context.WithValue(ctx, ctxDiagnostics, diags)
}

// Diagnostics will retrieve the diagnostics notifier from the provided context.
func Diagnostics(ctx context.Context) (*diagnostics.Notifier, error) {
	diags, ok := ctx.Value(ctxDiagnostics).(*diagnostics.Notifier)
	if !ok {
		return nil, fmt.Errorf("missing context: %s", ctxDiagnostics)
	}

	return diags, nil
}
