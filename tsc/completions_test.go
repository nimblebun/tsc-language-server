package tsc_test

import (
	"testing"

	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/tsc"
)

func TestGetCompletions(t *testing.T) {
	conf := config.New()
	completions := tsc.GetCompletions(&conf)

	definitionsLength := len(conf.GetTSCDefinitions())
	completionsLength := len(completions)

	if definitionsLength != completionsLength {
		t.Errorf("GetCompletions(&conf) -> TSC definition length %d doesn't match completions length %d", definitionsLength, completionsLength)
	}

	for _, val := range completions {
		definition, found := conf.GetTSCDefinition(val.Label)

		if !found {
			t.Errorf("GetCompletions(&conf) @ %s definition not found.", val.Label)
		}

		if val.Detail != definition.Detail {
			t.Errorf("GetCompletions(&conf) @ %s (detail) got %v, want %v", val.Label, val.Detail, definition.Detail)
		}

		if val.Documentation.Value != definition.Documentation {
			t.Errorf("GetCompletions(&conf) @ %s (documentation) got %v, want %v", val.Label, val.Documentation, definition.Documentation)
		}

		if val.InsertText != definition.GetInsertText() {
			t.Errorf("GetCompletions(&conf) @ %s (inserttext) got %v, want %v", val.Label, val.InsertText, definition.GetInsertText())
		}

		if val.Kind != lsp.CIKFunction {
			t.Errorf("GetCompletions(&conf) @ %s (kind) got %v, want %v", val.Label, val.Kind.String(), lsp.CIKFunction.String())
		}

		if definition.Nargs() > 0 && val.InsertTextFormat != lsp.ITFSnippet {
			t.Errorf("GetCompletions(&conf) @ %s *(inserttextformat) got %v, want %v", val.Label, val.InsertTextFormat, lsp.ITFSnippet)
		}
	}
}
