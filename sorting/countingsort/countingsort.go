package main

import "fmt"

func countingSort(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}
	var min,max = slice[0], slice[0]
	for i:=0; i<len(slice); i++ {
		if slice[i] > max {
			max = slice[i]
		} else if slice[i] < min {
			min = slice[i]
		}
	}
	countingSlice := make([]int, max-min+1)
	for i:=0; i<len(slice); i++ {
		countingSlice[slice[i]-min]++
	}
	k := 0
	for i:= 0; i < len(countingSlice); i++ {
		current := i + min
		for j := 0; j < countingSlice[i]; j++{
			slice[k] = current
			k++
		}
	}
	fmt.Print(countingSlice)
	return slice
}

func main() {
	//var test = []int {-5,-1,5,6,2,4,3,1,2,3,2,2,5}
	var test = []int{2, 253, 252, 7, 7, 155, 101, 82, 253, 0, 252, -7, 252, 7, 63, -95, 15, 154, 201, -29}
	fmt.Print(countingSort(test))
	//fmt.Print(countingSort([]int{}))
}