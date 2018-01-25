package sstats

import (
	"math"
)

// Summary collects summary float64 statistics.
// The zero-point is not usable: use MakeSummary to properly initialize a Summary.
type Summary struct {
	min float64
	max float64
	count int64
	sum float64
	sumOfSquares float64
}

// MakeSummary initialises and returns a blank Summary.
func MakeSummary() Summary {
	s := Summary{}
	s.min = math.NaN()
	s.max = math.NaN()
	s.count = 0
	s.sum = 0
	s.sumOfSquares = 0
	return s
}

// Meet adds updates a Summary with a new datapoint.
func (s *Summary) Meet(datapoint float64) {
	if ! (s.min <= datapoint) { s.min = datapoint }
	if ! (s.max >= datapoint) { s.max = datapoint }
	s.count ++
	s.sum += datapoint
	s.sumOfSquares += math.Pow(datapoint, 2)
}

// Min returns the minimum datapoint met by this Summary, or NaN if the Summary is blank.
func (s *Summary) Min() float64 {
	return s.min
}

// Max returns the maximum datapoint met by this Summary, or NaN if the Summary is blank.
func (s *Summary) Max() float64 {
	return s.max
}

// Count returns the number of datapoints met by this Summary (0 if the Summary is blank).
func (s *Summary) Count() int64 {
	return s.count
}

// Mean returns the mean of the datapoints met by this Summary, or NaN if the Summary is blank.
func (s *Summary) Mean() float64 {
	return s.sum / float64(s.count)
}

// Variance returns the variance of the datapoints met by this Summary, or NaN if the Summary is blank.
func (s *Summary) Variance() float64 {
	count := float64(s.count)
	return (s.sumOfSquares - (math.Pow(s.sum,2) / count) ) / count
}

// StandardDeviation returns the standard deviation of the datapoints met by this Summary, or NaN if the Summary is blank.
func (s *Summary) StandardDeviation() float64 {
	return math.Sqrt( s.Variance() )
}
