package tsc

import (
	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/config"
)

// GetCompletions will return a list of completions that can be used for a
// provided TSC command.
func GetCompletions(conf *config.Config) []lsp.CompletionItem {
	items := make([]lsp.CompletionItem, 0)

	definitions := conf.GetTSCDefinitions()

	for idx, def := range definitions {
		item := lsp.CompletionItem{
			Label:  def.Label,
			Detail: def.Detail,
			Documentation: lsp.MarkupContent{
				Kind:  lsp.MKMarkdown,
				Value: def.Documentation,
			},
			Kind:       lsp.CIKFunction,
			Data:       idx,
			InsertText: def.GetInsertText(),
		}

		if def.Nargs() > 0 {
			item.InsertTextFormat = lsp.ITFSnippet
		}

		items = append(items, item)
	}

	return items
}
