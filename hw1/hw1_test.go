package main

import "testing"

func TestNegativeNums(t *testing.T) {
	got := evaluation([]int{-5,-4,-2,-1})
	if got != 1 {
		t.Errorf("Negative number test result: %d; want 1", got)
	}
}

func TestZeroNums(t *testing.T) {
	got := evaluation([]int{0,0,0,0})
	if got != 0 {
		t.Errorf("Zero number test result: %d; want 0", got)
	}
}

func TestBigArray(t *testing.T) {
	got := evaluation([]int{1,5,2,4,1,3,2,4,1,3,2,4,5,6,7,7,2,3,5,2,3,4,2})
	if got != 7 {
		t.Errorf("Big array test result: %d; want 7", got)
	}
}