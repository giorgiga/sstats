package sstats

import (
	"math"
)

// Min returns the min index and value in the given values.
// Returns (-1, NaN) if called with an empty slice.
func Min(values []float64) (int,float64) {
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
func Max(values []float64) (int,float64) {
	if len(values) == 0 { return -1, math.NaN() }
	mi := -1
	mv := math.Inf(-1)
	for i,v := range values {
		if v > mv { mv = v ; mi = i }
	}
	return mi,mv
}

// MinOf is a convenience, variadic version of Min().
func MinOf(values ...float64) (int,float64) {
	return Min(values)
}

// MaxOf is a convenience, variadic version of Max().
func MaxOf(values ...float64) (int,float64) {
	return Max(values)
}
