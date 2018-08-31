package main

/*
This program converts all multiline fasta entries in a file to
single line fasta entries.
*/

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	// check if err and panic
	if e != nil {
		panic(e)
	}
}

func combineSeqLine(path string) {
	// read file by line
	inFile, err := os.Open(path)
	check(err)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var entry bytes.Buffer
	var first bool = true

	// scan and parse lines
	for scanner.Scan() {
		line := scanner.Text()

		// firstline case
		if first {
			entry.WriteString(line + "\n")
			first = false
		}
		// write header
		if strings.HasPrefix(line, ">") {
			header := line
			entry.WriteString("\n" + header + "\n")
		}
		// write seq
		if !strings.HasPrefix(line, ">") {
			entry.WriteString(line)
		}
	}
	fmt.Println(entry.String())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide path to multiline fasta file as command line argument")
		return
	}
	// read in the file
	combineSeqLine(os.Args[1])
}
