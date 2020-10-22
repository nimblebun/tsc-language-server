package validators

import (
	"github.com/nimblebun/tsc-language-server/config"
	"github.com/sourcegraph/go-lsp"
)

// Validate will return a slice of diagnostics based on the inspection of
// commands, events, and messages.
func Validate(textDocumentItem lsp.TextDocumentItem, conf *config.Config) []lsp.Diagnostic {
	text := textDocumentItem.Text

	var diagnostics []lsp.Diagnostic
	diagnostics = append(diagnostics, ValidateCommands(text, textDocumentItem, conf)...)
	diagnostics = append(diagnostics, ValidateEvents(text, textDocumentItem)...)
	diagnostics = append(diagnostics, ValidateMessages(text, textDocumentItem, conf)...)
	return diagnostics
}
