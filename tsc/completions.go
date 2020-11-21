package tsc

import (
	"github.com/sourcegraph/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
)

// GetCompletions will return a list of completions that can be used for a
// provided TSC command.
func GetCompletions(conf *config.Config) []lsp.CompletionItem {
	items := []lsp.CompletionItem{}

	definitions := conf.GetTSCDefinitions()

	for idx, def := range definitions {
		item := lsp.CompletionItem{
			Label:         def.Label,
			Detail:        def.Detail,
			Documentation: def.Documentation,
			Kind:          lsp.CIKFunction,
			Data:          idx,
			InsertText:    def.GetInsertText(),
		}

		if def.Nargs() > 0 {
			item.InsertTextFormat = lsp.ITFSnippet
		}

		items = append(items, item)
	}

	return items
}
