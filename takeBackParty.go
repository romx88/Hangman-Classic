package hangman

import (
	"encoding/json"
	"fmt"
	"os"
)

func TakeBackParty(file string) Party {
	// Lisez le contenu du fichier
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier : %v\n", err)
		os.Exit(1)
	}
	// Décodez les données JSON
	var p Party
	if err := json.Unmarshal(data, &p); err != nil {
		fmt.Printf("Erreur lors du décodage des données JSON : %v\n", err)

	}
	return p
}
