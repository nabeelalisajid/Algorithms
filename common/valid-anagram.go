package main

import (
	"fmt"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false;
	}
	s1LetterCount := make(map[rune]int)
	s2LetterCount := make(map[rune]int)

	for _, ch:= range s1 {
		s1LetterCount[ch]++
	}
	for _, ch :=range s2 {
		s2LetterCount[ch]++
	}

	for key, val := range s1LetterCount {
		if s2LetterCount[key] != val {
			return false
		}
	}
	return true;
}

func main() {
	s1:= "anagram"
	s2:= "nagaram"
	fmt.Println(anagram(s1, s2))
}