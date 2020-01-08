package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"unicode"

	"github.com/ernestas-poskus/syllables"
)

// Dictionary holds syllables and words
type Dictionary struct {
	words map[int][]string
}

func main() {
	dictionary := generateDictionary()

	dictionary.generateHaiku()
}

func ucFirst(phrase string) string {
	for i, v := range phrase {
		return string(unicode.ToUpper(v)) + phrase[i+1:]
	}

	return ""
}

func generateDictionary() Dictionary {
	dictionary := Dictionary{}
	words := make(map[int][]string, 0)

	file, err := os.Open("dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count := syllables.CountSyllables([]byte(scanner.Text()))

		if count <= 7 {
			words[count] = append(words[count], scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	dictionary.words = words

	return dictionary
}

func (dictionary *Dictionary) generateHaiku() {
	fmt.Println(dictionary.generateLine(5))
	fmt.Println(dictionary.generateLine(7))
	fmt.Println(dictionary.generateLine(5))
}

func (dictionary *Dictionary) generateLine(syllableCount int) string {
	var remainingCount = syllableCount
	line := ""

	for remainingCount > 0 {
		time.Sleep(time.Millisecond * 300)
		rand.Seed(time.Now().UnixNano())
		randomSyllable := rand.Intn(remainingCount) + 1

		max := len(dictionary.words[randomSyllable])

		rand.Seed(time.Now().UnixNano())
		randomWord := rand.Intn(max)

		line = line + dictionary.words[randomSyllable][randomWord] + " "

		remainingCount -= randomSyllable
	}

	return ucFirst(line)
}
