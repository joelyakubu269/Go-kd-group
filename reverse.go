package main

import (
	"fmt"
)

func main() {
	sli := []int{3, 7, 2, 9, 4}
	fmt.Println(reverse(sli))
}
func reverse(sli []int) []int {
	for i, j := len(sli)-1, 0; i > j; i, j = i-1, j+1 { // reverses a slice of int
		sli[i], sli[j] = sli[j], sli[i]
	}
	return sli
}
