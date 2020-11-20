package tsc

import (
	"testing"

	"github.com/sourcegraph/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
)

func TestGetCompletions(t *testing.T) {
	conf := config.New()
	completions := GetCompletions(&conf)

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

		if val.Documentation != definition.Documentation {
			t.Errorf("GetCompletions(&conf) @ %s (documentation) got %v, want %v", val.Label, val.Documentation, definition.Documentation)
		}

		if val.InsertText != definition.InsertText {
			t.Errorf("GetCompletions(&conf) @ %s (inserttext) got %v, want %v", val.Label, val.InsertText, definition.InsertText)
		}

		if val.Kind != lsp.CIKFunction {
			t.Errorf("GetCompletions(&conf) @ %s (kind) got %v, want %v", val.Label, val.Kind.String(), lsp.CIKFunction.String())
		}

		if definition.Nargs > 0 && val.InsertTextFormat != lsp.ITFSnippet {
			t.Errorf("GetCompletions(&conf) @ %s *(inserttextformat) got %v, want %v", val.Label, val.InsertTextFormat, lsp.ITFSnippet)
		}
	}
}
