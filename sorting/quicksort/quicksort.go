package main

func getPivot(slice []int) int {
	pivot := slice[0]
	return pivot
}

func quickSort(slice []int, getPivot int) {
	if len(slice) < 2 {
		return
	}
	left, right := 0, len(slice)-1
	pivot := getPivot

	slice[pivot], slice[right] = slice[right], slice[pivot]

	for i := range slice {
		if slice[i] < slice[right] {
			slice[left], slice[i] = slice[i], slice[left]
			left++
		}
	}
	slice[left], slice[right] = slice[right], slice[left]

	quickSort(slice[:left], getPivot)
	quickSort(slice[left+1:], getPivot)
}

func main() {
	testSlice := []int{1,2,4,-5,3,4,2,13,-6,5,3,-4,2}
	quickSort(testSlice, getPivot(testSlice))
}