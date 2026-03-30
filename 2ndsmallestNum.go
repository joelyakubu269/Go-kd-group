package main

import "fmt"

func main() {
	sli := []int{3, 7, 2, 9, 4}
	fmt.Println(secSmallestNum(sli))
}
func secSmallestNum(sli []int) int {
	minNum := sli[0]
	secMinNum := sli[0]
	for _, r := range sli {
		if r < minNum {
			secMinNum = minNum
			minNum = r
		} else if r < secMinNum && r != minNum {
			secMinNum = r
		}
	}
	return secMinNum
}
