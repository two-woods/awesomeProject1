package main

import "fmt"

//不多于3n/2的平均比较次数，在数组A[1...n]中找最大最小值
func find_min_max(sequence []int) (maximum int, minimum int) {
	maximum, minimum = sequence[0], sequence[0]
	for _, i := range sequence[1:] {
		if i > maximum {
			maximum = i
		}
		if i < minimum {
			minimum = i
		}
	}
	return
}

func main() {
	sequence := []int{1, 4, 2, 3}
	fmt.Println("sequence:", sequence)
	maximum, minimum := find_min_max(sequence)
	fmt.Println("Maximum in sequence is ", maximum)
	fmt.Println("The minimum in the sequence is", minimum)
}
