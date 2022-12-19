package main

import "fmt"

func Town() {
	fmt.Printf("You see a church in the distance with roads heading North and East.\nYour home neighbours another small house.\nWhere do you want to go?\n\n")
	fmt.Printf("1. North road\n2. East road\n3. Church\n4. Home\n5. Neighboring house\n\n")
	fmt.Scanln(&choice)
	if choice == 4 {
		Home()
	} else {
		fmt.Println("Feature coming soon")
		Town()
	}
}
