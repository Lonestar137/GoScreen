package main

import (
	"GoScreen/pkg/capture"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
  An application which takes a screenshot on keypress of a specified monitor(s)
*/

//TODO Parse/Qualify cmdline arguments.
//TODO Option to pass which monitor to capture.

func main() {

	if len(os.Args) == 1 {
		layout := capture.Keylayout{Capture: "j", Quit: "q"}
		capture.Key_press_listener(layout)
	} else if len(os.Args) == 3 {
		// If program called with 2 arguments, then
		layout := capture.Keylayout{Capture: os.Args[1], Quit: os.Args[2]}
		capture.Key_press_listener(layout)

		//TODO add an elif here for help text
	} else if strings.Contains(os.Args[1], "help") {
		fmt.Println("Usage: GoScreen [OPTION] [OPTION]")
		fmt.Println("Default keybindings: \n	j	Capture screenshot.\n	q	quit/close program.\n")
		fmt.Println("Examples: \n	GoScreen a k	Changes the capture keybind to `a` and the close keybind to `k`.")

	} else {
		log.Fatal("Invalid number of custom key arguments passed.  ")
		fmt.Println("Either call the function with no arguments, or pass two arguments for capture and quit keybinds respectively.")
		fmt.Println("Example: GoScreen j k")
	}

}
