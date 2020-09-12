package main

import (
	"fmt"
	"io"
	"os"
)

func openFileSample() {

	rfile, err := os.OpenFile("D:\\Enhance.txt", os.O_RDONLY, 0666)

	if err != nil {
		fmt.Println("unable to open file")
		os.Exit(1)
	}

	wfile, err := os.OpenFile("D:\\ScanOS.txt", os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("unable to open file")
		os.Exit(1)
	}

	n, err := io.Copy(wfile, rfile)
	if err != nil {
		fmt.Println("Copy failed")
		os.Exit(1)
	}
	fmt.Printf("Copied %d of bytes from %s  to %s ", n, wfile.Name(), rfile.Name())
}
