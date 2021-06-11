package split_test

import (
	"reflect"
	"testing"

	split "github.com/pen/split-slice-go"
)

func TestInt(t *testing.T) {
	t.Parallel()

	t.Run("split.IntSlice", func(t *testing.T) {
		t.Parallel()

		testCases := []struct {
			title    string
			sequence []int
			nPart    int
			greedy   bool
			want     []int
		}{
			{
				"normal: 5|55|55",
				[]int{5, 5, 5, 5, 5},
				3, false,
				[]int{0, 1, 3, 5},
			},
			{
				"greedy: 55|55|5",
				[]int{5, 5, 5, 5, 5},
				3, true,
				[]int{0, 2, 4, 5},
			},
		}

		for _, tc := range testCases {
			got := split.IntSlice(tc.sequence, tc.nPart, tc.greedy)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\n\"%s\"\n-- got --\n%v\n-- want --\n%v", tc.title, got, tc.want)
			}
		}
	})
}
