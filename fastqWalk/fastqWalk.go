package main

/*
Recursively search a directory for fastq files,
percent of seqs in each file greater than 30 nt long.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// define flag
var dir = flag.String("dir", "../testData", "specify top of directory to search")

func findFastq(rootpath string) []string {

	// list of fastq files, capacity 10
	fqList := make([]string, 0, 10)

	// walk command line dir for fastq of fq files
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".fastq" || filepath.Ext(path) == ".fq" {
			fqList = append(fqList, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Walk error!\n", err)
	}
	return fqList
}

func getLongSeqs(fqPath string) (string, float32) {
	// open fastq file
	file, err := os.Open(fqPath)
	if err != nil {
		fmt.Println("Error opening file:\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// line count
	NR := 1
	// seqs longer than 30
	longSeqs := 0

	for scanner.Scan() {
		// if a seqline
		if NR%4 == 2 {
			line := scanner.Text()
			if len(line) > 30 {
				longSeqs += 1
			}
		}
		NR += 1
	}
	// subtract NR starting value
	totalSeqs := (NR - 1) / 4
	percentLong := (float32(longSeqs) / float32(totalSeqs)) * 100
	return filepath.Base(fqPath), percentLong
}

func main() {
	// parse command line args
	flag.Parse()

	// absolute path of command line arg dir
	path, err := filepath.Abs(*dir)
	if err != nil {
		fmt.Println("Error getting filepath.Abs!\n", err)
	}
	fmt.Println("Executing walk in: ", path)

	// return fastq files
	fastqList := findFastq(path)

	// print the file and it's percent long seqs
	for _, file := range fastqList {
		fmt.Println(getLongSeqs(file))
	}
}
