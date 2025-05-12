package main

import (
	"regexp"
	"strings"
)

func LongestWord(sen string) string {

	listStr := strings.Split(sen, " ")
	longest := ""
	maxLen := 0

	re := regexp.MustCompile(`[A-Za-z]+`)

	for _, word := range listStr {
		cleaned := re.FindString(word)
		if len(cleaned) > maxLen {
			longest = cleaned
			maxLen = len(cleaned)
		}
	}

	return longest

}
