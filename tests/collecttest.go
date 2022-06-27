package collect

import (
	"fmt"
	"log"
	"testing"
)

func TestGetLastScreenshotNumber(t *testing.T) {
	lastScreenshot, err := collect.LastScreenshotNumber("")
	if err != nil {
		log.Fatal("COLLECTEST failed: ", err)
	}

	if lastScreenshot == 0 {
		fmt.Println("Success")
	}

}
