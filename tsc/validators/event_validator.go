package validators

import (
	"fmt"
	"regexp"

	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/langserver/textdocument"
	"pkg.nimblebun.works/tsc-language-server/utils"
)

// ValidateEvents will ensure that the defined events are correct. Currently it
// checks if an event was re-defined in the same file. Doesn't keep track of
// global events defined in Head.tsc
func ValidateEvents(text string, textDocumentItem lsp.TextDocumentItem, conf *config.Config) []lsp.Diagnostic {
	document := textdocument.From(textDocumentItem)

	// this will match #0000
	re := regexp.MustCompile("#(?:[0-9]{4})")

	if conf.Setup.LooseChecking.Events {
		// this will match #0000, #0ABC
		re = regexp.MustCompile("#(?:.{4})")
	}

	occurrences := make(map[string]int)
	diagnostics := []lsp.Diagnostic{}

	for _, match := range re.FindAllStringIndex(text, -1) {
		from, to := match[0], match[1]
		event := utils.Substring(text, from, 5)

		if occurrences[event] < 2 {
			occurrences[event]++
		}

		if occurrences[event] == 2 {
			diagnostic := lsp.Diagnostic{
				Severity: lsp.DSError,
				Range: lsp.Range{
					Start: document.PositionAt(from),
					End:   document.PositionAt(to),
				},
				Message: fmt.Sprintf("Event %s has already been declared in this file.", event),
				Source:  "no-event-dup",
			}

			diagnostics = append(diagnostics, diagnostic)
		}
	}

	return diagnostics
}
