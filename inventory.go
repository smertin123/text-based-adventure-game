package main

import "fmt"

var apples int

func Inventory(return_function func()) {
	if apples == 0 {
		fmt.Printf("\nInventory empty\n\n")
		return_function()
	} else if apples > 0 {
		fmt.Printf("\nYou have %d apples\n\n", apples)
		fmt.Printf("\n1. Use item\n2. Continue\n\n")
		fmt.Scanln(&choice)
		if choice == 1 {
			fmt.Printf("\nYou have eaten an apple\n\n")
			Bite_sound()
			fmt.Printf(apple_eaten + "\n\n")
			apples--
			if health < 100 {
				health = health + 10
			}
			fmt.Println("\nHealth is", health, "%\n\n")
			Sleep(2)
			Inventory(return_function)
		} else if choice == 2 {
			return_function()
		} else {
			fmt.Printf("\nInvalid input\n\n")
			Inventory(return_function)
		}
	}

}
