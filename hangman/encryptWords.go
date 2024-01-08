package hangman

import (
	"math/rand"
	"time"
)

func EncryptWord(wordschoice string) ([]string, []string) {
	var tabOfLetter []string
	var wordUser []string
	var randindex []int
	rand.Seed(time.Now().UnixNano())
	nbLettre := (len(wordschoice) / 2) - 1
	if nbLettre == 0 {
		nbLettre = 1
	}
	for i := 0; i <= nbLettre; i++ {
		randindex = append(randindex, rand.Intn(len(wordschoice)-1))
	}
	for i := 0; i < len(wordschoice); i++ {
		wordUser = append(wordUser, "_")
		for c := 0; c < nbLettre; c++ {
			if i == randindex[c] {
				new := true
				wordUser[i] = string(wordschoice[i])
				for f := 0; f < len(tabOfLetter); f++ {
					new = true
					if string(wordschoice[i]) == tabOfLetter[f] {
						new = false
					}
				}
				if new {
					tabOfLetter = append(tabOfLetter, string(wordschoice[i]))
				}
			}
		}
	}

	for c := 0; c < len(randindex); c++ {
		temprang := randindex[c]
		for i := 0; i < len(wordUser); i++ {
			if string(wordUser[temprang]) == string(wordschoice[i]) {
				wordUser[i] = string(wordUser[temprang])

			}
		}
	}

	return wordUser, tabOfLetter
}
