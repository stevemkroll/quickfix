package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"os"

	"golang.org/x/lint"
)

func getFileLines(path string) ([]string, error) {
	if path == "" {
		return nil, errors.New("err... no file path")
	}
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, errors.New("err... cannot find file")
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("err... parsing file")
	}
	return lines, nil
}

func getWarnings(path string) ([]lint.Problem, error) {
	if path == "" {
		return nil, errors.New("err... no file path")
	}
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, errors.New("err... cannot find file")
	}
	defer file.Close()
	fbytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("err... cannot read file")
	}
	l := lint.Linter{}
	problems, err := l.Lint(path, fbytes)
	if err != nil {
		return nil, errors.New("err... lint error")
	}
	var warnings []lint.Problem
	for _, p := range problems {
		if p.Confidence == 0.2 {
			continue
		}
		warnings = append(warnings, p)
	}
	return warnings, nil
}

func generateComment(problem lint.Problem) (string, error) {
	line := problem.LineText
	lineSlice := strings.Split(line, " ")

	name := ""
	for _, l := range lineSlice {
		match, err := regexp.MatchString(`^([A-Z]{1,}[a-zA-Z()]{1,})$`, l)
		if err != nil {
			return "", errors.New("err... regex error")
		}
		if match {
			name = l
			nameSlice := strings.Split(name, "(")
			return fmt.Sprintf("// %+s ...", nameSlice[0]), nil
		}
	}
	return "", errors.New("err... generating func name")
}
