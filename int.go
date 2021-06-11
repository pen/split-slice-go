// Package split provides algorithm to split array of value into sub-arrays which have near sum of values.
package split

// IntSlice calculates best points to split.
// Returns []int{0, ...(points)..., len(intSlice)} for convenience.
func IntSlice(intSlice []int, nPart int, greedy bool) []int {
	if nPart <= 0 || len(intSlice) < nPart {
		return []int{0, len(intSlice)}
	}

	s := newState(nPart, intSlice, greedy)
	s.search(nPart, 0)

	return s.bestSplitPoints()
}

// state is workspace object for a calculation.
type state struct {
	slice   []int // working slice
	indexOf []int // candidate split points (reversed)

	// best record
	minIndexOf  []int
	minPartSize int

	// table to calculate slice[i:j].sum() fast
	totalUpto []int

	greedy bool
}

// newState is kinda constructor of state
func newState(nPart int, slice []int, greedy bool) *state {
	indexOf := make([]int, nPart+1)
	indexOf[0] = len(slice)

	if greedy {
		copySlice := make([]int, len(slice))
		copy(copySlice, slice)
		reverseInt(copySlice)
		slice = copySlice
	}

	totalUpto := make([]int, len(slice)+1)

	total := 0
	for i, m := range slice {
		total += m
		totalUpto[i+1] = total
	}

	return &state{
		slice:       slice,
		indexOf:     indexOf,
		totalUpto:   totalUpto,
		minPartSize: total,
		minIndexOf:  make([]int, len(indexOf)),
		greedy:      greedy,
	}
}

// bestSplitPoints makes result of exported function IntSlice.
func (s *state) bestSplitPoints() []int {
	// recycle s.indexOf that finished its original role.
	copy(s.indexOf, s.minIndexOf)

	if s.greedy {
		for i, mi := range s.minIndexOf {
			s.indexOf[i] = len(s.slice) - mi
		}
	} else {
		reverseInt(s.indexOf)
	}

	return s.indexOf
}

// function search examines combinations of split points.
/*
   ex. split.IntSlice([3, 3, 5, 3, 4, 2], 3, false)

   move splitting points   max sum
   --1---2---3---4---5-----------------------------
   3 | 3 | 5   3   4   2  => 14
   3 | 3   5 | 3   4   2  =>  9    new record
   3 | 3   5   3 | 4   2  break cause 3+5+3 > 9
   3 | 3   5   3   4 | 2  (skip)
   3   3 | 5 | 3   4   2  =>  9    tie (do nothing)
   3   3 | 5   3 | 4   2  =>  8    new record 8
   3   3 | 5   3   4 | 2  break cause 5+3+4 > 8
   3   3   5 | 3 | 4   2  break cause 3+3+5 > 8
   3   3   5 | 3   4 | 2  (skip)
   3   3   5   3 | 4 | 2  (skip)

   it returns [ 0, 2, 4, 6 ]
   which can make [ [ 3, 3 ], [ 5, 3 ], [ 4, 2 ] ])
*/
func (s *state) search(nRest, index int) {
	if nRest == 0 {
		if max := s.maxPartSize(); max <= s.minPartSize {
			s.recordMin(max)
		}

		return
	}

	s.indexOf[nRest] = index

	// move a splitting point
	for nextIndex := index + 1; nextIndex <= len(s.slice)-nRest; nextIndex++ {
		if s.partSize(index, nextIndex) >= s.minPartSize {
			// no hope after current part has grew too large
			break
		}

		s.search(nRest-1, nextIndex)
	}
}

// recordMin keeps best candidate so far.
func (s *state) recordMin(partSize int) {
	s.minPartSize = partSize
	copy(s.minIndexOf, s.indexOf)
}

// partSize is like slice[start:next].sum() .
func (s *state) partSize(startIndex, nextIndex int) int {
	return s.totalUpto[nextIndex] - s.totalUpto[startIndex]
}

// maxPartSize is score of a candidate. lower is better.
func (s *state) maxPartSize() int {
	max := 0
	for i := 0; i < len(s.indexOf)-1; i++ {
		if size := s.partSize(s.indexOf[i+1], s.indexOf[i]); size > max {
			max = size
		}
	}

	return max
}

// reverseInt is a utility to reverse []int .
func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
