package hangman

import "fmt"

func IfAlreadyTested(wordsOrLettre string, tab []string) bool {
	if len(tab) >= 1 {
		for i := 0; i < len(tab); i++ {
			if tab[i] == wordsOrLettre {
				fmt.Println("Already tested, try another one")
				return true
			}
		}
	}
	return false
}
