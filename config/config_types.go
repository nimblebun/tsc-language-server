package config

import (
	"fmt"
	"strings"
)

// MaxMessageLineLength specifies the maximum length of a line of text
// with and without portait image
type MaxMessageLineLength struct {
	Plain    int `json:"plain,omitempty"`
	Portrait int `json:"portrait,omitempty"`
}

// LooseChecking specifies whether the language server should perform checks on
// event IDs and arguments in a loosely fashion. For more information, see:
// https://docs.nimblebun.works/tscrc-json#setup-loose-checking
type LooseChecking struct {
	Events    bool `json:"events,omitempty"`
	Arguments bool `json:"arguments,omitempty"`
}

// SetupConfig contains options for the TSC diagnostics
type SetupConfig struct {
	MaxMessageLineLength MaxMessageLineLength `json:"maxMessageLineLength,omitempty"`
	LooseChecking        LooseChecking        `json:"looseChecking,omitempty"`
}

// TSCDefinition is the definition of a TSC command
type TSCDefinition struct {
	Label         string   `json:"label"`
	Detail        string   `json:"detail"`
	Documentation string   `json:"documentation"`
	Format        string   `json:"format"`
	InsertText    string   `json:"insertText"`
	ArgType       []string `json:"argtype,omitempty"`
}

// TSCConfig is a map that resolves into a TSC definition
type TSCConfig map[string]TSCDefinition

// GenericConfig is a map that resolves into a string
type GenericConfig map[string]string

// CustomConfig is a map that resolves into a GenericConfig
type CustomConfig map[string]GenericConfig

// Config is a tscrc configuration file
type Config struct {
	Setup         SetupConfig   `json:"setup,omitempty"`
	TSC           TSCConfig     `json:"tsc,omitempty"`
	Directions    GenericConfig `json:"directions,omitempty"`
	Faces         []string      `json:"faces,omitempty"`
	Maps          GenericConfig `json:"maps,omitempty"`
	Weapons       GenericConfig `json:"weapons,omitempty"`
	Items         GenericConfig `json:"items,omitempty"`
	Equipables    GenericConfig `json:"equipables,omitempty"`
	Illustrations GenericConfig `json:"illustrations,omitempty"`
	Songs         GenericConfig `json:"songs,omitempty"`
	SFX           GenericConfig `json:"sfx,omitempty"`
	Custom        CustomConfig  `json:"custom,omitempty"`
}

// Nargs will return the number of arguments a TSC definition
func (definition *TSCDefinition) Nargs() int {
	return len(definition.ArgType)
}

// GetInsertText will retrieve a text to be inserted into the document. If the
// insertText property is provided in .tscrc.json, it will just return that.
// Otherwise, it will generate a zero-filled argument list.
func (definition *TSCDefinition) GetInsertText() string {
	if len(definition.InsertText) != 0 {
		return definition.InsertText
	}

	if len(definition.Label) == 0 {
		return ""
	}

	command := definition.Label[1:]

	if definition.Nargs() == 0 {
		return command
	}

	placeholders := make([]string, definition.Nargs())

	for idx := range definition.ArgType {
		placeholders[idx] = fmt.Sprintf("${%d:0000}", idx+1)
	}

	return fmt.Sprintf("%s%s", command, strings.Join(placeholders, ":"))
}
