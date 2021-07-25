package main

import (
	"testing"
	"math"
)

func TestTimestep(t *testing.T) {
	planets := NewSolar()
	x_e_0 := (*planets)[1].Pos[0]
	n := int(365/4)
	for i:=0;i<n;i++ {
		timestep(planets,float32(3600*24*365/n))
	}
	x_e_1 := (*planets)[1].Pos[0]
	dist := math.Abs(float64(x_e_1 -x_e_0))
	if dist > 1e5 {
		t.Errorf("deviation too large %g m", dist)
	}
}

