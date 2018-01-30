package sstats

import (
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
