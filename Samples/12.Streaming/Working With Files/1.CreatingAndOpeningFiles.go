package main

import (
	. "fmt"
	"io"
	"os"
)

func creatingAndOpeningFiles() {

	file1, err := os.Open("D:\\Enhance.txt")
	defer file1.Close()
	if err != nil {
		Println("Unbale to open the source file")
		os.Exit(1)
	}

	file2, err := os.Create("D:\\ScanOS.txt")

	defer file2.Close()
	if err != nil {
		Println("Unbale to Create the destination file")
		os.Exit(1)
	}

	n, err := io.Copy(file2, file1)
	if err != nil {
		Println("Copy failed")
		os.Exit(1)
	}
	Printf("Copied %d of bytes from %s  to %s ", n, file1.Name(), file2.Name())
}
