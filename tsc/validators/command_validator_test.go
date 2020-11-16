package validators_test

import (
	"testing"

	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/tsc/validators"
)

func TestValidateCommands(t *testing.T) {
	const ok = "<MNA<CMU0023<FAI0000<END"
	const tooMany = "<MNA2222"
	const tooFew = "<TRA0002"

	conf := config.New()

	t.Run("should return empty diagnostics slice when there are no argc issues", func(t *testing.T) {
		document := dummyTextDocument(ok)
		diagnostics := validators.ValidateCommands(ok, document, &conf)

		if len(diagnostics) != 0 {
			t.Errorf(
				"tsc.validators.ValidateCommands(ok, document, &conf) got %v (%v), want %v",
				len(diagnostics),
				diagnostics,
				0,
			)
		}
	})

	t.Run("should return errors when a command has too many arguments", func(t *testing.T) {
		document := dummyTextDocument(tooMany)
		diagnostics := validators.ValidateCommands(tooMany, document, &conf)

		if len(diagnostics) != 1 {
			t.Errorf(
				"tsc.validators.ValidateCommands(tooMany, document, &conf) -> len got %v (%v), want %v",
				len(diagnostics),
				diagnostics,
				1,
			)
		}

		expectedMessage := "Too many arguments provided to <MNA. Expected 0, got 1."
		if diagnostics[0].Message != expectedMessage {
			t.Errorf(
				"tsc.validators.ValidateCommands(tooMany, document, &conf)[0].Message got %v, want %v",
				diagnostics[0].Message,
				expectedMessage,
			)
		}
	})

	t.Run("should return errors when a command has too few arguments", func(t *testing.T) {
		document := dummyTextDocument(tooFew)
		diagnostics := validators.ValidateCommands(tooFew, document, &conf)

		if len(diagnostics) != 1 {
			t.Errorf(
				"tsc.validators.ValidateCommands(tooFew, document, &conf) -> len got %v (%v), want %v",
				len(diagnostics),
				diagnostics,
				1,
			)
		}

		expectedMessage := "Too few arguments provided to <TRA. Expected 4, got 1."
		if diagnostics[0].Message != expectedMessage {
			t.Errorf(
				"tsc.validators.ValidateCommands(tooFew, document, &conf)[0].Message got %v, want %v",
				diagnostics[0].Message,
				expectedMessage,
			)
		}
	})
}
