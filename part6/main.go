package main

import (
	"fmt"
)

func main() {
	strings := []string{
		"short",
		"a much longer string",
		"tiny",
		"this is the longest string in the slice",
		"medium length",
	}

	longestString := findLongestString(strings)

	fmt.Println("The longest string is:", longestString)
}

func findLongestString(strings []string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, str := range strings {
		if len(str) > len(longest) {
			longest = str
		}
	}

	return longest
}
