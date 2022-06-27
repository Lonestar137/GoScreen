package capture

import (
	"GoScreen/pkg/collect"
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"log"
	"os"
	"strconv"
)

// Keylayout -- used for creating keybinds in the application, used inside the Key_press_listener...
type Keylayout struct {
	Capture string
	Quit    string
}

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

//High level function which takes a screenshot with default parameters such as screenshot number etc. . .
func Snapshot() {
	var curr_screenshot int
	var err error
	curr_screenshot, err = collect.LastScreenshotNumber("./")
	if err != nil {
		log.Fatal(err)
	}
	var screenshotName int = curr_screenshot + 1
	fmt.Println("")
	TakeScreenshot(screenshotName, 0)
}

/*
  A function which when called pauses the terminal to await a key press.
  @param key  -- Custom type Keylayout. This param contains keybinds for the application such as the quit and capture keybind which are just strings in the type.
*/
func Key_press_listener(key Keylayout) {
	ch := make(chan string)
	go func(ch chan string) {
		var b = make([]byte, 1)
		for {
			os.Stdin.Read(b)
			ch <- string(b)
		}
	}(ch)

	for {
		stdin, _ := <-ch
		if stdin == key.Capture {
			Snapshot()
		} else if stdin == key.Quit {
			//cleanup  bfor exit would go here
			break
		}

	}
}
