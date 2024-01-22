package main

import (
	_ "fmt"
)

func FindOddNumber(input []int) int {
	if len(input) == 0 {
		return 0
	}

	if len(input) == 1 {
		return input[0]
	}

	counter := make(map[int]int)

	for _, value := range input {
		counter[value] += 1
	}

	for index, value := range counter {
		if value%2 != 0 {
			return index
		}
	}

	return 0
}
