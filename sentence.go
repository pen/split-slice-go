package split

import "strings"

// Sentence splits one-line string into N-lines minimal width.
func Sentence(sentence string, n int, greedy bool) string {
	words := strings.Split(sentence, " ")

	points := Slice(
		words, func(i int) int {
			return len(words[i]) + 1 // +1 for space separator
		},
		n, greedy,
	)
	nPart := len(points) - 1

	lines := make([]string, nPart)
	for i := 0; i < nPart; i++ {
		lines[i] = strings.Join(words[points[i]:points[i+1]], " ")
	}

	return strings.Join(lines, "\n")
}
