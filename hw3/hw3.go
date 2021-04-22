package main

import "fmt"

func printPascal()  {
	const n = 6
	var array [n][n] int
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
}

func main() {
	printPascal()
}
