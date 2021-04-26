package main

import "fmt"

func isRepetitive(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[0] {
			return false
		}
	}
	return true
}

func msBits(slice []int, sliceElement int) int {
	max,min := slice[0], slice[0]
	for i:=0;i<len(slice);i++ {
		if slice[i] > max {
			max = slice[i]
		} else if slice[i] < min {
			min = slice[i]
		}
	}
	number := int(float64(sliceElement-min)*float64(len(slice))/float64(max-min+1))
	return number
}

func bucketSort(slice []int, msBits func(slice []int, sliceElement int) int) []int{
	buckets := make([][]int, len(slice))

	for i:=0; i<len(slice);i++ {
		buckets[i] = make([]int, len(slice))
		for j:=0;j<len(slice);j++ {
			buckets[i] = nil
		}
	}

	for i:=0;i<len(slice);i++ {
		bucketNum := msBits(slice, slice[i])
		buckets[bucketNum] = append(buckets[bucketNum], slice[i])
	}

	for i:=0;i<len(buckets);i++ {
		if len(buckets[i]) > 1 && isRepetitive(buckets[i]) == false{
			buckets[i] = bucketSort(buckets[i], msBits)
		}
	}
	k:=0
	for i:=0;i<len(buckets);i++ {
		for j:=0;j<len(buckets[i]);j++{
			slice[k] = buckets[i][j]
			k++
		}
	}
	return slice
}

func main() {
	test := []int{1,5,6,7,3,-5,-1,4,-5}
	fmt.Print(bucketSort(test, msBits))
}


