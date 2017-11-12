package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {

	file, err := os.OpenFile("db.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	shrts = []*ShirtField{}

	if err := gocsv.UnmarshalFile(file, &shrts); err != nil {
		panic(err)
	}

	fileContents, err := json.Marshal(&shrts) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.OpenFile("db.json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	jsonFile.Write(fileContents)

}

var shrts []*ShirtField
var file *os.File

func readFile() {

}

func writeFile() {
}

// ShirtField field that holds information for a shirt
type ShirtField struct {
	ID, Description string
}

func (sf *ShirtField) String() string {
	return fmt.Sprintf("%s: %s", sf.ID, sf.Description)
}
