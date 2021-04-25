package main

import "testing"

func TestNegativeNums(t *testing.T) {
	got := printPascal(-3)
	if got != 0 {
		t.Errorf("Negative number test result: %d; want 0", got)
	}
}

func TestZeroNums(t *testing.T) {
	got := printPascal(0)
	if got != 0 {
		t.Errorf("Zero number test result: %d; want 0", got)
	}
}