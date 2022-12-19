package main

import "fmt"

func Church() {
	fmt.Printf("\nYou see 2 people at the alter getting married.\n\n1. Talk to them\n2. Leave\n\n")
	fmt.Scanln(&choice)
	if choice == 1 {
		fmt.Printf("\nGroom: \"The road to the east is very dangerous. I wouldnt recommend heading there without a weapon!\"\nWife: \"A woman lives in town that may help you find one.\"\n\n")
		Church()
	} else if choice == 2 {
		Scene(town)
		Town()
	} else {
		fmt.Printf("Invalid input")
		Church()
	}
}
