package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {

	ch <- "Yerubandi"
	ch <- "Durga"
	ch <- "Kishore"
	ch <- "Rishika"
	ch <- "Mrudula"
}

func getData(ch chan string) {

	for {
		fmt.Printf("%s \n", <-ch)
	}
}
