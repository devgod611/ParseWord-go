package parser

import (
	"log"
	"regexp"
	"sort"
	"strings"
)

// WordCount holds word and count pair
type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

// parse text and get most used top ten words
func ParseText(text string) []WordCount {

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(text, " ")
	words := strings.Split(processedString, " ")

	// count same words in s
	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	// create and fill slice of word-count pairs for sorting by count
	wordCounts := make([]WordCount, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, WordCount{Word: key, Count: val})
	}

	// sort wordCount slice by decreasing count number
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	if len(wordCounts) > 10 {
		wordCounts = wordCounts[0:11] // get top ten words only
	}

	return wordCounts
}
