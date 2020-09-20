/*
1. Find the matched word in "text" field data and then print 4 before and 4 after words with respect to matched word.
2. This program will read all json files in a directory and then print before and after 4 words in the "text" tag for the flag value.*/
package main

import (
	"archive/zip"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	word := flag.String("word", "", "a word")
	flag.Parse()

	r, err := zip.OpenReader("foc-slack-export.zip")
	defer r.Close()
	if err != nil {
		panic(err)
	}

	for _, f := range r.File {
		if !f.FileInfo().IsDir() {
			if filepath.Ext(f.Name) == ".json" {
				m, err := readDataFromJSONFile(f)
				if err != nil {
					fmt.Printf("error while reading : %v, file : %s", err, f.Name)
					continue
				}
				checkAndPrintWordsFromText(m, f.Name, *word)
			}
		}
	}
}

// This function reads the data from available JSON files in the Zipped folder directory.
func readDataFromJSONFile(file *zip.File) ([]map[string]interface{}, error) {

	var data []map[string]interface{}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from ", r)
		}
	}()

	rc, err := file.Open()
	defer rc.Close()
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(rc)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}
	return data, nil
}

/*This function split the "text" field data into words by delimiting with space and
if macthes found with the flag word then will display 4 words before and after the matched word.
For Testing purpose, return statement is added.*/
func checkAndPrintWordsFromText(data []map[string]interface{}, fileName string, word string) bool {
	bRet := false
	for _, m := range data {

		if text, ok := m["text"].(string); ok {
			words := strings.Split(text, " ")
			len := len(words)
			for i, w := range words {
				if w == word {
					if i >= 4 && i < len-4 {
						bRet = true
						r := strings.Trim(fmt.Sprint(words[i-4:i+5]), "[]")
						fmt.Printf("File name %s : %s \n\n", fileName, r)
					} else {
						fmt.Printf("Search word found in %s but required number of adjacent words are not available in that message\n", fileName)
					}
				}
			}
		}
	}
	return bRet
}

/* Review comments
1. Use filepath package to do file path,extension related operatons.
2. defer statement will execute in panic condtions also but we should register it before panic arise.
3. If it is known error and still have to proceed further then better to avoid panic at those places. Instead return err string.
4. In panic situations also, if we want to continue execution further then recover it.
5. Try to write the small code blocks to handle the functionality. It is very useful for understanding and issue debugging.
6. len(words)-- one time length calculation is enough instead of calculate and check in "if" condition everytime.
7. No need to add and fill 2 string arrays for before and after words. We can directly use "words" slice.
8. We can use string functions and format specifiers more effectively to print the output.
*/
