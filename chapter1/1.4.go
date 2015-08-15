package chapter1

// Write a method to replace all spaces in a string with '%20'.
// You may assume that the string has sufficient space at the end of the string
// to hold the additional characters, and that you are given the "true"
// length of the string. (Note: if implementing in Java, please use a
// character array so that you can perform this operation in place.)
// EXAMPLE
// Input: "Mr John Smith      "
// Output: "Mr%20Dohn%20Smith"

// Replace whitespace characters with %20 string
func ReplaceSpaces(s *string) {
	exchange_string := "%20"
	temp := ""
	trimString(s)
	for _, char := range *s {
		if string(char) == " " {
			temp += exchange_string
		} else {
			temp += string(char)
		}
	}
	*s = temp
}

// trim space characters at the end of the string
func trimString(s *string) {
	temp := ""
	st := *s
	for i := len(st) - 1; i >= 0; i-- {
		if string(st[i]) == " " {
			continue
		} else {
			temp = st[0 : i+1]
			break
		}
	}
	*s = temp
}
