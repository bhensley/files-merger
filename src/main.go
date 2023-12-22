package main

import (
	"io"
	"os"
	"strconv"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func makeDir(dir string) {
	err := os.Mkdir(dir, 0755)
	checkErr(err)
}

func copyDir(dir string, outputDir string, dirLabel string) error {
	files, err := os.ReadDir(dir)
	checkErr(err)

	for i, file := range files {
		if !file.IsDir() {
			// Copy the file
			src, err := os.Open(dir + file.Name())
			checkErr(err)
			defer src.Close()

			dst, err := os.Create(outputDir + "/" + dirLabel + "_" + strconv.Itoa(i) + file.Name())
			checkErr(err)
			defer dst.Close()

			_, err = io.Copy(dst, src)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	// Expecting 3 arguments to start: directory_1 directory_2 output_directory
	args := os.Args[1:]

	if len(args) != 3 {
		panic("Invalid number of arguments. Expected 3.")
	}

	// Make sure our input directories exist
	if _, err := os.Stat(args[0]); os.IsNotExist(err) {
		panic("Directory 1 does not exist.")
	}

	if _, err := os.Stat(args[1]); os.IsNotExist(err) {
		panic("Directory 2 does not exist.")
	}

	// Does our output directory exist? If not, create it.
	if _, err := os.Stat(args[2]); os.IsNotExist(err) {
		makeDir(args[2])
	} else {
		// It exists, so let's delete it and recreate it.
		err := os.RemoveAll(args[2])
		checkErr(err)

		makeDir(args[2])
	}

	copyDir(args[0], args[2], "a")
	copyDir(args[1], args[2], "b")
}
