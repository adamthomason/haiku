package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	"unicode"

	"github.com/ernestas-poskus/syllables"
)

// Dictionary holds syllables and words
type Dictionary struct {
	words map[int][]string
}

// Poem contains the poem structure
type Poem struct {
	First, Second, Third string
}

func main() {
	http.HandleFunc("/", healthcheck)
	http.HandleFunc("/poem", poemHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok."))
}

func poemHandler(w http.ResponseWriter, r *http.Request) {
	dictionary := generateDictionary()
	poem := dictionary.generateHaiku()
	result, _ := json.Marshal(&poem)

	fmt.Println(string(result))

	w.Write(result)
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

func (dictionary *Dictionary) generateHaiku() Poem {
	poem := Poem{
		dictionary.ensureSyllables(dictionary.generateLine(5), 5),
		dictionary.ensureSyllables(dictionary.generateLine(7), 7),
		dictionary.ensureSyllables(dictionary.generateLine(5), 5),
	}

	return poem
}

func (dictionary *Dictionary) ensureSyllables(line string, length int) string {
	checkedLine := line

	for (syllables.CountSyllables([]byte(checkedLine))) != length {
		checkedLine = dictionary.generateLine(length)
	}

	return checkedLine
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
