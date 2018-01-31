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
	if (sz == 0) { return math.NaN() }
	i := sz/2 // if even, we want the mean of indexes i and i-1 (eg: size is 4 -> we want indexes 1 and 2 and i = 2)
	          // if odd, we just want index i (eg: size is 5 -> we want index 3 and i = 3)
	return _median(values, i, sz % 2 == 0)
}

func _median(values []float64, i int, domean bool) (float64) { // domean: if we should return the mean of [i] and [i-1]
	// fmt.Printf("looking for #%v in %v (domean? %v)\n", i, values, domean)
	sz := len(values)

	// let's pick a pivot and run a round of quick sort

	pi := sz/2 // pivot index (will be overwritten)

	{
		pv := values[pi] // pivot value, for reference
		values[pi], values[sz-1] = values[sz-1], values[pi] // park pivot in the rightmost position
		l := 0 // initial index for left
		for i, x := range values {
			if x < pv { // elements smaller than the pivot go to the left
				values[i], values[l] = values[l], values[i]
				l++
			}
		}
		pi = l
		values[pi], values[sz-1] = values[sz-1], values[pi] // move the pivot after the last "left" element
	}
	// at this point the pivot has pi elements smaller than itself, and they are all to its left

	if i < pi {
		// both i and i-1 are to the left of pi -> discard from pi rightwards and recurse
		return _median(values[:pi], i, domean)
	} else if i > pi {
		// i is to the right of pi and i-1 is not left of pi (at worst, it's pi) -> discard leftwards of pi and recurse
		return _median(values[pi:], i - pi, domean)
	} else { // i == pi
		if domean {
				// we must search for the value at i-1
				// we could recurse _median(values, i-1), but really
				// what we want is in fact the max value to the left of i
				_,prev := Max(values[:i]...)
				return ( values[i] + prev ) / 2
		} else {
				return values[pi]
		}
	}

	panic("reached unreachable code")
}
