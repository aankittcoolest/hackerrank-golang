package easy

import "fmt"

/*
Hackerrank Kangaroo jump problem
Link: https://www.hackerrank.com/challenges/kangaroo/problem
*/

func Kangaroo() {
	x := [5] int32 {0,3,4,2}
	output := kangaroo(x[0],x[1],x[2],x[3])
	fmt.Print(output)
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if x2>x1 && v2>v1 {
		return "NO"
	}

	var kang1 int32 = x1 + v1
	var kang2 int32 =  x2 + v2
	for i:=0; i<=10000; i++ {
		if kang1 == kang2 {
			return "YES"
		}
		kang1 = v1 +kang1
		kang2 = v2 + kang2
	}
	return "NO"

}
