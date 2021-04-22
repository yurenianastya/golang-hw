package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func cmdReader() []int {
	args := os.Args[1:]
	separatedArgs := strings.Join(args, ",")
	stringSlice := strings.Split(separatedArgs, ",")
	intSlice := make([]int, len(stringSlice))
	for i, s := range stringSlice {
		intSlice[i], _ = strconv.Atoi(s)
	}
	return intSlice
}

func scanner() []int {
	var slice []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type numbers:")
	for scanner.Scan() {
		value := scanner.Text()
		i, err := strconv.Atoi(value)
		if len(value) == 0 {
			break
		}
		if err != nil {
			fmt.Println("Enter a valid number or empty line to break")
		} else {
			slice = append(slice, i)
		}
	}
	return slice
}

func factorial (number int)(result int) {
	result = 1
	var i int
	for i = 1; i <= number; i++ {
		result *= i
	}
	return result
}

func evaluation(slice []int) int{
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
	fmt.Print(evaluation(cmdReader()))
}