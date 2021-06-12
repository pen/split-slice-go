// Package split provides algorithm to split array of value into sub-arrays which have near sum of values.
package split

import "reflect"

// Slice calculates best points to split.
// Returns []int{0, ...(points)..., len(slice)} for convenience.
func Slice(slice interface{}, getSize func(i int) int, nPart int, greedy bool) []int {
	sliceLen := reflect.ValueOf(slice).Len()
	if sliceLen == 0 || nPart <= 0 {
		return nil
	}

	if sliceLen <= nPart {
		splitPoints := make([]int, sliceLen+1)
		for i := 0; i <= sliceLen; i++ {
			splitPoints[i] = i
		}

		return splitPoints
	}

	s := newState(sliceLen, getSize, nPart, greedy)
	s.search(nPart, 0)

	return s.bestSplitPoints()
}

// state is workspace object for a calculation.
type state struct {
	sliceLen int
	greedy   bool

	// candidate split points (reversed)
	indexOf []int

	// best record
	minIndexOf  []int
	minPartSize int

	// table to calculate slice[i:j].sum() fast
	totalUpto []int
}

// newState is kinda constructor of state
func newState(sliceLen int, getSize func(i int) int, nPart int, greedy bool) *state {
	indexOf := make([]int, nPart+1)
	indexOf[0] = sliceLen

	totalUpto := make([]int, sliceLen+1)
	total := 0

	for i := 0; i < sliceLen; i++ {
		j := i
		if greedy {
			j = sliceLen - 1 - i
		}

		total += getSize(j)
		totalUpto[i+1] = total
	}

	return &state{
		sliceLen:    sliceLen,
		greedy:      greedy,
		indexOf:     indexOf,
		minIndexOf:  make([]int, len(indexOf)),
		minPartSize: total,
		totalUpto:   totalUpto,
	}
}

// bestSplitPoints makes result of exported function Slice.
func (s *state) bestSplitPoints() []int {
	// recycle s.indexOf that finished its original role.
	copy(s.indexOf, s.minIndexOf)

	if s.greedy {
		for i, mi := range s.minIndexOf {
			s.indexOf[i] = s.sliceLen - mi
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
	for nextIndex := index + 1; nextIndex <= s.sliceLen-nRest; nextIndex++ {
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

// partSize is like slice[start:next].sum() .
func (s *state) partSize(startIndex, nextIndex int) int {
	return s.totalUpto[nextIndex] - s.totalUpto[startIndex]
}

// reverseInt is a utility to reverse []int .
func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
