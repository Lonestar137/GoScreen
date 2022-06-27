package capture

import (
	"testing"
)

//PadToTest
func PadIntTest(t *testing.T) {
	var paddedInt string = PadInt(10, 4)

	if paddedInt != "0010" {
		t.Errorf("PadIntTest: Padding incorrect, got: %s, want: %s", paddedInt, "0010")
	}

}
