package main

import "fmt"

func addShirt() {

	sf := ShirtField{}
	sf.Description = readLine("Description")
	sf.ID = readLine("Item ID")

	fmt.Println(sf)

	// Replace shirt if exists

	foundBefore := false
	for i := range shrts {
		if shrts[i].ID == sf.ID {
			shrts[i] = &sf
			foundBefore = true
			fmt.Println("found shirt")
			break
		}
	}
	if !foundBefore {
		shrts = append(shrts, &sf)
	}
	writeFile()
}

func listShirt() {

	shrtID := readLine("List Item ID")

	// Replace shirt if exists

	for i := range shrts {
		if shrts[i].ID == shrtID {
			fmt.Printf("found shirt: %s\n", shrts[i])
		}
	}

}

func deleteShirt() {
	shrtID := readLine("Delete Item ID")

	deleteID := -1

	for i := range shrts {
		if shrts[i].ID == shrtID {
			deleteID = i
			break
		}
	}

	if deleteID >= 0 {
		fmt.Printf("Deleting shirt %s \n", shrts[deleteID])
		shrts = append(shrts[:deleteID], shrts[deleteID+1:]...)
	}
	writeFile()
}
