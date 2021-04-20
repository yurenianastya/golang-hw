package main

import "fmt"

func insertionSort(slice []int)  []int{
	var k = 0
	for i := 1; i < len(slice); i++ {
		for slice[i] < slice[k] && k > 0 {
				swap := slice[i]
				slice[i] = slice[k]
				slice[k] = swap
				k--
				i--
		}
		k = i
	}
	return slice
}

func main() {
	fmt.Print(insertionSort([]int{1,5,23,2,56,7,8,4}))
}