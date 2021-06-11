package split

import (
	"strings"
)

// Sentence splits one-line string into n-lines string.
func Sentence(sentence string, n int, greedy bool) string {
	words := strings.Split(sentence, " ")

	widths := make([]int, len(words))
	for i, word := range words {
		widths[i] = len(word) + 1 // +1 for trailing space
	}

	points := IntSlice(widths, n, greedy)
	nPart := len(points) - 1

	lines := make([]string, nPart)
	for i := 0; i < nPart; i++ {
		lines[i] = strings.Join(words[points[i]:points[i+1]], " ")
	}

	return strings.Join(lines, "\n")
}
