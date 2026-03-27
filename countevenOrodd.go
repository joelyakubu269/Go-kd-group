package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	even, odd := (countEvenorOdd(nums))
	fmt.Printf("num of even is %d, num of odd is %d:\n", even, odd)
}

func countEvenorOdd(sli []int) (int, int) {
	even := 0
	odd := 0
	for _, r := range sli {
		if r%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even, odd

}
