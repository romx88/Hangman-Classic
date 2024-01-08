package hangman

import "fmt"

func Help() {
	fmt.Println("Hangman is a game of hangman :) ")
	fmt.Println("To start the game, run the following command:")
	fmt.Println("If you don't want ASCII art, use the following commands:")
	fmt.Println("go run main/main.go [word file]")
	fmt.Println("")
	fmt.Println("Otherwise, run the following command:")
	fmt.Println("go run main/main.go [word file] [letter file]")
	fmt.Println("")
	fmt.Println("To save a game, run the following command:")
	fmt.Println("Enter 'STOP' when prompted to enter a letter, then follow the instructions.")
	fmt.Println("")
	fmt.Println("To resume a game, run the following command:")
	fmt.Println("go run main/main.go --startWith [save file path]")
	fmt.Println("")

}
