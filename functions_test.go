package sstats

import (
	"math"
	"math/rand"
	"sort"
	"testing"
)

func TestMin1(t *testing.T) {
	i,v := Min(0,1,2)
	if i != 0 { t.Errorf("wrong index") }
	if v != 0 { t.Errorf("wrong value") }
}

func TestMin2(t *testing.T) {
	i,v := Min(2,1,0)
	if i != 2 { t.Errorf("wrong index") }
	if v != 0 { t.Errorf("wrong value") }
}

func TestMax1(t *testing.T) {
	i,v := Max(0,1,2)
	if i != 2 { t.Errorf("wrong index") }
	if v != 2 { t.Errorf("wrong value") }
}

func TestMax2(t *testing.T) {
	i,v := Max(2,1,0)
	if i != 0 { t.Errorf("wrong index") }
	if v != 2 { t.Errorf("wrong value") }
}

func TestMinEq(t *testing.T) {
	i,_ := Min(0,0,0)
	if i != 0 { t.Errorf("wrong index") }
}

func TestMaxEq(t *testing.T) {
	i,_ := Max(0,0,0)
	if i != 0 { t.Errorf("wrong index") }
}

func TestMedianEven(t *testing.T) {
	m := Median(0,2)
	if m != 1 { t.Errorf("wrong median (%v)", m) }
}

func TestMedianOdd(t *testing.T) {
	m := Median(0,2,100)
	if m != 2 { t.Errorf("wrong median (%v)", m) }
}

func TestMedianRandom(t *testing.T) {
	lengths := [...]int{ 30, 31, 1000, 1001 }
	rng := rand.New(rand.NewSource(0))
	for tcase,l := range lengths {
		values := make([]float64, l)
		for i := range values { values[i] = rng.Float64() }
		actual := Median(values...)
		sort.Float64s(values)
		expected := math.NaN()
		if l % 2 == 0 {
			expected = (values[l/2] + values[l/2-1]) / 2
		} else {
			expected = values[l/2]
		}
		if expected != actual { t.Errorf("case %v: expected %v but got %v", tcase+1, expected, actual) }
	}
}

