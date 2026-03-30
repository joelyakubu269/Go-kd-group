package main

import "fmt"

func main() {
	sli := []int{3, 7, 2, 9, 4}
	fmt.Println(smallestNum(sli))
}
func smallestNum(sli []int) int {
	minNum := sli[0]
	for _, r := range sli {
		if r < minNum {
			minNum = r
		}
	}
	return minNum
}
