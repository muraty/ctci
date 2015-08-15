package chapter1

import (
	"strconv"
)

// Implement a method to perform basic string compression using the counts of
// repeated characters.
// For example, the string aabcccccaaa would become a2b1c5a3.
// If the "compressed" string would not become smaller than the original string,
// your method should return the original string.

// compress given string in order to repeated characters.
// aabcccccaaa -> a2b1c5a3
func Compression(s string) string {
	result := ""
	count := 1
	prev := s[0]
	for i := 1; i < len(s); i++ {
		// Last element ?
		if i == len(s)-1 {
			if s[i] == prev {
				result += string(prev) + strconv.Itoa(count+1)
			} else {
				result += string(prev) + strconv.Itoa(1)
			}
		} else if s[i] == prev {
			count++
		} else if i != len(s)-1 {
			result += string(prev) + strconv.Itoa(count)
			count = 1
			prev = s[i]
		}
	}
	if len(result) > len(s) {
		return s
	}
	return result
}
