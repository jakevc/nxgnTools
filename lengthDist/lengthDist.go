package main

/*
Compute the read length distribution of a 2line fasta file
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func lengthDist(path string) (dist map[int]int) {
	fi, err := os.Open(path)
	// catch the error if there is one
	if err != nil {
		panic(err)
	}
	// might as well close now
	defer fi.Close()
	scanner := bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)

	// initailize empty map to write to
	dist = make(map[int]int)

	for scanner.Scan() {
		// firstline
		line := scanner.Text()
		if !strings.HasPrefix(line, ">") {
			dist[len(line)] += 1
		}
	}

	fmt.Println("lenth\tcount")
	for k, v := range dist {
		fmt.Printf("%d\t%d\n", k, v)
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide path to twoline fasta file")
		return
	}
	// read the second command line arg
	lengthDist(os.Args[1])
}
