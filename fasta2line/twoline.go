package main

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
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	var entry bytes.Buffer
	var first bool = true

	for scanner.Scan() {
		line := scanner.Text()

		// firstline case
		if first {
			entry.WriteString(line + "\n")
			first = false
		}
		if strings.HasPrefix(line, ">") {
			header := line
			entry.WriteString("\n" + header + "\n")
		}
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
