package sstats

import (
	"fmt"
	"math"
	"testing"
)

func TestEmpty(t *testing.T) {
	s := MakeSummary()
	if !math.IsNaN(s.Min())               { t.Fail() }
	if !math.IsNaN(s.Max())               { t.Fail() }
	if s.Count() != 0                     { t.Fail() }
	if !math.IsNaN(s.Mean())              { t.Fail() }
	if !math.IsNaN(s.Variance())          { t.Fail() }
	if !math.IsNaN(s.StandardDeviation()) { t.Fail() }
}

func TestSimplest(t *testing.T) {
	s := MakeSummary()
	s.Meet(1)
	s.checkStats(t, 1, 1, 1, 1, 0)
}

func TestSimple(t *testing.T) {
	s := MakeSummary()
	s.Meet(8)
	s.Meet(10)
	s.Meet(12)
	s.checkStats(t, 8, 12, 3, 10, 8.0/3)
}

// examples

func ExampleSimple() {
	var summary Summary = MakeSummary()
	summary.Meet(8)
	summary.Meet(10)
	summary.Meet(12)
	fmt.Printf("mean: %v\n", summary.Mean())
	fmt.Printf("variance: %v\n", summary.Variance())
	// Output:
	// mean: 10
	// variance: 2.6666666666666665
}

// --- here be utilities ---

func (s *Summary) checkStats(t *testing.T, min float64, max float64, count int64, mean float64, variance float64) {
	pass := true
	// yes, floats should not be compared with ==, but... it happens to work with these tests :)
	pass = pass && s.Min() == min
	pass = pass && s.Max() == max
	pass = pass && s.Count() == count
	pass = pass && s.Mean() == mean
	pass = pass && s.Variance() == variance
	if !pass {
		t.Errorf("min: %v %v, max: %v %v, count: %v %v, mean: %v %v, variance: %v %v",
		         min, s.Min(),
		         max, s.Max(),
		         count, s.Count(),
		         mean, s.Mean(),
		         variance, s.Variance())
	}
}
