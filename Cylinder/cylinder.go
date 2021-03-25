package main

import (
	"math"
)

type CylinderAtt struct {
	rad    float64
	height float64
}

func SurfaceArea() float64 {
	var cylinder CylinderAtt

	return 2.0 * math.Pi * cylinder.rad * (cylinder.rad + cylinder.height)
}
