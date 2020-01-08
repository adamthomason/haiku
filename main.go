package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ernestas-poskus/syllables"
)

// Dictionary holds syllables and words
type Dictionary struct {
	words map[int][]string
}

func main() {
	dictionary := generateDictionary()

	for count, words := range dictionary.words {
		fmt.Printf("%v:\n", count)

		for _, word := range words {
			fmt.Printf("    %v\n", word)
		}
	}
}

func generateDictionary() Dictionary {
	dictionary := Dictionary{}
	words := make(map[int][]string)

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count := syllables.CountSyllables([]byte(scanner.Text()))
		words[count] = append(words[count], scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	dictionary.words = words

	return dictionary
}
