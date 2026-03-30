package main

import "fmt"

func main() {
	sli := []int{1, 3, 9, 5, 8, 7}
	fmt.Println(sort(sli))
}

//	func sort(sli []int) []int {
//		for i := 0; i < len(sli)-1; i++ {
//			for j := 0; j < len(sli)-i-1; j++ { // j < len(sli)-i-1 the -i is to prevent checking sorted numbers
//				if sli[j] > sli[j+1] {  // compare to adjacent pairs, then swap if condition is met
//					sli[j], sli[j+1] = sli[j+1], sli[j]
//				}
//			}
//		}
//		return sli
//	}
func sort(sli []int) []int { // the bubble sort method
	for i := 0; i < len(sli); i++ { // in this code we keep i and compare with [i:]
		min := i                            // use that find least number ideology but in this case we store position
		for j := i + 1; j < len(sli); j++ { // thats why j= i +1
			if sli[j] < sli[min] { // we use j to compare
				min = j // then we store the position
			}
		}
		sli[i], sli[min] = sli[min], sli[i] // then we swap their positions
	}
	return sli
}
