package firstclass

import "fmt"

//AnonymousFunctions func
func AnonymousFunctions() {

	a := func() {
		fmt.Println("first class function is invoked")
	}
	a()
	fmt.Println("anonymousFunctions function is invoked")

	func(n string) {
		fmt.Println("Welcome ", n)
	}("Kishore")
}
