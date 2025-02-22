package main

import (
	"fmt"
	"io"
	"os"
)

var (
	dirsToCheck = []string{"."}
	dirCount    uint
	fileCount   uint
)

func main() {

	if len(os.Args) > 1 {
		dirsToCheck = os.Args[1:]
	}

	for _, dir := range dirsToCheck {
		getCountsForDir(dir, 0)
	}

	fmt.Printf("%v directories, %v files\n", dirCount, fileCount)

}

func getCountsForDir(dirName string, level int) {

	f, err := os.Open(dirName)
	if err != nil {
		fmt.Printf("error opening directory: %v, error: %v\n", dirName, err)
		return
	}
	defer f.Close()

	dir, err := f.Readdir(-1)
	if err == io.EOF {
		return
	}

	lastIndex := len(dir) - 1
	s := "├── "
	for i, entry := range dir {
		if i == lastIndex {
			s = "└── "
		}
		fmt.Printf("%*v%v\n", level, s, entry.Name())
		if entry.IsDir() {
			dirCount++
			getCountsForDir(dirName+string(os.PathSeparator)+entry.Name(), level+4)
		}

		fileCount++
	}

}
