package main

import (
	"testing"
)

func TestCountSmileyFace(t *testing.T) {
	countCase1 := CountSmilyFace([]string{
		":)",
		";(",
		";}",
		":-D",
	})

	if countCase1 != 2 {
		t.Errorf("Incorrect expect %d != result %d", 2, countCase1)
	}

	countCase2 := CountSmilyFace([]string{
		";D",
		":-(",
		":-)",
		";~)",
	})

	if countCase2 != 3 {
		t.Errorf("Incorrect expect %d != result %d", 3, countCase2)
	}

	countCase3 := CountSmilyFace([]string{
		";]",
		":[",
		";*",
		":$",
		";-D",
	})

	if countCase3 != 1 {
		t.Errorf("Incorrect expect %d != result %d", 1, countCase3)
	}
}
