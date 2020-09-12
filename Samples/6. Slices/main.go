package main

import "fmt"

func main() {

	//multiDimentionalSlice()
	//sliceSplitting()

	pls := [][]string{
		{"C", "C++", "java"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	fmt.Println(countriess())
}

func countriess() []string {

	countries := []string{"USA", "INDIA", "AUSTRALIA", "CANADA", "GERMANY"}
	needCountries := countries[(len(countries)-2):]
	countriesCpy := make([]string, len(needCountries))
	copy(countriesCpy, needCountries)
	return countriesCpy
}
