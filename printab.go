package hangman

import "fmt"

func Printab(tab []string) {
	for i := 0; i < len(tab); i++ {
		fmt.Printf(tab[i])
		fmt.Printf(" ")
	}
	fmt.Println("")
}
