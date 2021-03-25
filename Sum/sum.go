package Sum

import "fmt"

func SumValue() int {
	var val int
	var jumlah int = 0

	fmt.Print("Input value: ")
	fmt.Scanln(&val)

	for i := 1; i <= val; i++ {
		jumlah += i
	}

	return jumlah
}
