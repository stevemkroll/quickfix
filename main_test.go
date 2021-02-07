package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
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
		t.Fatal(err)
	}
	for i := range files {
		err := fixFile(files[i])
		if err != nil {
			t.Fatal(err)
		}
	}
}
