package chapter1

import "strings"

// Given two strings, write a method to decide if one is a permutation
// of the other.

// Checks whether the target string is a permutation of the source string.
// We assume that, strings are case sensitive and
// whitespace character matters.
func isPermutation(source string, target string) bool {
	if len(source) == len(target) {
		for _, elem := range target {
			if !strings.Contains(source, string(elem)) {
				return false
			}
		}
	} else {
		return false
	}
	return true
}
