package main

import "fmt"

var draw_apple = 2

func Home() {
	fmt.Printf("\nWhat do you want to do?\n1. Look around\n2. Leave\n0. View inventory\n\n")
	fmt.Scanln(&choice)
	if choice == 1 {
		Look_around()
	} else if choice == 2 {
		fmt.Printf("You arrive in town.\n")
		Sleep(2)
		Scene(town)
		Sleep(2)
		Town()
	} else if choice == 0 {
		Inventory(Home)
	} else {
		fmt.Printf("invalid input, try again: \n\n")
		Home()
	}
}

func Look_around() {
	fmt.Printf("\n1. Check birdcage\n2. Check dresser\n0. View inventory\n")
	fmt.Scanln(&choice)
	if choice == 1 {
		Scene(birdcage)
		Birdcage()
	} else if choice == 2 {
		Scene(dresser)
		Dresser()
	} else if choice == 0 {
		Inventory(Look_around)
	} else {
		fmt.Printf("invalid input, try again: \n\n")
		Look_around()
	}
}

func Birdcage() {
	fmt.Printf("\n1. Check bird\n2. Leave\n0. View inventory\n")
	fmt.Scanln(&choice)
	if choice == 1 {
		fmt.Printf(bird)
		Bird_sound()
		Bird()
	} else if choice == 2 {
		Home()
	} else if choice == 0 {
		Inventory(Birdcage)
	} else {
		fmt.Printf("invalid input, try again: \n\n")
		Birdcage()
	}
}

func Bird() {
	fmt.Printf("\n1. Stroke bird\n2. Leave\n0. View inventory\n")
	fmt.Scanln(&choice)
	if choice == 1 {
		fmt.Printf(bird)
		fmt.Printf("\nHe bites your finger\n\n")
		Ouch_sound()
		lose_10()
		Bird()
	} else if choice == 2 {
		Scene(birdcage)
		Birdcage()
	} else if choice == 0 {
		Inventory(Bird)
	} else {
		fmt.Printf("\nInvalid input\n\n")
		Bird()
	}
}

func Dresser() {
	fmt.Printf("\n1. Look in mirror\n2. Check drawer\n3. Leave\n0. View inventory\n")
	fmt.Scanln(&choice)
	if choice == 1 {
		Mirror()
	} else if choice == 2 {
		Scene(drawer)
		Drawer()
	} else if choice == 3 {
		Home()
	} else if choice == 0 {
		Inventory(Dresser)
	} else {
		fmt.Printf("invalid input, try again: \n\n")
		Scene(dresser)
		Dresser()
	}
}

func Mirror() {
	if char == "The Archer" {
		fmt.Printf(archer + "\n\n")
	} else if char == "Dark Knight" {
		fmt.Printf(dark_knight + "\n\n")
	} else if char == "Barbara" {
		fmt.Printf(barbara + "\n\n")
	} else {
		fmt.Printf("No character" + "\n\n")
	}
	Dresser()
}

func Drawer() {
	fmt.Printf("\n1. Take apple\n2. Leave\n0. View inventory\n")
	fmt.Scanln(&choice)
	if choice == 1 && draw_apple > 0 {
		apples++
		draw_apple--
		fmt.Printf(apple+"\n\n You have %d apples in your inventory\n\n", apples)
		Drawer()
	} else if choice == 1 && draw_apple < 1 {
		fmt.Printf("\nYou've taken them all you greedy bastard\n\n")
		Crowd_shocked_sound()
		Drawer()
	} else if choice == 0 {
		Inventory(Drawer)
	} else if choice == 2 {
		Scene(dresser)
		Dresser()
	} else {
		fmt.Printf("Invalid input")
		Scene(drawer)
		Drawer()
	}
}
