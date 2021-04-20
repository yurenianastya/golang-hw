package main

import (
	"fmt"
	"math"
)

func factorial (number int)(result int) {
	result = 1
	var i int
	for i = 1; i <= number; i++ {
		result *= i
	}
	return result
}

func evaluation2(slice []int) int{
	var result []int
	for k := 0; k < len(slice); k++ {
		calculatedFactorial := factorial(k)
		if float64(slice[k]) > math.Pow(2,float64(k)) && slice[k] < calculatedFactorial {
			result = append(result, slice[k])
		}
	}
	return len(result)
}

func main() {
	var sequence = []int{0,0,0,0,0,0,0,0,0,600}
	fmt.Print("the answer is: ",evaluation2(sequence))
}