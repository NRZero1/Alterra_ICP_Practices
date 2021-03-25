package Grade

import "fmt"

func checkGrade() string {
	var nilai int
	var finalGrade string

	fmt.Print("Cek nilai yang akan diinput: ")
	fmt.Scanln(nilai)

	if nilai > 85 {
		finalGrade = "Grade A"
	} else if nilai > 75 {
		finalGrade = "Grade B"
	} else if nilai > 65 {
		finalGrade = "Grade C"
	} else if nilai > 55 {
		finalGrade = "Grade D"
	} else {
		finalGrade = "Tidak Lulus"
	}

	return finalGrade
}
