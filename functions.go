package sstats

import (
	"math"
)

// Min returns the min index and value in the given slice.
// Returns (-1, NaN) if called with an empty slice.
func Min(slice []float64) (int,float64) {
	if len(slice) == 0 { return -1, math.NaN() }
	mi := -1
	mv := math.Inf(1)
	for i,v := range slice {
		if v < mv { mv = v ; mi = i }
	}
	return mi,mv
}

// Max returns the max index and value in the given slice.
// Returns (-1, NaN) if called with an empty slice.
func Max(slice []float64) (int,float64) {
	if len(slice) == 0 { return -1, math.NaN() }
	mi := -1
	mv := math.Inf(-1)
	for i,v := range slice {
		if v > mv { mv = v ; mi = i }
	}
	return mi,mv
}
