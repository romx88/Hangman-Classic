package hangman

import (
	"bufio"
	"log"
	"os"
)

func FetchFile(file string) []string {
	var tab []string
	data, err := os.Open(file)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
		os.Exit(1)
	}
	defer data.Close()
	fileScanner := bufio.NewScanner(data)
	for fileScanner.Scan() {
		tab = append(tab, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
		os.Exit(1)
	}
	return tab
}
