// count the occurrence of words in a text

package main

import "strings"

func WordCount(text string) map[string]int {
	counts := map[string]int{}
	words := strings.Fields(text)
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	return counts
}
