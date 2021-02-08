package quickfix

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"os"

	"golang.org/x/lint"
)

func getFileLines(path string) ([]string, error) {
	if path == "" {
		return nil, errors.New("err... getting file lines")
	}
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, errors.New("err... getting file lines")
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("err... getting file lines")
	}
	return lines, nil
}

func getWarnings(path string) ([]lint.Problem, error) {
	if path == "" {
		return nil, errors.New("err... getting lint warnings")
	}
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		return nil, errors.New("err... getting lint warnings")
	}
	defer file.Close()
	fbytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("err... getting lint warnings")
	}
	l := lint.Linter{}
	problems, err := l.Lint(path, fbytes)
	if err != nil {
		return nil, errors.New("err... getting lint warnings")
	}
	var warnings []lint.Problem
	for _, p := range problems {
		if p.Confidence == 0.2 {
			continue
		}
		if p.Category != "comments" {
			continue
		}
		text := strings.Split(p.Text, " ")
		if text[0] != "exported" {
			continue
		}

		line := strings.Split(p.LineText, " ")
		if line[0] != "var" &&
			line[0] != "const" &&
			line[0] != "type" &&
			line[0] != "func" {
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
			return "", errors.New("err... generating comment")
		}
		if match {
			name = l
			nameSlice := strings.Split(name, "(")
			return fmt.Sprintf("// %+s ...", nameSlice[0]), nil
		}
	}
	return "", errors.New("err... generating comment")
}

func insertComment(file []string, comment string, offset int) []string {
	pre := make([]string, len(file))
	post := make([]string, len(file))
	copy(pre, file)
	copy(post, file)
	pre = pre[:offset]
	post = post[offset:]
	newFile := append(pre, comment)
	newFile = append(newFile, post...)
	return newFile
}

func generateNewFile(path string) ([]string, error) {
	fileSlice, err := getFileLines(path)
	if err != nil {
		return nil, errors.New("err... generating file")
	}
	warningSlice, err := getWarnings(path)
	if err != nil {
		return nil, errors.New("err... generating file")
	}
	for i, w := range warningSlice {
		offset := (w.Position.Line + i) - 1
		comment, err := generateComment(w)
		if err != nil {
			return nil, errors.New("err... generating file")
		}
		fileSlice = insertComment(fileSlice, comment, offset)
	}
	return fileSlice, nil
}

func fixFile(path string) error {
	newFile, err := generateNewFile(path)
	if err != nil {
		return errors.New("err... fixing file")
	}

	file, err := os.OpenFile(path, os.O_WRONLY, os.ModeDir)
	if err != nil {
		return errors.New("err... fixing file")
	}
	defer file.Close()

	for i := range newFile {
		_, err := file.WriteString(newFile[i] + "\n")
		if err != nil {
			log.Println(err)
			return errors.New("err... fixing file")
		}
	}
	return nil
}
