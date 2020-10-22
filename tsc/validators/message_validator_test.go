package validators_test

import (
	"testing"

	"github.com/nimblebun/tsc-language-server/config"
	"github.com/nimblebun/tsc-language-server/tsc/validators"
)

func TestValidateMessages(t *testing.T) {
	const ok = `<MSGSuuuuue!<NOD
Answer me!<NOD
I'm so hungry...<NOD
There's nothing to eat and I've
been reduced to feeding on
cockroaches.<NOD<CLO`

	const invalidPlain = `<MSGSuuuuue!<NOD
Answer me! I'm so hungry... There's nothing to eat and I've been
reduced to feeding on cockroaches<NOD<CLO`

	const invalidPortrait = `<MSG<FAC0010Suuuuue!<NOD
Answer me! I'm so hungry... There's<NOD
<FAC0000to eat and I've been reduced to
feeding on cockroaches.<NOD<CLO`

	conf := config.New()

	t.Run("should return empty diagnostics when there are no text overflows", func(t *testing.T) {
		document := dummyTextDocument(ok)
		diagnostics := validators.ValidateMessages(ok, document, &conf)

		if len(diagnostics) != 0 {
			t.Errorf(
				"tsc.validators.ValidateMessages(ok, document, &conf) got %v (%v), want %v",
				len(diagnostics),
				diagnostics,
				0,
			)
		}
	})

	t.Run("should return warnings when a plain message overflows", func(t *testing.T) {
		document := dummyTextDocument(invalidPlain)
		diagnostics := validators.ValidateMessages(invalidPlain, document, &conf)

		if len(diagnostics) != 1 {
			t.Errorf(
				"tsc.validators.ValidateMessages(invalidPlain, document, &conf) -> len got %v (%v), want %v",
				len(diagnostics),
				diagnostics,
				0,
			)
		}

		expectedMessage := "Message exceeds 35 characters (current length: 64). This may cause text overflow issues."
		if diagnostics[0].Message != expectedMessage {
			t.Errorf(
				"tsc.validators.ValidateMessages(invalidPlain, document, &conf)[0].Message got %v, want %v",
				diagnostics[0].Message,
				expectedMessage,
			)
		}
	})

	t.Run("should return warnings when a message with portrait overflows", func(t *testing.T) {
		document := dummyTextDocument(invalidPortrait)
		diagnostics := validators.ValidateMessages(invalidPortrait, document, &conf)

		if len(diagnostics) != 1 {
			t.Errorf(
				"tsc.validators.ValidateMessages(invalidPortrait, document, &conf) -> len got %v (%v), want %v",
				len(diagnostics),
				diagnostics,
				0,
			)
		}

		expectedMessage := "Message exceeds 28 characters (current length: 35). This may cause text overflow issues."
		if diagnostics[0].Message != expectedMessage {
			t.Errorf(
				"tsc.validators.ValidateMessages(invalidPortrait, document, &conf)[0].Message got %v, want %v",
				diagnostics[0].Message,
				expectedMessage,
			)
		}
	})
}
