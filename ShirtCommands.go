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

func useShirt() {
	shrtID := readLine("Use Item ID")

	for i := range shrts {
		if shrts[i].ID == shrtID {
			fmt.Printf("Using shirt %s \n", shrts[i])
			shrts[i].Location = "in use"
			return
		}
	}
	writeFile()
}

func setLocation(location string) {

	shrtID := readLine("List Item ID")

	// Replace shirt if exists

	for i := range shrts {
		if shrts[i].ID == shrtID {
			shrts[i].Location = location
			fmt.Printf("Set Location: %s\n", shrts[i])
		}
	}
	writeFile()
}

func setAllLocations() {
	oldLocation := readLine("Set all location from")
	newLocation := readLine("Set all location to")

	for i := range shrts {
		if shrts[i].Location == oldLocation {

			shrts[i].Location = newLocation
			fmt.Printf("Relocated shirt %s\n", shrts[i])
		}
		writeFile()
	}

}

func retagShirts(start, stop int, format string) {
	initial := start
	end := stop
	direction := 1

	if initial > end {
		direction = -1

	}

	for i := initial; true; i += direction {
		oldShrtID := fmt.Sprintf(format, i)

		var shrt *ShirtField
		for i := range shrts {
			if shrts[i].ID == oldShrtID {
				shrt = shrts[i]
			}
		}
		if shrt == nil {
			fmt.Printf("shirt %s not found\n", oldShrtID)
			return
		}

		fmt.Printf("Old Shirt - %s\n", shrt)
		newShrtID := readLine("New shirt ID")
		shrt.ID = newShrtID

		fmt.Printf("Changed shirt ID from %s to %s - %s\n", oldShrtID, newShrtID, shrt)
		writeFile()
	}
}
