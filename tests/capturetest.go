package unitest

import (
	"GoScreen/pkg/capture"
	"fmt"
	"testing"
)

//PadToTest
func PadIntTest(t *testing.T) {
	var paddedInt string = capture.PadInt(10, 4)

	if paddedInt != "0010" {
		t.Errorf("PadIntTest: Padding incorrect, got: %s, want: %s", paddedInt, "0010")
	}

}
