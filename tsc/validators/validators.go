package validators

import (
	"github.com/sourcegraph/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
)

// Validate will return a slice of diagnostics based on the inspection of
// commands, events, and messages.
func Validate(textDocumentItem lsp.TextDocumentItem, conf *config.Config) []lsp.Diagnostic {
	text := textDocumentItem.Text

	diagnostics := make([]lsp.Diagnostic, 0)
	diagnostics = append(diagnostics, ValidateCommands(text, textDocumentItem, conf)...)
	diagnostics = append(diagnostics, ValidateEvents(text, textDocumentItem)...)
	diagnostics = append(diagnostics, ValidateMessages(text, textDocumentItem, conf)...)
	return diagnostics
}
