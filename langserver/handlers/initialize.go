package handlers

import (
	"context"

	"github.com/creachadair/jrpc2"
	"github.com/sourcegraph/go-lsp"
)

// ServerCapabilities is a temporary workaround for the missing folding range
// provider field in Sourcegraph's go-lsp package. We will work on an in-house
// version of the LSP implementation in Go in the future.
type ServerCapabilities struct {
	TextDocumentSync                 *lsp.TextDocumentSyncOptionsOrKind   `json:"textDocumentSync,omitempty"`
	HoverProvider                    bool                                 `json:"hoverProvider,omitempty"`
	FoldingRangeProvider             bool                                 `json:"foldingRangeProvider,omitempty"`
	CompletionProvider               *lsp.CompletionOptions               `json:"completionProvider,omitempty"`
	SignatureHelpProvider            *lsp.SignatureHelpOptions            `json:"signatureHelpProvider,omitempty"`
	DefinitionProvider               bool                                 `json:"definitionProvider,omitempty"`
	TypeDefinitionProvider           bool                                 `json:"typeDefinitionProvider,omitempty"`
	ReferencesProvider               bool                                 `json:"referencesProvider,omitempty"`
	DocumentHighlightProvider        bool                                 `json:"documentHighlightProvider,omitempty"`
	DocumentSymbolProvider           bool                                 `json:"documentSymbolProvider,omitempty"`
	WorkspaceSymbolProvider          bool                                 `json:"workspaceSymbolProvider,omitempty"`
	ImplementationProvider           bool                                 `json:"implementationProvider,omitempty"`
	CodeActionProvider               bool                                 `json:"codeActionProvider,omitempty"`
	CodeLensProvider                 *lsp.CodeLensOptions                 `json:"codeLensProvider,omitempty"`
	DocumentFormattingProvider       bool                                 `json:"documentFormattingProvider,omitempty"`
	DocumentRangeFormattingProvider  bool                                 `json:"documentRangeFormattingProvider,omitempty"`
	DocumentOnTypeFormattingProvider *lsp.DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`
	RenameProvider                   bool                                 `json:"renameProvider,omitempty"`
	ExecuteCommandProvider           *lsp.ExecuteCommandOptions           `json:"executeCommandProvider,omitempty"`
	SemanticHighlighting             *lsp.SemanticHighlightingOptions     `json:"semanticHighlighting,omitempty"`
	XWorkspaceReferencesProvider     bool                                 `json:"xworkspaceReferencesProvider,omitempty"`
	XDefinitionProvider              bool                                 `json:"xdefinitionProvider,omitempty"`
	XWorkspaceSymbolByProperties     bool                                 `json:"xworkspaceSymbolByProperties,omitempty"`

	Experimental interface{} `json:"experimental,omitempty"`
}

// InitializeResult is a temporary workaround for the missing folding range
// provider field in Sourcegraph's go-lsp package. We will work on an in-house
// version of the LSP implementation in Go in the future.
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
}

// Initialize is the callback that runs on the "initialize" method
func (mh *MethodHandler) Initialize(ctx context.Context, _ *jrpc2.Request) (InitializeResult, error) {
	result := InitializeResult{
		Capabilities: ServerCapabilities{
			TextDocumentSync: &lsp.TextDocumentSyncOptionsOrKind{
				Options: &lsp.TextDocumentSyncOptions{
					OpenClose: true,
					Change:    lsp.TDSKIncremental,
				},
			},
			CompletionProvider: &lsp.CompletionOptions{
				ResolveProvider: false,
			},
			HoverProvider:        true,
			FoldingRangeProvider: true,
		},
	}

	return result, nil
}
