package collect

import (
	"io/fs"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

/*
  Given a dir returns a list of files in that dir
  @param dirname is the fully qualified path to the dir you want to scan.
  @param showdirs is optional
*/
func ReadDir(dirname string, showDirs bool) []fs.FileInfo {

	var xfiles []fs.FileInfo // List of filenames to be returned.
	if showDirs == true {
		//Print dirs
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			xfiles = append(xfiles, file)
		}
	} else {
		// Dont print dirs
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			if !file.IsDir() {
				xfiles = append(xfiles, file)
			}
		}
	}

	return xfiles
}

/*
  Given a directory(default is current), returns the number of the last .png file generated in the dir.
  @path the directory to look inside of for other screenshots.
*/
func LastScreenshotNumber(path string) (int, error) {
	dir := ""
	if path != "" {
		dir = path
	}

	var files []fs.FileInfo = ReadDir(dir, false)
	var screenshotNumber string = "0" // holds string value of the last reported screenshot number.
	for _, file := range files {
		if strings.Contains(file.Name(), "_") {
			var s []string = strings.Split(file.Name(), "_")

			//Type conversion to compare strings by value.
			currNum, err := strconv.Atoi(s[0])
			savedNum, yerr := strconv.Atoi(screenshotNumber)
			if err != nil {
				log.Fatal(err, "\n\nFailed to convert screenshot string names to int in collect.LastScreenshotNumber function. Failed number: ", s[0])
			} else if yerr != nil {
				// Basically if screenshot number isn't basenum and it fails for some reason report it.
				log.Fatal(err, "\n\nFailed to convert screenshotNumber string to int in collect.LastScreenshotNumber function. Failed number: ", screenshotNumber)
			}

			if currNum > savedNum {
				screenshotNumber = s[0]
			}

		}

	}

	//convert string to int and return it.
	return strconv.Atoi(screenshotNumber)
}
