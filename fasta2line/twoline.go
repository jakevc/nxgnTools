package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLine(path string) {
	// read file by line
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide input file path")
		return
	}
	// read in the file
	readLine(os.Args[1])
}
