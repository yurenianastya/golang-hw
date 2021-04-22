package main

import "fmt"

func insertionSort(slice []int)  []int{
	var k = 0
	for i := 1; i < len(slice); i++ {
		for slice[i] < slice[k] {
			swap := slice[i]
			slice[i] = slice[k]
			slice[k] = swap
			if k > 0 {
				k--
			}
			i--
		}
		k = i
	}
	return slice
}

func main() {
	fmt.Print(insertionSort([]int{82, 253, 252, 7, 7, 130, 101, 82, 253, 0, 252, -7, 252, 7, 63, -95, 15, 154, 98, -29}))
}