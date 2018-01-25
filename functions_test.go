package sstats

import (
	"testing"
)

func TestMin1(t *testing.T) {
	i,v := Min([]float64{0,1,2})
	if i != 0 { t.Errorf("wrong index") }
	if v != 0 { t.Errorf("wrong value") }
}

func TestMin2(t *testing.T) {
	i,v := Min([]float64{2,1,0})
	if i != 2 { t.Errorf("wrong index") }
	if v != 0 { t.Errorf("wrong value") }
}

func TestMax1(t *testing.T) {
	i,v := Max([]float64{0,1,2})
	if i != 2 { t.Errorf("wrong index") }
	if v != 2 { t.Errorf("wrong value") }
}

func TestMax2(t *testing.T) {
	i,v := Max([]float64{2,1,0})
	if i != 0 { t.Errorf("wrong index") }
	if v != 2 { t.Errorf("wrong value") }
}
