package easy

import "fmt"

/*
Hackerrank: Grading problem
Link: https://www.hackerrank.com/challenges/grading/problem
*/
func Grading() {
	out := gradingStudents([]int32 {73, 67, 38, 33})
	fmt.Println(out)
}

func gradingStudents(grades []int32) []int32 {

	roundedGrades := make([]int32, 0, len(grades))

	for _,v := range grades {
		if v < 38 {
			roundedGrades = append(roundedGrades, v)
			continue
		}
		quotient := v/5
		remainder := v%5
		if remainder == 0 {
			roundedGrades = append(roundedGrades, v)
			continue
		}
		quotient++
		newVal := quotient * 5
		remainder = newVal - v
		if remainder < 3 {
			v = v + remainder
		}
		roundedGrades = append(roundedGrades, v)
	}
	return  roundedGrades
}
