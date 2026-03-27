package main

import "fmt"

func main() {
	nums := []int{3, 7, 2, 9, 4}
	fmt.Println(findMaxNum(nums))

	//even, odd := (countEvenorOdd(nums))
	//fmt.Printf("num of even is %d, num of odd is %d:\n", even, odd)
}

// func countEvenorOdd(sli []int) (int, int) {
// 	even := 0
// 	odd := 0
// 	for _, r := range sli {
// 		if r%2 == 0 {
// 			even++
// 		} else {
// 			odd++
// 		}
// 	}
// 	return even, odd

// }
func findMaxNum(sli []int) int {
	if len(sli) == 0 {
		return 0
	}
	maxNum := sli[0]
	for _, r := range sli {
		if r > maxNum {
			maxNum = r
		}
	}
	return maxNum
}
