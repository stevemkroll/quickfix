package quickfix

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
		log.Printf("\n\n%+v\n\n", err.Error())
		panic(err)
	}
	for i := range files {
		warnings, err := getWarnings(files[i])
		if err != nil {
			log.Printf("\n\n%+v\n\n", err.Error())
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
		log.Printf("\n\n%+v\n\n", err.Error())
		panic(err)
	}
	for i := range files {
		err := fixFile(files[i])
		if err != nil {
			log.Printf("\n\n%+v\n\n", err.Error())
			panic(err)
		}
	}
}
