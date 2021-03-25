package Cylinder

import (
	"fmt"
	"math"
)

func SurfaceArea() float64 {
	var rad float64
	var height float64

	fmt.Print("Input radius: ")
	fmt.Scanln(&rad)
	fmt.Print("Input height: ")
	fmt.Scanln(&height)

	return 2.0 * math.Pi * rad * (rad + height)
}
