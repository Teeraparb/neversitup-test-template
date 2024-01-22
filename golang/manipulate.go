package main

import _ "fmt"

func Manipulate(text string) []string {
	n := len(text)

	if n == 0 {
		return []string{}
	}

	var answer []string
	duplicateChecker := make(map[string]bool)

	shuffling([]rune(text), duplicateChecker, &answer, 0, n-1)

	return answer
}

func shuffling(cha []rune, seen map[string]bool, answer *[]string, left int, right int) {
	if left == right {
		chaString := string(cha)
		if !seen[chaString] {
			seen[chaString] = true
			*answer = append(*answer, chaString)
		}
		return
	}

	for i := left; i <= right; i++ {
		cha[left], cha[i] = cha[i], cha[left]
		shuffling(cha, seen, answer, left+1, right)
		cha[left], cha[i] = cha[i], cha[left]
	}
}
