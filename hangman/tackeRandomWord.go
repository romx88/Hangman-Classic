package hangman

import (
	"math/rand"
	"time"
)

func TackeRandomWord(fileWords string) string {
	// Take words in tab
	tabWords := FetchFile(fileWords)
	// Random Words
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(len(tabWords))
	return tabWords[randomNum]
}
