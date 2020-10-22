package validators

import (
	"testing"
)

func TestValidateEvents(t *testing.T) {
	const ok = `#0090
<MNA<FL-0263<CMU0002<FAI0000<END
#0091
<MNA<FL-0263<CMU0002<FAI0001<END
#0092
<MNA<FL-0263<CMU0002<FAI0002<END
#0093
<MNA<FL-0263<CMU0002<FAI0003<END
#0094
<MNA<FL-0263<FLJ0341:0099<CMU0002<FAI0004<END`

	const bad = `#0090
<MNA<FL-0263<CMU0002<FAI0000<END
#0091
<MNA<FL-0263<CMU0002<FAI0001<END
#0092
<MNA<FL-0263<CMU0002<FAI0002<END
#0093
<MNA<FL-0263<CMU0002<FAI0003<END
#0092
<MNA<FL-0263<CMU0002<FAI0002<END
#0094
<MNA<FL-0263<FLJ0341:0099<CMU0002<FAI0004<END`

	t.Run("should return empty diagnostics slice when there are no duplicate events", func(t *testing.T) {
		document := dummyTextDocument(ok)
		diagnostics := validateEvents(ok, document)

		if len(diagnostics) != 0 {
			t.Errorf(
				"tsc.validators.validateEvents(ok, document) got %v (%v), want %v",
				len(diagnostics),
				diagnostics,
				0,
			)
		}
	})

	t.Run("should return diagnostics when there are duplicate events", func(t *testing.T) {
		document := dummyTextDocument(bad)
		diagnostics := validateEvents(bad, document)

		if len(diagnostics) != 1 {
			t.Errorf(
				"tsc.validators.validateEvents(bad, document) got %v (%v), want %v",
				len(diagnostics),
				diagnostics,
				1,
			)
		}
	})
}
