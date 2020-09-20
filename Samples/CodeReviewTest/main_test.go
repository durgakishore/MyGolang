package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// execute command:  go test -v

//TestCheckAndPrintWordsFromText is testing the JSON file reading
func TestCheckAndPrintWordsFromText(t *testing.T) {

	var tests = []struct {
		fileName string // .json file path
		word     string //search word
		result   bool   //expected result
	}{
		{"TestFiles/valid.json", "hunt", true},
		{"TestFiles/textempty.json", "data", false},
		{"TestFiles/valid.json", "", false},
		{"TestFiles/valid.json", "design,", true},
		{"TestFiles/valid.json", "channe", true}, //exact 4 words before and after "channe"
	}
	for _, test := range tests {

		jsonFile, err := os.Open(test.fileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer jsonFile.Close()

		b, _ := ioutil.ReadAll(jsonFile)

		var result []map[string]interface{}
		json.Unmarshal([]byte(b), &result)

		fmt.Printf("Checking file %s\n", test.fileName)
		if checkAndPrintWordsFromText(result, test.fileName, test.word) != test.result {
			fmt.Printf("Validation failed : %s\n", test.fileName)
			t.Fail()
		} else {
			fmt.Printf("Validation succeeded : %s\n", test.fileName)
		}
	}

}
