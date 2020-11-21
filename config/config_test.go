package config

import (
	"testing"
)

// this test can be improved
func TestNew(t *testing.T) {
	t.Run("should return the default configuration", func(t *testing.T) {
		config := New()

		// default config has 91 TSC command definitions
		if len(config.TSC) != 91 {
			t.Errorf("config.New(): got %v, want %v", len(config.TSC), 91)
		}
	})
}

func TestGetTSCDefinition(t *testing.T) {
	conf := New()

	t.Run("should return false when the command is not found", func(t *testing.T) {
		_, found := conf.GetTSCDefinition("<SUE")

		if found {
			t.Errorf("config.Config#GetTSCDefinition(\"<SUE\") got %v, want %v", found, false)
		}
	})

	t.Run("should return correct TSC definition", func(t *testing.T) {
		targetDefinition := TSCDefinition{
			Label:         "<AM-",
			Detail:        "ArMs -",
			Documentation: "Remove weapon W.",
			Format:        "<AM-WWWW",
			InsertText:    "AM-${1:0000}",
			ArgType:       []string{"weapon"},
		}

		definition, found := conf.GetTSCDefinition("<AM-")
		ok := true

		if !found {
			t.Errorf("config.Config#GetTSCDefinition(\"<AM-\") second return value, got %v, want %v", found, true)
		}

		ok = definition.Label == targetDefinition.Label
		ok = definition.Detail == targetDefinition.Detail
		ok = definition.Documentation == targetDefinition.Documentation
		ok = definition.Format == targetDefinition.Format
		ok = definition.GetInsertText() == targetDefinition.InsertText
		ok = definition.Nargs() == targetDefinition.Nargs()
		ok = len(definition.ArgType) == len(targetDefinition.ArgType)
		ok = definition.ArgType[0] == targetDefinition.ArgType[0]

		if !ok {
			t.Errorf("config.Config#GetTSCDefinition(\"<AM-\") got %v, want %v", definition, targetDefinition)
		}
	})
}

func TestGetArgumentValue(t *testing.T) {
	conf := New()
	tra, _ := conf.GetTSCDefinition("<TRA")
	fai, _ := conf.GetTSCDefinition("<FAI")
	fac, _ := conf.GetTSCDefinition("<FAC")
	itplus, _ := conf.GetTSCDefinition("<IT+")
	eqplus, _ := conf.GetTSCDefinition("<EQ+")
	sil, _ := conf.GetTSCDefinition("<SIL")
	cmu, _ := conf.GetTSCDefinition("<CMU")
	sou, _ := conf.GetTSCDefinition("<SOU")
	git, _ := conf.GetTSCDefinition("<GIT")

	t.Run("should return number", func(t *testing.T) {
		val := conf.GetArgumentValue(tra, 2, "0002")
		want := "0002"

		if val != want {
			t.Errorf("config.Config#GetArgumentValue(tra, 2, \"0002\") got %v, want %v", val, want)
		}
	})

	t.Run("should return direction", func(t *testing.T) {
		val := conf.GetArgumentValue(fai, 0, "0001")
		want := "Up"

		if val != want {
			t.Errorf("config.Config#GetArgumentValue(fai, 0, \"0001\") got %v, want %v", val, want)
		}
	})

	t.Run("should return face", func(t *testing.T) {
		val := conf.GetArgumentValue(fac, 0, "0001")
		want := "Sue (smiling)"

		if val != want {
			t.Errorf("config.Config#GetArgumentValue(fac, 0, \"0001\") got %v, expect %v", val, want)
		}
	})

	t.Run("should return map", func(t *testing.T) {
		val := conf.GetArgumentValue(tra, 0, "0010")
		want := "Sand - Sand Zone"

		if val != want {
			t.Errorf("config.Config#GetArgumentValue(tra, 0, \"0010\") got %v, want %v", val, want)
		}
	})

	t.Run("should return item", func(t *testing.T) {
		val := conf.GetArgumentValue(itplus, 0, "0002")
		want := "Map System"

		if val != want {
			t.Errorf("config.Config#GetArgumentValue(itplus, 0, \"0002\") got %v, want %v", val, want)
		}
	})

	t.Run("should return equippable", func(t *testing.T) {
		val := conf.GetArgumentValue(eqplus, 0, "0004")
		want := "Arms Barrier"

		if val != want {
			t.Errorf("config.Config#GetArgumentValue(eqplus, 0, \"0004\") got %v, want %v", val, want)
		}
	})

	t.Run("should return illustration", func(t *testing.T) {
		val := conf.GetArgumentValue(sil, 0, "0017")
		want := "King, Jack, Sue"

		if val != want {
			t.Errorf("config.Config#GetArgumentValue(sil, 0, \"0017\") got %v, want %v", val, want)
		}
	})

	t.Run("should return song", func(t *testing.T) {
		val := conf.GetArgumentValue(cmu, 0, "0036")
		want := "Running Hell"

		if val != want {
			t.Errorf("config.Config#GetArgumentValue(cmu, 0, \"0036\") got %v, want %v", val, want)
		}
	})

	t.Run("should return sfx", func(t *testing.T) {
		val := conf.GetArgumentValue(sou, 0, "0003")
		want := "bump head"

		if val != want {
			t.Errorf("config.Config#GetArgumentValue(sou, 0, \"0003\") got %v, want %v", val, want)
		}
	})

	t.Run("on edge cases", func(t *testing.T) {
		t.Run("should default to number on OOB", func(t *testing.T) {
			val := conf.GetArgumentValue(tra, 20, "0003")
			want := "0003"

			if val != want {
				t.Errorf("config.Config#GetArgumentValue(tra, 20, \"0003\") got %v, want %v", val, want)
			}
		})

		t.Run("should get weapon on <GIT with value that starts with 0", func(t *testing.T) {
			val := conf.GetArgumentValue(git, 0, "0003")
			want := "Fireball"

			if val != want {
				t.Errorf("config.Config#GetArgumentValue(git, 0, \"0003\") got %v, want %v", val, want)
			}
		})

		t.Run("should get item on <GIT with value that doesn't start with 0", func(t *testing.T) {
			val := conf.GetArgumentValue(git, 0, "1003")
			want := "Santa's Key"

			if val != want {
				t.Errorf("config.Config#GetArgumentValue(git, 0, \"1003\") got %v, want %v", val, want)
			}
		})
	})

	t.Run("on custom commands", func(t *testing.T) {
		customCommand := TSCDefinition{
			Label:         "<HEY",
			Detail:        "HEY!",
			Documentation: "Custom Command",
			Format:        "<HEYXXXX:YYYY",
			InsertText:    "HEY${1:0000}:${2:0000}",
			ArgType:       []string{"custom:", "custom:test"},
		}

		customConf := New()
		customConf.Custom = CustomConfig{
			"test": map[string]string{
				"0001": "Chie",
			},
		}

		t.Run("should resolve custom argument type", func(t *testing.T) {
			val := customConf.GetArgumentValue(customCommand, 1, "0001")
			want := "Chie"

			if val != want {
				t.Errorf("config.Config#GetArgumentValue(customCommand, 1, \"0001\") got %v, want %v", val, want)
			}
		})

		t.Run("should return raw value on empty custom type", func(t *testing.T) {
			val := customConf.GetArgumentValue(customCommand, 0, "0002")
			want := "0002"

			if val != want {
				t.Errorf("config.Config#GetArgumentValue(customCommand, 0, \"0002\") got %v, want %v", val, want)
			}
		})
	})
}

func TestUpdate(t *testing.T) {
	json := `{
	"setup": {
		"maxMessageLineLength": {
			"portrait": 30
		}
	},

	"tsc": {
		"<MIM": {
      "label": "<MIM",
      "detail": "MImiga Mask",
      "documentation": "Give player Mimiga mask X.",
      "format": "<MIMXXXX",
      "insertText": "MIM${1:0000}",
      "argtype": [ "number" ]
    }
	},

	"faces": [
		"reset",
		"Sue",
		"Curly"
	]
}`

	t.Run("should keep untouched values", func(t *testing.T) {
		conf := New()
		conf.Update([]byte(json))

		if conf.Setup.MaxMessageLineLength.Plain != 35 {
			t.Errorf(
				"config.Config#Update() -> Setup.MaxMessageLineLength.Plain got %v, want %v",
				conf.Setup.MaxMessageLineLength.Plain,
				35,
			)
		}

		_, found := conf.GetTSCDefinition("<TRA")

		if !found {
			t.Errorf(
				"config.Config#Update() -> config.Config#GetTSCDefinition(\"<TRA\") `found` got %v, want %v",
				found,
				true,
			)
		}
	})

	t.Run("should overwrite existing stuff and add new stuff", func(t *testing.T) {
		conf := New()
		conf.Update([]byte(json))

		if conf.Setup.MaxMessageLineLength.Portrait != 30 {
			t.Errorf(
				"config.Config#Update() -> Setup.MaxMessageLineLength.Portrait got %v, want %v",
				conf.Setup.MaxMessageLineLength.Portrait,
				30,
			)
		}

		mim, found := conf.GetTSCDefinition("<MIM")

		if !found {
			t.Errorf(
				"config.Config#Update() -> config.Config#GetTSCDefinition(\"<MIM\") `found` got %v, want %v",
				found,
				true,
			)
		}

		if mim.Label != "<MIM" {
			t.Errorf(
				"config.Config#Update() -> config.Config#GetTSCDefinition(\"<MIM\") `mim` got %v, want %v",
				mim.Label,
				"<MIM",
			)
		}

		expectedFaces := []string{"reset", "Sue", "Curly"}
		for i := range expectedFaces {
			got := conf.Faces[i]
			want := expectedFaces[i]

			if got != want {
				t.Errorf(
					"config.Config#Update() -> config.Config.Faces@%d got %v, want %v",
					i,
					got,
					want,
				)
			}
		}
	})
}

func TestNargs(t *testing.T) {
	t.Run("should return 0 when argtypes is empty/undefined", func(t *testing.T) {
		def := TSCDefinition{}

		if def.Nargs() != 0 {
			t.Errorf("config.TSCDefinition#Nargs(): got %v, want %v", def.Nargs(), 0)
		}
	})

	t.Run("should return correct number of arguments based on argtypes", func(t *testing.T) {
		def := TSCDefinition{
			ArgType: []string{"number", "number"},
		}

		if def.Nargs() != 2 {
			t.Errorf("config.TSCDefinition#Nargs(): got %v, want %v", def.Nargs(), 2)
		}
	})
}

func TestGetInsertText(t *testing.T) {
	t.Run("should return inserttext when provided", func(t *testing.T) {
		expected := "AME${1:0000}"

		def := TSCDefinition{
			InsertText: expected,
		}

		if def.GetInsertText() != expected {
			t.Errorf("config.TSCDefinition#GetInsertText(): got %v, want %v", def.GetInsertText(), expected)
		}
	})

	t.Run("should return empty string when command label is not provided", func(t *testing.T) {
		def := TSCDefinition{}

		if len(def.GetInsertText()) != 0 {
			t.Errorf("config.TSCDefinition#GetInsertText(): got length of %v, want 0", len(def.GetInsertText()))
		}
	})

	t.Run("should return command when there are no args", func(t *testing.T) {
		def := TSCDefinition{
			Label: "<AME",
		}

		expected := "AME"

		if def.GetInsertText() != expected {
			t.Errorf("config.TSCDefinition#GetInsertText(): got %v, want %v", def, expected)
		}
	})

	t.Run("should generate insert text on the fly", func(t *testing.T) {
		def := TSCDefinition{
			Label:   "<AME",
			ArgType: []string{"number", "number"},
		}

		expected := "AME${1:0000}:${2:0000}"

		if def.GetInsertText() != expected {
			t.Errorf("config.TSCDefinition#GetInsertText(): got %v, want %v", def.GetInsertText(), expected)
		}
	})
}
