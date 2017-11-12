package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

var shrts []*ShirtField
var file *os.File

func readLine(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + ": ")
	text, _ := reader.ReadString('\n')
	return text[0 : len(text)-1]
}

func main() {

	var err error
	file, err = os.OpenFile("db.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cmd := readLine("Enter a command")

	doFunc := func() { log.Fatal("Command not found") }

	switch cmd {
	case "list", "ls":
		doFunc = listShirt
	case "remove", "rm", "del", "delete":
		doFunc = deleteShirt
	case "add", "a":
		doFunc = addShirt
	}

	for true {
		readFile()
		doFunc()

	}
}

func readFile() {

	shrts = []*ShirtField{}
	if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	if err := gocsv.UnmarshalFile(file, &shrts); err != nil {
		panic(err)
	}

}

func writeFile() {
	if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	file.Truncate(1)
	err := gocsv.MarshalFile(&shrts, file) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}
	file.Sync()

}
