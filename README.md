# nxgnTools in Golang
Learning golang by rewriting nxgn tools in golang

# FASTA Tools 
	- fasta2line : Converts multiline fasta to twoline fasta file and prints to stdout.

	- lengthDist : Calculates the readlength distribution of a fasta file and prints to stdout.

	- fastqWalk : Recursively search a directory for fastq files, and print the file wiht the percent of seqs in the file greater than 30 nt long. 

# Kmer Tools

	- knorm : removes reads from a fastq file with kmer coverage greater than c.

	- kspec : writes kmer distribution of fastq file to stdout


# Alignment tools 

	- getGenename : given a file with the chromosome and position fields, and a gff annotation, return the original file with a gene field, if any

# TODO

	 - kspec
	 - getGenname
