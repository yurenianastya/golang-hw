package main

import "fmt"

func evaluation1(slice []int) int{
	var result []int
	for k := 1; k < len(slice)-1; k++ {
		if slice[k] < ((slice[k-1]+slice[k+1]) / 2) {
			result = append(result, slice[k])
		}
	}
	return len(result)
}

func main() {
	var sequence = []int{5,8,2,6,4,5,1,2,4,1}
	fmt.Print("the answer is: ",evaluation1(sequence))
}