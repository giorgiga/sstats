package sstats

import (
	"math"
	"math/rand"
	"sort"
	"testing"
)

func TestMin_1(t *testing.T) {
	i,v := Min(0,1,2)
	if i != 0 { t.Errorf("wrong index") }
	if v != 0 { t.Errorf("wrong value") }
}

func TestMin_2(t *testing.T) {
	i,v := Min(2,1,0)
	if i != 2 { t.Errorf("wrong index") }
	if v != 0 { t.Errorf("wrong value") }
}

func TestMax_1(t *testing.T) {
	i,v := Max(0,1,2)
	if i != 2 { t.Errorf("wrong index") }
	if v != 2 { t.Errorf("wrong value") }
}

func TestMax_2(t *testing.T) {
	i,v := Max(2,1,0)
	if i != 0 { t.Errorf("wrong index") }
	if v != 2 { t.Errorf("wrong value") }
}

func TestMin_Eq(t *testing.T) {
	i,_ := Min(0,0,0)
	if i != 0 { t.Errorf("wrong index") }
}

func TestMax_Eq(t *testing.T) {
	i,_ := Max(0,0,0)
	if i != 0 { t.Errorf("wrong index") }
}

func TestMedian_Even(t *testing.T) {
	m := Median(0,2)
	if m != 1 { t.Errorf("wrong median (%v)", m) }
}

func TestMedian_Odd(t *testing.T) {
	m := Median(0,2,100)
	if m != 2 { t.Errorf("wrong median (%v)", m) }
}

func TestMedian_HandPicked(t *testing.T) {
	tcases := [][]float64{
		[]float64{ 1 },
		[]float64{ 1, 2 },
		[]float64{ 1, 2, 3 },
		[]float64{ 1, 2, 3, 4 },
		[]float64{ 0, 0, 1, 8, 9 },
		[]float64{ 0, 0, 0, 0, 1 },
		[]float64{ 0, 0, 0, 1, 1 },
		[]float64{ 0, 0, 1, 1, 1 },
		[]float64{ 0, 1, 1, 1, 1 },
		[]float64{ 0, 1, 3, 4, 2 },
	}
	for i,values := range tcases {
		actual := Median(values...)
		l := len(values)
		sort.Float64s(values)
		expected := math.NaN()
		if l % 2 == 0 {
			expected = (values[l/2] + values[l/2-1]) / 2
		} else {
			expected = values[l/2]
		}
		if expected != actual { t.Errorf("case %v: expected %v but got %v for %v", i+1, expected, actual, values) }
	}
}

func TestMedian_Rng(t *testing.T) {
	lengths := [...]int{ 6, 7, 300, 301, 1000, 1001 }
	for i,l := range lengths {
		rng := rand.New(rand.NewSource(int64(l)))
		values := make([]float64, l)
		for j := range values { values[j] = rng.Float64() }
		actual := Median(values...)
		sort.Float64s(values)
		expected := math.NaN()
		if l % 2 == 0 {
			expected = (values[l/2] + values[l/2-1]) / 2
		} else {
			expected = values[l/2]
		}
		if expected != actual { t.Errorf("case %v (%v): expected %v but got %v", i+1, l, expected, actual) }
	}
}

func TestMedianDoesntLoopForever(t *testing.T) {
	// in the initial implementation, this would loop forever
	m := Median(0,0,0,1,1)
	if m != 0 { t.Errorf("wrong median (%v)", m) }
}

type tCase struct{ v []float64; p int; ep int }

func TestQSortRoundR1(t *testing.T) {
	cases := []tCase{
		tCase{ []float64{ 4, 3, 2, 1, 0 }, 0, 4 },
		tCase{ []float64{ 4, 3, 2, 1, 0 }, 1, 3 },
		tCase{ []float64{ 4, 3, 2, 1, 0 }, 2, 2 },
		tCase{ []float64{ 4, 3, 2, 1, 0 }, 3, 1 },
		tCase{ []float64{ 4, 3, 2, 1, 0 }, 4, 0 },
		tCase{ []float64{ 0, 0, 0 }, 0, 2 },
		tCase{ []float64{ 0, 0, 1 }, 1, 1 },
		tCase{ []float64{ 0, 0, 1 }, 2, 2 },
	}

	for i,tc := range cases {
		ap := qSortRoundR(tc.v, tc.p)
		if tc.ep != ap {
			t.Errorf("case %v: wrong pivot position after sorting: expected %v got %v", i+1, tc.ep, ap)
		} else {
			pv := tc.v[ap]
			for _,v := range tc.v[:ap]   { if v >  pv { t.Errorf("%v >  %v left  of pivot in %v", v, pv, tc.v) } }
			for _,v := range tc.v[ap+1:] { if v <= pv { t.Errorf("%v <= %v right of pivot in %v", v, pv, tc.v) } }
		}
	}

}

// --- here be utilities -----------------------------------------------------------------------------------------------

func eq(a []float64, b []float64) bool {
	if a == nil && b == nil { return true }
	if a == nil || b == nil { return false }
	if len(a) != len(b) { return false }
	for i := range a { if a[i] != b[i] { return false } }
	return true
}
