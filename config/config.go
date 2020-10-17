package config

import (
	"encoding/json"
	"strconv"
	"strings"
)

func getGenericItem(id string, value string) string {
	if len(value) == 0 {
		return id
	}

	return value
}

// New creates a new TSC configuration
func New() Config {
	config := Config{}
	json.Unmarshal(DefaultConfig, &config)
	return config
}

// GetTSCDefinitions will return a list of TSC definitions
func (config *Config) GetTSCDefinitions() []TSCDefinition {
	definitions := make([]TSCDefinition, 0, len(config.TSC))

	for _, definition := range config.TSC {
		definitions = append(definitions, definition)
	}

	return definitions
}

// GetTSCDefinition will return a singular TSC definition
func (config *Config) GetTSCDefinition(key string) (TSCDefinition, bool) {
	definition, found := config.TSC[key]
	return definition, found
}

func (config *Config) getDirection(id string) string {
	return getGenericItem(id, config.Directions[id])
}

func (config *Config) getFace(id int) string {
	if len(config.Faces) < id {
		return strconv.Itoa(id)
	}

	return config.Faces[id]
}

func (config *Config) getMap(id string) string {
	return getGenericItem(id, config.Maps[id])
}

func (config *Config) getWeapon(id string) string {
	return getGenericItem(id, config.Weapons[id])
}

func (config *Config) getItem(id string) string {
	return getGenericItem(id, config.Items[id])
}

func (config *Config) getEquippable(id string) string {
	return getGenericItem(id, config.Equippables[id])
}

func (config *Config) getIllustration(id string) string {
	return getGenericItem(id, config.Illustrations[id])
}

func (config *Config) getSong(id string) string {
	return getGenericItem(id, config.Songs[id])
}

func (config *Config) getSFX(id string) string {
	return getGenericItem(id, config.SFX[id])
}

func (config *Config) getCustomValue(key string, id string) string {
	targetMap, ok := config.Custom[key]

	if !ok {
		return id
	}

	return getGenericItem(id, targetMap[id])
}

// GetArgumentValue will retrieve an argument from .tscrc.json for the given
// TSC definition, argument type index, and value string
func (config *Config) GetArgumentValue(definition TSCDefinition, idx int, value string) string {
	var argtype string

	if len(definition.ArgType) < idx {
		argtype = "number"
	} else {
		argtype = definition.ArgType[idx]
	}

	if definition.Key == "<GIT" {
		if value[0] == '0' {
			return config.getWeapon(value)
		}

		value = "0" + value[1:]
		return config.getItem(value)
	}

	if strings.HasPrefix(argtype, "custom:") {
		key := strings.Replace(argtype, "custom:", "", 1)

		if len(key) == 0 {
			return value
		}

		return config.getCustomValue(key, value)
	}

	switch argtype {
	case "number":
		return value
	case "direction":
		return config.getDirection(value)
	case "face":
		targetFace, _ := strconv.Atoi(value)
		return config.getFace(targetFace)
	case "map":
		return config.getMap(value)
	case "item":
		return config.getItem(value)
	case "equippable":
		return config.getEquippable(value)
	case "illustration":
		return config.getIllustration(value)
	case "song":
		return config.getSong(value)
	case "sfx":
		return config.getSFX(value)
	default:
		return value
	}
}

// Update will merge a given .tscrc.json string with the default tscrc
// configuration
func (config *Config) Update(newConfigJSON []byte) {
	newConfig := New()
	json.Unmarshal(newConfigJSON, &newConfig)
	*config = newConfig
}
