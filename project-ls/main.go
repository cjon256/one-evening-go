package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	listAll := flag.Bool("a", false, "List all files")
	flag.Parse()
	testdata := "testdata"
	files := listFiles(testdata)
	for _, f := range files {
		if !*listAll && strings.HasPrefix(f, ".") {
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
