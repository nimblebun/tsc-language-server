package diagnostics

import (
	"context"
	"log"
	"sync"

	"github.com/creachadair/jrpc2"
	"github.com/sourcegraph/go-lsp"
)

// DocumentContext contains data about the document on which we have performed
// the diagnostics on.
type DocumentContext struct {
	Ctx         context.Context
	URI         lsp.DocumentURI
	Diagnostics []lsp.Diagnostic
}

// Notifier contains a channel representing TSC documents with diagnostics, as
// well as methods for publishing new diagnostics.
type Notifier struct {
	SessionCtx       context.Context
	TSCDocs          chan DocumentContext
	CloseTSCDocsOnce sync.Once
}

func publishDiagnostics(docs <-chan DocumentContext, logger *log.Logger) {
	for doc := range docs {
		if err := jrpc2.PushNotify(doc.Ctx, "textDocument/publishDiagnostics", lsp.PublishDiagnosticsParams{
			URI:         doc.URI,
			Diagnostics: doc.Diagnostics,
		}); err != nil {
			logger.Printf("Failed to publish diagnostics: %s", err)
		}
	}
}

// NewNotifier creates a new Notifier instance for running a goroutine that
// sends out diagnostics to all documents in the channel.
func NewNotifier(sessionCtx context.Context, logger *log.Logger) *Notifier {
	TSCDocs := make(chan DocumentContext, 50)
	go publishDiagnostics(TSCDocs, logger)
	return &Notifier{
		TSCDocs:    TSCDocs,
		SessionCtx: sessionCtx,
	}
}

// Diagnose will enqueue a document and a slice of diagnostics for publishing.
// Notifications will be delivered in the order of the queue.
func (n *Notifier) Diagnose(ctx context.Context, uri lsp.DocumentURI, diagnostics []lsp.Diagnostic) {
	select {
	case <-n.SessionCtx.Done():
		n.CloseTSCDocsOnce.Do(func() {
			close(n.TSCDocs)
		})
		return
	default:
	}
	n.TSCDocs <- DocumentContext{
		Ctx:         ctx,
		URI:         uri,
		Diagnostics: diagnostics,
	}
}
