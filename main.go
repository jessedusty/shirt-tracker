package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
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

	openFile()
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
	case "location", "loc":
		location := readLine("New Location")
		doFunc = func() { setLocation(location) }
	case "aloc", "alocation":
		readFile()
		setAllLocations()
		os.Exit(0)
	case "retag":
		readFile()
		format := readLine("Shrt ID format ex shirt%d")
		start, err := strconv.Atoi(readLine("Starting Index"))
		if err != nil {
			log.Fatal("Starting index invalid")
		}
		stop, err := strconv.Atoi(readLine("Stop at Index"))
		if err != nil {
			log.Fatal("Stop index invalid")
		}
		retagShirts(start, stop, format)
		os.Exit(0)
	}

	for true {
		readFile()
		doFunc()
	}
}

func openFile() {
	var err error
	file, err = os.OpenFile("db.json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

}

func readFile() {

	shrts = []*ShirtField{}
	if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(fileContents, &shrts); err != nil {
		log.Fatal(err)
	}

}

func writeFile() {

	fileContents, err := json.MarshalIndent(&shrts, "", " ") // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}

	if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}

	file.Truncate(1)

	file.Write(fileContents)
	file.Sync()

}
