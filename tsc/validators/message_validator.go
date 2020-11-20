package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sourcegraph/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/langserver/textdocument"
)

// ValidateMessages will warn if it notices any text overflow issues in the
// displayed messages (<MSG, <MS2, <MS3 commands). It takes into account the
// cases in which <FAC was previously defined. The limits are defined in the
// .tscrc.json configuration file
func ValidateMessages(text string, textDocumentItem lsp.TextDocumentItem, conf *config.Config) []lsp.Diagnostic {
	document := textdocument.From(textDocumentItem)

	// this will match anything that starts with <MSG, <MS2, or <MS3 and ends
	// <CLO, <END, <ESC, <EVE, <INI, <TRA, or <XX1
	re := regexp.MustCompile("(?:<MS[G|2|3]\n?)((.|\n)+?)(?:<CLO|<END|<ESC|<EVE|<INI|<TRA|<XX1)")

	diagnostics := []lsp.Diagnostic{}

	for _, match := range re.FindAllStringIndex(text, -1) {
		from, to := match[0], match[1]
		message := text[from:to]

		linesRe := regexp.MustCompile("\n|<CLR")
		lines := linesRe.Split(message, -1)

		hasPortrait := false

		for _, line := range lines {
			// this will remove all TSC tags from the string, giving us a clean output
			cleanLineRe := regexp.MustCompile("<(([A-Z0-9+-]){3}(([0-9]){4})?)((:([0-9]){4})?){0,3}")
			cleanLine := cleanLineRe.ReplaceAllString(line, "")

			if strings.Contains(line, "<FAC0000") {
				hasPortrait = false
			}

			if strings.Contains(line, "<FAC") && !strings.Contains(line, "<FAC0000") {
				hasPortrait = true
			}

			limit := conf.Setup.MaxMessageLineLength.Plain
			if hasPortrait {
				limit = conf.Setup.MaxMessageLineLength.Portrait
			}

			startPos := strings.Index(message, line)

			if len(cleanLine) > limit {
				diagnostic := lsp.Diagnostic{
					Severity: lsp.Warning,
					Range: lsp.Range{
						Start: document.PositionAt(from + startPos),
						End:   document.PositionAt(from + startPos + len(line)),
					},
					Message: fmt.Sprintf(
						"Message exceeds %d characters (current length: %d). This may cause text overflow issues.",
						limit,
						len(cleanLine),
					),
					Source: "text-overflow",
				}

				diagnostics = append(diagnostics, diagnostic)
			}
		}
	}

	return diagnostics
}
