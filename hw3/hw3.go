package main

import (
	"fmt"
	"os"
	"strconv"
)

func cmdReader() int {
	arg := os.Args[1]
	argument, _ := strconv.Atoi(arg)
	return argument
}

//func scanner() []int {
//	var slice []int
//	scanner := bufio.NewScanner(os.Stdin)
//	fmt.Print("Type numbers:")
//	for scanner.Scan() {
//		value := scanner.Text()
//		i, err := strconv.Atoi(value)
//		if len(value) == 0 {
//			break
//		}
//		if err != nil {
//			fmt.Println("Enter a valid number or empty line to break")
//		} else {
//			slice = append(slice, i)
//		}
//	}
//	return slice
//}

func printPascal(n int) int{
	if n < 0 {
		return 0
	}
	var array = make([][]int, n)
	for i:= range array {
		array[i] = make([]int, n)
	}
	for line := 0; line < n; line++ {
		for i := 0; i<= line; i++ {
			if line == i || i == 0 {
				array[line][i] = 1
			} else {
				array[line][i] = array[line-1][i-1] + array[line-1][i]
			}
			fmt.Print(array[line][i]," ")
		}
	}
	return 0
}

func main() {
	fmt.Print(printPascal(cmdReader()))
}
