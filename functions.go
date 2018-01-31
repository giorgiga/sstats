package sstats

import (
	"math"
)

// Min returns the min index and value in the given values.
// Returns (-1, NaN) if called with an empty slice.
func Min(values ...float64) (int,float64) {
	if len(values) == 0 { return -1, math.NaN() }
	mi := -1
	mv := math.Inf(1)
	for i,v := range values {
		if v < mv { mv = v ; mi = i }
	}
	return mi,mv
}

// Max returns the max index and value in the given values.
// Returns (-1, NaN) if called with an empty slice.
func Max(values ...float64) (int,float64) {
	if len(values) == 0 { return -1, math.NaN() }
	mi := -1
	mv := math.Inf(-1)
	for i,v := range values {
		if v > mv { mv = v ; mi = i }
	}
	return mi,mv
}

// Median returns the median of the given values.
// Returns NaN if called with an empty slice.
func Median(values... float64) float64 {
	sz := len(values)
	switch sz {
		case 0:  return math.NaN()
		case 1:  return values[0]
		default: return findIth(values, sz/2, sz % 2 == 0)
	}
}

// the domean parameter holds wheter  initially we had an even or odd number of values:
// if size is even (domean = true), we want the mean of indexes i and i-1 (eg: size is 4 -> i = 2 but we want them mean of indexes 1 and 2)
// if size is odd (domean = false), we just want index i (eg: size is 5 -> we just want index 3) )
func findIth(values []float64, i int, domean bool) (float64) {
	pi := qSortRoundR(values, len(values)/2)
	// the pivot values[pi] now has pi elements <= than itself, all to its left
	//    ie: values is arranged like [<<<<][====P][>>>>]
	if i < pi { // both i and i-1 are to the left of pi
		return findIth(values[:pi], i, domean) // discard from pi rightwards and recurse
	} else if pi < i { // i is to the right of pi and i-1 is not left of pi (at worst, it's pi)
		return findIth(values[pi:], i - pi, domean)  // discard leftwards of pi (keep pi) and recurse
	} else if domean { // i == pi, but we must search for the value at i-1 too
		// we could recurse:
		//     prev = findIth(values, i-1)
		// but what we want is in fact the max value to the left of i
		_,prev := Max(values[:i]...)
		return ( values[i] + prev ) / 2
	} else {  // i == pi and we don't care about i-1
		return values[i]
	}
}

// qSortRoundR sorts "big" elements to the right of the pivot
//    ie: it arranges values like [<<<<][====P][>>>>]
// Takes the index of the pivot and returns its new index after the sort.
func qSortRoundR(values []float64, pivot int) (int) {
	// park pivot in the leftmost position
	values[pivot], values[0] = values[0], values[pivot]
	right := len(values) -1
	for i := right; i > 0; i-- {
		if values[i] > values[0] {
			values[i], values[right] = values[right], values[i]
			right--
		}
	}
	// move pivot to position right (1 left of where the last "big" element was moved)
	values[right], values[0] = values[0], values[right]
	return right
}
