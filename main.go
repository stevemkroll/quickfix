package main

import (
	"flag"
	"log"
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
	flag.StringVar(&fixFlag, "fix", "none", "defines the quick fix you would like to perform")
	flag.Parse()
	log.Printf("fix flag is...%+v", fixFlag)
}
