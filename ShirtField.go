package main

import "fmt"

// ShirtField field that holds information for a shirt
type ShirtField struct {
	ID, Description, Location string
}

func (sf *ShirtField) String() string {
	return fmt.Sprintf("%s \t %s \t %s", sf.ID, sf.Description, sf.Location)
}
