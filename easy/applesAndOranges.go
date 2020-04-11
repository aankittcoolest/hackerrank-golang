package easy

import "fmt"


/*
Hackerrank: Apples and Oranges Problem
Link: https://www.hackerrank.com/challenges/apple-and-orange/problem
*/
func AppleAndOrange() {
	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}

// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var appleCount, orangeCount int32
	for _, v := range apples {
		rangeApple := v + a
		if s <= rangeApple && t >= rangeApple {
			appleCount++
		}
	}

	for _, v := range oranges {
		rangeOrange := v + b
		if s <= rangeOrange && t >= rangeOrange {
			orangeCount++
		}
	}

	fmt.Println(appleCount)
	fmt.Println(orangeCount)
}
