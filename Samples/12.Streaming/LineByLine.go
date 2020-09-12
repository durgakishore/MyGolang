package main

import (
	"bufio"
	"fmt"
	"os"
)

func lineByLineReading() {

	file, _ := os.Open("./CustomIOReader.go")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
