package main

import (
	numeral_parser "github.com/0B1t322/russian-words-numbers-to-numbers/internal/numeral-parser"
	"os"
)

const (
	fileNameToRead  = "test_text.txt"
	fileNameToWrite = "out_text.txt"
)

func main() {
	fileRead, err := os.Open("./" + fileNameToRead)
	if err != nil {
		panic(err)
	}
	defer fileRead.Close()

	fileWrite, err := os.Create("./" + fileNameToWrite)
	if err != nil {
		panic(err)
	}
	defer fileWrite.Close()

	parser := numeral_parser.NewParser(fileRead, fileWrite)
	defer parser.Close()

	err = parser.ParseAllLines()
	if err != nil {
		panic(err)
	}
}
