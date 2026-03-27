package main

import "fmt"

func main() {
	nums := []int{3, 3, 3}

	fmt.Println(secondMax(nums))
}
func secondMax(sli []int) int {
	maxNum := sli[0]
	secondMaxn := sli[0] //The way we initialize our tracking variables matters.
	// If we pick a number like 0, it might not make sense for negative numbers.”
	if len(sli) < 2 {
		return sli[0]
	}
	for _, r := range sli {
		if r > maxNum {
			secondMaxn = maxNum
			maxNum = r

		} else if r > secondMaxn && r != maxNum { // handles duplicates
			secondMaxn = r
		}
	}
	return secondMaxn
}
