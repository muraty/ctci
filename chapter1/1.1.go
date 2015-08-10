package chapter1

// Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures?

// Assumes that s string has only ASCII characters
func IsAllUnique(s string) bool {
	// exceeds the size of ASCII table
	if len(s) > 128 {
		return false
	}
	boolean_list := make([]bool, 128)
	for _, character := range s {
		if boolean_list[character] == true {
			return false
		}
		boolean_list[character] = true
	}
	return true
}
