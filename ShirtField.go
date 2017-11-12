package main

import "fmt"

// ShirtField field that holds information for a shirt
type ShirtField struct {
	ID, Description string
}

func (sf *ShirtField) String() string {
	return fmt.Sprintf("%s: %s", sf.ID, sf.Description)
}
