package main

import "testing"

func TestFindOddNumber(t *testing.T) {
	resultCase1 := FindOddNumber([]int{7})

	if resultCase1 != 7 {
		t.Errorf("Incorrect input %d != output %d", 7, resultCase1)
	}

	resultCase2 := FindOddNumber([]int{0})

	if resultCase2 != 0 {
		t.Errorf("Incorrect input %d != output %d", 0, resultCase2)
	}

	resultCase3 := FindOddNumber([]int{1, 1, 2})

	if resultCase3 != 2 {
		t.Errorf("Incorrect input %d != output %d", 2, resultCase3)
	}

	resultCase4 := FindOddNumber([]int{0, 1, 0, 1, 0})

	if resultCase4 != 0 {
		t.Errorf("Incorrect input %d != output %d", 0, resultCase4)
	}

	resultCase5 := FindOddNumber([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})

	if resultCase5 != 4 {
		t.Errorf("Incorrect input %d != output %d", 4, resultCase5)
	}
}
