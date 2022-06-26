package main

import (
	"GoScreen/pkg/capture"
	"GoScreen/pkg/collect"
	"fmt"
	"log"
)

/*
  An application which takes a screenshot on keypress of a specified monitor(s)
*/

func snapshot() {
	var curr_screenshot int
	var err error
	curr_screenshot, err = collect.LastScreenshotNumber("./")
	if err != nil {
		log.Fatal(err)
	}
	curr_screenshot += 1
	fmt.Println("MAIN: Curr screenshot ", curr_screenshot)
	capture.TakeScreenshot(curr_screenshot, 0)

	fmt.Println(curr_screenshot)

}

func main() {
	//TODO keyboard event loop
	// onKeyPress{
	// 	snapshot()
	// }
	//

	snapshot()
}
