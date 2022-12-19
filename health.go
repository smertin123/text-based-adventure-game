package main

import (
	"fmt"
	"os"
)

var health = 100

func lose_10() {
	health = health - 10
	Health_check()
	fmt.Println("Your health goes down to", health, "%\n")
}

func Health_check() {
	if health == 0 {
		fmt.Printf("\nYou Died\n\n")
		Sleep(1)
		fmt.Printf("         _                                                 _.--.    __  _\n	| |                                               ) \\   `.,'  \\| |\n	| |                                            (`'       |     : |\n	| |                             _..-.-.-.-._    )     ,),'.    | |\n	| |('.                    __..-' ) ) ) ) ) )``-'    _.-| \\     | |\n	| | \\ `...------''``--'''' \\   )_____....---     ,''           ; |\n	| |_(_..-......_________..._,-'_,..__....____..-'.._________..'| |\n        | |____________________________________________________________| |\n      __|_|____________________________________________________________|_|__\n\n")
		Dead_sound()
		fmt.Printf("Restart?\n\n1. Yes\n2. No\n\n")
		fmt.Scanln(&choice)
		if choice == 1 {
			main()
		} else if choice == 2 {
			os.Exit(0)
		} else {
			fmt.Printf("Invalid input\n")
			fmt.Printf("Restart?\n\n1. Yes\n2. No\n\n")
			fmt.Scanln(&choice)
		}
	}
}
