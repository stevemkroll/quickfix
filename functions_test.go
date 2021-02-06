package main

import "testing"

var path string

func init() {
	path = "./person.go"
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
