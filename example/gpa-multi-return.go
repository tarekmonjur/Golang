package example

import "fmt"

func myGPA(midExam float32, finalExam float32) (float32, string) {
	avgMark := (midExam + finalExam) / 2
	var gradeLetter string
	if avgMark > 90 {
		gradeLetter = "A+"
	} else if avgMark > 80 {
		gradeLetter = "A"
	} else if avgMark > 70 {
		gradeLetter = "B"
	} else if avgMark > 60 {
		gradeLetter = "C"
	} else {
		gradeLetter = "F"
	}
	return avgMark, gradeLetter
}

func main() {
	var midExamMark, finalExamMark, avgMark float32
	var gpaGrade string

	fmt.Scan(&midExamMark)
	fmt.Scanf("%f", &finalExamMark)

	avgMark, gpaGrade = myGPA(midExamMark, finalExamMark)
	fmt.Println(gpaGrade, avgMark)
}
