package split_test

import (
	"testing"

	split "github.com/pen/split-slice-go"
)

func TestSentence(t *testing.T) {
	t.Parallel()

	t.Run("split.Sentence", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			title    string
			sentence string
			nPart    int
			greedy   bool
			want     string
		}{
			{
				"nulstring",
				"", 3, false,
				"",
			},
			{
				"split one word",
				"word", 3, false,
				"word",
			},
			{
				"normal",
				"the sun shines blight on the old kentucky home", 3, false,
				"the sun shines\nblight on the\nold kentucky home",
			},
			{
				"greedy",
				"the sun shines blight on the old kentucky home", 3, true,
				"the sun shines\nblight on the old\nkentucky home",
			},
		}

		for _, tc := range testCases {
			got := split.Sentence(tc.sentence, tc.nPart, tc.greedy)
			if got != tc.want {
				t.Errorf("\n\"%s\"\n-- got --\n%s\n-- want --\n%s", tc.title, got, tc.want)
			}
		}
	})
}
