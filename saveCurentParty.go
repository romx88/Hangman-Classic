package hangman

import (
	// "encoding/json"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func SaveCurentParty(p *Party) {
	fmt.Println("The game has ended. In which file in the 'data/saves' directory do you want to save your game? \n (ex: save.txt) : ")
	reader := bufio.NewReader(os.Stdin)
	file, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	// // Encode the game state to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error encoding game state:", err)
		return
	}
	nameFile := "data/saves/"
	nameFile += file[0:(len(file) - 1)]
	fmt.Println(nameFile)
	// Write the JSON data to a file
	err = os.WriteFile(nameFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing :", err)
		return
	}
	fmt.Printf("Game state saved to %v", file)
}
