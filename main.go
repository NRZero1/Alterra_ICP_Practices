package main

import (
	"fmt"

	"github.com/NRZero1/Alterra_ICP_Practices/Cylinder"
	"github.com/NRZero1/Alterra_ICP_Practices/Grade"
	"github.com/NRZero1/Alterra_ICP_Practices/Sum"
)

func main() {
	fmt.Println("Challenge 1: Cylinder Surface Area")
	fmt.Println(Cylinder.SurfaceArea())
	fmt.Print("\n")
	fmt.Println("Challenge 2: Checking Grade")
	fmt.Println(Grade.CheckGrade())
	fmt.Print("\n")
	fmt.Println("Challenge 3: Sum")
	fmt.Println("Jumlah = " + Sum.SumValue())
}
