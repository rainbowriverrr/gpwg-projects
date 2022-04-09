package main

import (
	"fmt"
	"regexp"
	"strings"
)

func testWords() {

	var testStr string = `As far as eye could reach he saw nothing but the stems of the great plants about him
	receding in the violet shade, and far overhead the multiple transparency of huge leaves
	filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever
	he felt able he ran again; the ground continued soft and springy, covered with the same
	resilient weed which was the first thing his hands had touched in Malacandra. Once or
	twice a small red creature scuttled across his path, but otherwise there seemed to be no
	life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned
	and alone in a forest of unknown vegetation thousands or millions of miles beyond the
	reach or knowledge of man.`

	countWordFreq(testStr)

}

func countWordFreq(input string) {

	lowered := strings.ToLower(input)
	cleaned := removePunctuation(lowered)
	wordList := strings.Split(cleaned, " ")

	frequency := make(map[string]int)

	for _, word := range wordList {
		word = strings.TrimSpace(word)
		frequency[word]++
	}

	for word, count := range frequency {
		if count > 1 {
			fmt.Printf("%s : %d times \n", word, count)
		}
	}

}

func removePunctuation(input string) string {
	reg, err := regexp.Compile("[^a-z]+")

	if err != nil {
		println("err")
	}

	processedString := reg.ReplaceAllString(input, " ")

	return processedString
}
