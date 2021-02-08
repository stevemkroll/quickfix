package quickfix

import (
	"testing"
)

var path string

func init() {
	path = "./pet/pet.go"
}

func TestGetFileLines(t *testing.T) {
	lines, err := getFileLines(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, l := range lines {
		t.Log(l)
	}
}

func TestGetWarnings(t *testing.T) {
	warnings, err := getWarnings(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, w := range warnings {
		t.Log(w)
	}
}

func TestGenerateComments(t *testing.T) {
	warnings, err := getWarnings(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, w := range warnings {
		comment, err := generateComment(w)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(comment)
	}
}

func TestGenerateNewFile(t *testing.T) {
	newFile, err := generateNewFile(path)
	if err != nil {
		t.Fatal(err)
	}
	for i := range newFile {
		t.Log(newFile[i])
	}
}

func TestFixFile(t *testing.T) {
	err := fixFile(path)
	if err != nil {
		t.Fatal(err)
	}
}
