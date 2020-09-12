package main

import (
	"io"
	"log"
	"os"
)

func createFile() {

	file, err := os.Create("D:\\Files\\Log\\123.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(file)
	file.Close()
}

func createDirOrFolder() {

	_, err := os.Stat("D:\\Files\\Log\\Kishore")

	if os.IsNotExist(err) {

		errDir := os.MkdirAll("D:\\Files\\Log\\Kishore", 0755)
		if errDir != nil {
			log.Fatal(errDir)
		}
	}
}

func fileRename() {

	oldName := "D:\\Files\\Log\\123.txt"
	newName := "D:\\Files\\Log\\File.txt"

	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

func copyFile() {

	newFile, err := os.Create("D:\\Files\\123.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer newFile.Close()

	
	sourceFile,err := os.OpenFile("D:\\Files\\Log\\123.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)

	if err != nil{
		log.Fatal(err)
	}
	defer sourceFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("copied %d bytes", bytesCopied)

}
