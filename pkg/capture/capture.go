package capture

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
	"strconv"
)

/*
  Takes an integer and pads 0's to the front of the int, returns a string
  @param num to pad
  @padTo how many zeros should be added.
*/
func PadInt(num int, padTo int) string {

	var numStr string = strconv.Itoa(num)

	for i := 0; i < padTo; i++ {
		if len([]rune(numStr)) < padTo {
			numStr += "0" + numStr
		}
	}

	return ""
}

/*
  Takes a screenshot of your entire display.
  @param startingNum, the index number you want assign to the screenshot.  This number will be incremented if you are capturing more than one monitor.
  @param monitor, the display you want to capture. -1 captures all monitors.  0 will capture your first monitor and so on. . .
*/
func TakeScreenshot(screenshotNumber int, monitor int) {

	n := screenshot.NumActiveDisplays()

	if monitor == -1 {
		for i := 0; i < n; i++ {
			bounds := screenshot.GetDisplayBounds(i)
			img, err := screenshot.CaptureRect(bounds)
			if err != nil {
				panic(err)
			}

			fileName := fmt.Sprintf("%d_%s.png", screenshotNumber+i, "GoScreen")
			file, _ := os.Create(fileName)
			defer file.Close()
			png.Encode(file, img)

		}
	} else {
		bounds := screenshot.GetDisplayBounds(monitor)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		fileName := fmt.Sprintf("%d_%s.png", screenshotNumber, "GoScreen")
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", screenshotNumber, bounds, fileName)
	}

}
