package counter

import "strings"

// using string functions
func CountWords(sentence string) int {

	words := strings.Fields(sentence)
	return len(words)
}

// Custom Split function written from scratch
func CountWords2(sentence string) int {
	words := 0
	i := 0
	last := ""
	for i < len(sentence) {
		if sentence[i] == '\t' || sentence[i] == '\n' {
			words++
			i += 2
			last = ""
			continue
		}
		if sentence[i] == ' ' {
			words++
			last = ""
			for i < len(sentence) && sentence[i] == ' ' {
				i++
			}
		} else {
			last = last + string(sentence[i])
		}
		i++
	}
	if len(last) > 0 {
		words++
	}
	return words
}
