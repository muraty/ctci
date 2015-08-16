package chapter1

import "strings"

// Assume you have a method isSubstring which checks if one word is a
// substring of another. Given two strings, s1 and s2, write code to check
// if s2 is a rotation of s1 using only one call to isSubstring
// (e.g.,"waterbottle"is a rotation of"erbottlewat").

// is s2 is substring of s1 ?
func isSubstring(s1 string, s2 string) bool {
	return strings.Contains(s1, s2)
}

// is s2 is rotation of s1 ?
func IsRotation(s1 string, s2 string) bool {
	// make s1 double, so it contains rotated s2
	s1s1 := s1 + s1
	return isSubstring(s1s1, s2)
}
