package config

// MaxMessageLineLength specifies the maximum length of a line of text
// with and without portait image
type MaxMessageLineLength struct {
	Plain    int `json:"plain,omitempty"`
	Portrait int `json:"portrait,omitempty"`
}

// SetupConfig contains options for the TSC diagnostics
type SetupConfig struct {
	MaxMessageLineLength `json:"maxMessageLineLength,omitempty"`
}

// TSCDefinition is the definition of a TSC command
type TSCDefinition struct {
	Key           string   `json:"key"`
	Label         string   `json:"label"`
	Detail        string   `json:"detail"`
	Documentation string   `json:"documentation"`
	Format        string   `json:"format"`
	InsertText    string   `json:"insertText"`
	Nargs         int      `json:"nargs"`
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
	Equippables   GenericConfig `json:"equippables,omitempty"`
	Illustrations GenericConfig `json:"illustrations,omitempty"`
	Songs         GenericConfig `json:"songs,omitempty"`
	SFX           GenericConfig `json:"sfx,omitempty"`
	Custom        CustomConfig  `json:"custom,omitempty"`
}
