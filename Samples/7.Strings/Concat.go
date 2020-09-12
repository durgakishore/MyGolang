package main

import (
	"errors"
	. "fmt"
	"strings"
)

func concat(parts ...string) (string, error) {

	if len(parts) == 0 {
		return "", errors.New("No strings supplied")
	}
	return strings.Join(parts, " "), nil
}

func simpleStringConcat() {
	if result, err := concat("Yerubandi", "Durga", "Kishore", "From", "Rajahmundry"); err != nil {
		println("Error ", err)
	} else {
		Println("New string :", result)
	}
}
