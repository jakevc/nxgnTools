package main

/*
Filter out sequences in a fastq file whose kmer coverage
is greater than the specified kcov with k=ksize.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/montanaflynn/stats"
	"os"
)

var (
	kcov    = flag.Int("c", 11, "Specify the limit of the kmer coverage to filter at")
	ksize   = flag.Int("k", 10, "Specify the kmer size to filter on")
	infile  = flag.String("i", "", "Specify a fastq file for kmer filtering")
	outfile = flag.String("o", "", "Specify an output filename")
)

type fastqEntry struct {
	Head    string
	SeqLine string
	Plus    string
	Qual    string
}

func keepSeq(seq string, k int, kcov int) bool {
	// dict to store kmer counts
	kdict := make(map[string]float64)

	num := len(seq) - k + 1
	// kmerize the seq
	for i := 1; i < num; i++ {
		kmer := seq[i : k+i]
		// frequency of each kmer
		kdict[kmer] += 1
	}

	// get coverage
	kmerCoverage := make([]float64, 0, len(kdict))
	for k := range kdict {
		kmerCoverage = append(kmerCoverage, kdict[k])
	}

	// convert to stats type
	var covList stats.Float64Data = kmerCoverage

	// return true if median coverage is below coverage lmit
	med, err := covList.Median()
	if err != nil {
		fmt.Println("Error in Median Func:", err)
		return false
	}
	if med < float64(kcov) {
		return true
	}
	return false
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	// parse and check flags
	flag.Parse()
	if *infile == "" {
		fmt.Println("No input file!")
		flag.Usage()
		os.Exit(1)
	} else if *outfile == "" {
		fmt.Println("No output file!")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(flag.Args())

	infi, err := os.Open(*infile)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer infi.Close()

	// read lines into buffer
	scanner := bufio.NewScanner(infi)
	scanner.Split(bufio.ScanLines)

	// open file for writing
	of, err := os.Create(*outfile)
	if err != nil {
		panic(err)
	}
	defer of.Close()

	// write buffer
	w := bufio.NewWriter(of)

	// line num
	NR := 1

	// entry
	entry := fastqEntry{}

	// counter for number of seqs filtered
	numFiltered := 0

	for scanner.Scan() {
		switch lintype := NR % 4; lintype {
		case 0:
			entry.Qual = scanner.Text()
			fourLines := fmt.Sprintf("%s\n%s\n%s\n%s\n",
				entry.Head,
				entry.SeqLine,
				entry.Plus,
				entry.Qual)

			// write to file if keepSeq returns True
			if keepSeq(entry.SeqLine, *ksize, *kcov) {
				_, err := w.WriteString(fourLines)
				if err != nil {
					fmt.Println("Error writing entry: ", err)
				}
			} else {
				numFiltered += 1
			}

		// wrtie entry to file if keepseq true then flush entry

		case 1:
			entry.Head = scanner.Text()
		case 2:
			entry.SeqLine = scanner.Text()
		case 3:
			entry.Plus = scanner.Text()
		}
		NR += 1
	}

	w.Flush()
	fmt.Printf("Done. %d sequences were filtered out.\n", numFiltered)
}
