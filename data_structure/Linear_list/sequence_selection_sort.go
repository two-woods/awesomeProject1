package main

import "fmt"

func main() {
	waitSortSequence := []int{3, 2, 4, 1}
	selectionSort(waitSortSequence)
}

func selectionSort(sequence []int) {
	var d int
	for k, _ := range sequence[:len(sequence)-1] {
		d = k
		for j := k + 1; j < len(sequence); j++ {
			if sequence[j] <= sequence[d] {
				d = j
			}
			if d != k {
				temp := sequence[d]
				sequence[d] = sequence[k]
				sequence[k] = temp
			}
		}
	}
	fmt.Println("The ranking result is ", sequence)
}
