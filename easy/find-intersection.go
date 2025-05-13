package main

import (
	"strings"
)

func FindIntersection(strArr []string) string {
	list1 := strings.Split(strArr[0], ", ")
	list2 := strings.Split(strArr[1], ", ")

	numMap := make(map[string]bool)
	for _, val := range list1 {
		numMap[val] = true
	}

	var result []string
	for _, val := range list2 {
		if numMap[val] {
			result = append(result, val)
		}
	}

	if len(result) == 0 {
		return "false"
	}

	return strings.Join(result, ",")
}
