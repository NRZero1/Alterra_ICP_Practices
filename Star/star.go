package Star

import (
	"fmt"
	"strings"
)

func CreateStar() {
	var loop int

	fmt.Print("Input loop value: ")
	fmt.Scanln(&loop)

	for i := 1; i <= loop; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
