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
	fmt.Printf("%s\n%s\n%s", dictionary.generateLine(5), dictionary.generateLine(7), dictionary.generateLine(5))
}

func (dictionary *Dictionary) generateLine(syllableCount int) string {
	var remainingCount = syllableCount
	line := ""

	for remainingCount > 0 {
		time.Sleep(time.Millisecond * 3)

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		r1.Seed(time.Now().UnixNano())

		randomSyllable := r1.Intn(remainingCount) + 1

		time.Sleep(time.Millisecond * 3)

		s2 := rand.NewSource(time.Now().UnixNano())
		r2 := rand.New(s2)
		r2.Seed(time.Now().UnixNano())

		max := len(dictionary.words[randomSyllable])
		randomWord := r2.Intn(max)

		line = line + dictionary.words[randomSyllable][randomWord] + " "

		remainingCount -= randomSyllable
	}

	return ucFirst(line)
}
