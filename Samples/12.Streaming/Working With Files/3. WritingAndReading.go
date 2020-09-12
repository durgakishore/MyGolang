package main

import (
	"fmt"
	"os"
)

func sampleWritingAndReading() {

	rows := []string{
		"Yerubandi Durga Kishore",
		"He is trying to learn Go language"}

	fout, err := os.Create("D:\\SampleIo.txt")

	if err != nil {
		fmt.Println("failed to create a new file")
		os.Exit(1)
	}
	defer fout.Close()

	for _, row := range rows {
		fout.WriteString(row)
	}
}
