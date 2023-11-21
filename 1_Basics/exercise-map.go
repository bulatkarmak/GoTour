// https://go.dev/tour/moretypes/23

package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	wordCountMap := make(map[string]int)
	wordsSlice := strings.Fields(s)

	for _, word := range wordsSlice {
		wordCountMap[word]++
	}

	return wordCountMap
}

// func main() {
// 	fmt.Println(WordCount("hello hello white black"))

// }
