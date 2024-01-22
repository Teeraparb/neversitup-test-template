package main

import _ "fmt"

func CountSmilyFace(text []string) int {
	if len(text) == 0 {
		return 0
	}

	counter := 0

	for _, value := range text {
		if len(value) == 2 {
			if (value[0] == ':' || value[0] == ';') && (value[1] == ')' || value[1] == 'D') {
				counter += 1
			}
		} else {
			if (value[0] == ':' || value[0] == ';') && (value[1] == '-' || value[1] == '~') && (value[2] == ')' || value[2] == 'D') {
				counter += 1
			}
		}
	}

	return counter
}
