package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// flags
var fixFlag string

// strings
var (
	fixNone    string = "none"
	fixAll     string = "all"
	fixStructs string = "structs"
)

// init
func init() {
	// flag.StringVar(&fixFlag, "fix", "none", "defines the quick fix you would like to perform")
	// flag.Parse()
}

// main
func main() {
	// flag.StringVar(&fixFlag, "fix", "none", "defines the quick fix you would like to perform")
	// flag.Parse()
	// log.Printf("fix flag is...%+v", fixFlag)
}

// PreviewWarnings displays all the warnings that will be fixed
func PreviewWarnings() {
	var files []string
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		path = filepath.Clean(path)
		dir, file := filepath.Split(path)
		if !strings.HasPrefix(dir, ".") &&
			!strings.HasPrefix(file, ".") &&
			strings.Contains(file, ".go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for i := range files {
		warnings, err := getWarnings(files[i])
		if err != nil {
			panic(err)
		}
		if len(warnings) > 0 {
			for i := range warnings {
				log.Println(warnings[i].Text)
			}
		}
	}
}

// FixAll fixes all supported warnings
func FixAll() {
	var files []string
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		path = filepath.Clean(path)
		dir, file := filepath.Split(path)
		if !strings.HasPrefix(dir, ".") &&
			!strings.HasPrefix(file, ".") &&
			strings.Contains(file, ".go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	for i := range files {
		err := fixFile(files[i])
		if err != nil {
			panic(err)
		}
	}
}
