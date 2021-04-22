package main

import (
	"bufio"
	"fmt"
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

func evaluation(slice []int) int{
	var result []int
	for k := 1; k < len(slice)-1; k++ {
		if slice[k] < ((slice[k-1]+slice[k+1]) / 2) {
			result = append(result, slice[k])
		}
	}
	return len(result)
}

func main() {
	fmt.Print(evaluation(cmdReader()))
}