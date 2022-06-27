package collect

import (
	"GoScreen/pkg/capture"
	"fmt"
	"log"
	"testing"
)

//Test TestLastScreenshotNumber
func TestLastScreenshotNumber(t *testing.T) {
	lastScreenshot, err := LastScreenshotNumber("../../")
	if err != nil {
		log.Fatal("COLLECTEST.TestLastScreenshotNumber failed: ", err)
	}
	capture.TakeScreenshot(lastScreenshot, 0)

}

// Example function -- it's a real test btw.
func ExampleLastScreenshotNumber() {
	lastScreenshot, err := LastScreenshotNumber("../../tests/")
	if err != nil {
		log.Fatal("COLLECTEST.ExampleLastScreenshotNumber failed: ", err)
	}

	if lastScreenshot == 0 {
		fmt.Println(lastScreenshot)
	}

	// Output:
	// 0
}
