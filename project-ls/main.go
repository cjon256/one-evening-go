package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	testdata := "testdata"
	files := listFiles(testdata)
	for _, f := range files {
		if strings.HasPrefix(f, ".") {
			continue
		}
		fmt.Println(f)
	}
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		dirs = append(dirs, f.Name())
	}

	return dirs
}
