package validators

import (
	"fmt"
	"regexp"

	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/langserver/textdocument"
	"pkg.nimblebun.works/tsc-language-server/tsc"
	"pkg.nimblebun.works/tsc-language-server/utils"
)

// ValidateCommands will ensure that the arguments provided to a command are
// correct. Current criteria include:
//
// - Number of arguments must be the same as the number defined in .tscrc.json
func ValidateCommands(text string, textDocumentItem lsp.TextDocumentItem, conf *config.Config) []lsp.Diagnostic {
	document := textdocument.From(textDocumentItem)

	// this will match <ABC, <ABC0000, <ABC0000:0001, <ABC0000:0001:0002, <ABC0000:0001:0002:0003,
	// <ABC0000?0001, <ABC0000 0001b0002, <ABC0000$0001\0002^0003, <FAOV100, <FLJV020:V102
	re := regexp.MustCompile("<(([A-Z0-9+-]){3}(([0-9V])([0-9]){3})?)((.([0-9V])([0-9]){3})?){0,3}")

	diagnostics := []lsp.Diagnostic{}

	for _, match := range re.FindAllStringIndex(text, -1) {
		from, to := match[0], match[1]
		input := text[from:to]

		targetCommand := utils.Substring(input, 0, 4)
		command, found := conf.GetTSCDefinition(targetCommand)

		if !found {
			continue
		}

		argc := 0
		inputWithoutCommand := input[4:]

		for i := 0; i < len(inputWithoutCommand); i++ {
			arg := utils.Substring(inputWithoutCommand, i*5, 4)

			if tsc.IsValidArgument(arg) {
				argc++
			}
		}

		if argc != command.Nargs() {
			quantity := "few"
			if argc > command.Nargs() {
				quantity = "many"
			}

			diagnostic := lsp.Diagnostic{
				Severity: lsp.DSError,
				Range: lsp.Range{
					Start: document.PositionAt(from),
					End:   document.PositionAt(to),
				},
				Message: fmt.Sprintf(
					"Too %s arguments provided to %s. Expected %d, got %d.",
					quantity,
					command.Label,
					command.Nargs(),
					argc,
				),
				Source: "tsc-argc",
			}

			diagnostics = append(diagnostics, diagnostic)
		}
	}

	return diagnostics
}
