package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sumOfNum(nums))
}
func sumOfNum(sli []int) int {
	conti := 0
	for i := 0; i < len(sli); i++ {
		conti += sli[i]
	}
	return conti
	// alternative code
	// for _,r:= range sli {
	// 	sum+= r
	// }
	// return sum
}
