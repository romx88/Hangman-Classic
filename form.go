package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Party struct {
	Life        int
	WordsChoice string
	WordUser    []string
	TabOfLetter []string
	TabOfWords  []string
	Win         bool
}

func Form(p *Party) {
	trouve := false
	for p.Life != 10 && !p.Win {
		trouve = false
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if input == "" || input == "\n" || input[0:(len(input)-1)] == " " {
			fmt.Println("Invalid character, only letters are allowed")
		} else {
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
			if len(input[0:(len(input)-1)]) > 1 && !IfAlreadyTested(string(input[0:(len(input)-1)]), p.TabOfWords) {
				switch input[0:(len(input) - 1)] {
				case "STOP":
					SaveCurentParty(p)
					return
				case p.WordsChoice:
					p.Win = true
				default:
					p.TabOfWords = append(p.TabOfWords, input[0:(len(input)-1)])
					if p.Life == 9 {
						p.Life = 10
					} else {
						p.Life += 2
					}
				}
			} else {
				if !IfAlreadyTested(string(input[0]), p.TabOfLetter) {
					p.TabOfLetter = append(p.TabOfLetter, string(input[0]))
					for i := 0; i < len(p.WordsChoice); i++ {
						if string(input[0]) == string(p.WordsChoice[i]) {
							p.WordUser[i] = string(input[0])
							trouve = true
						}
					}
					if !trouve {
						p.Life++
					}
					for i := 0; i < len(p.WordsChoice); i++ {
						if string(p.WordsChoice[i]) != string(p.WordUser[i]) {
							break
						}
						if i == len(p.WordsChoice)-1 {
							p.Win = true
							break
						}
					}

				}
			}
		}
	}
}
