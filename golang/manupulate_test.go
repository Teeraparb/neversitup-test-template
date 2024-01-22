package main

import (
	"reflect"
	"testing"
)

func TestManipulate(t *testing.T) {
	shuffleCase1 := Manipulate("a")
	expectedCase1 := []string{"a"}
	if !reflect.DeepEqual(shuffleCase1, expectedCase1) {
		t.Errorf("Incorrect result %v != expect %v", shuffleCase1, expectedCase1)
	}

	shuffleCase2 := Manipulate("ab")
	expectedCase2 := []string{"ab", "ba"}
	if !reflect.DeepEqual(shuffleCase2, expectedCase2) {
		t.Errorf("Incorrect result %v != expect %v", shuffleCase2, expectedCase2)
	}

	shuffleCase3 := Manipulate("abc")
	expectedCase3 := []string{"abc", "acb", "bac", "bca", "cba", "cab"}
	if !reflect.DeepEqual(shuffleCase3, expectedCase3) {
		t.Errorf("Incorrect result %v != expect %v", shuffleCase3, expectedCase3)
	}

	shuffleCase4 := Manipulate("aabb")
	expectedCase4 := []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}
	if !reflect.DeepEqual(shuffleCase4, expectedCase4) {
		t.Errorf("Incorrect result %v != expect %v", shuffleCase4, expectedCase4)
	}
}
