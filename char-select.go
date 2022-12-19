package main

import "fmt"

// create choice and character variables
var choice int
var char string

func Introduce() {
	//introduce the game and players
	fmt.Printf("Welcome to this text based adventure game!\n")
	Sleep(4)
	fmt.Printf("Meet the characters\n\n")
	Sleep(2)
	fmt.Printf("       The Archer\n" + archer + "\n\n")
	Sleep(2)
	fmt.Printf("     Dark Knight\n" + dark_knight + "\n\n")
	Sleep(2)
	fmt.Printf("   Barbara\n\n" + barbara + "\n\n")
	Sleep(2)
	Char_select()
}

func Char_select() {
	//ask player for character choice
	fmt.Printf("1. The Archer\n2. Dark Knight\n3. Barbara\n\nChoose your hero (1, 2 or 3): ")

	//grab input from player
	fmt.Scanln(&choice)

	//if player chooses 1 set character to archer
	if choice == 1 {
		char = "The Archer"

		//if player chooses 2 set character to Dark Knight
	} else if choice == 2 {
		char = "Dark Knight"

		//if player chooses 3 set character to barbara
	} else if choice == 3 {
		char = "Barbara"

		//if input invalid tell player to choose again
	} else {
		fmt.Printf("invalid input, try again: \n\n")
		//return player to character select function
		Char_select()
	}
	fmt.Printf("\n" + char + ", great choice!\n")
	Char_hello_sound(char)
	fmt.Printf("Let the games begin........\n\n")
	Start_game_sound()
	fmt.Printf(alarm)
	Alarm_sound()
	fmt.Printf("\n\nYou open your eyes and realise its your alarm, it's time to get up!\n")
	Sleep(3)
	fmt.Printf("With a big stretch you jump out of bed.\n")
	Yawn_sound()
	Sleep(2)
	fmt.Printf("\n" + bed + "\n\n")
	Sleep(2)
	Home()
}
