package chapter1

// Given an image represented by an NxN matrix, where each pixel in the
// image is 4 bytes, write a method to rotate the image by 90 degrees.
// Can you do this in place?

// rotates given matrix 90degree
func Rotate(m [4][4]int) [4][4]int {
	for layer := 0; layer < len(m[0])/2; layer++ {
		first := layer
		last := len(m[0]) - 1 - layer
		for i := layer; i < last; i++ {
			offset := i - first
			top := m[first][i]

			// top <- left
			m[first][i] = m[last-offset][first]

			// left <- bottom
			m[last-offset][first] = m[last][last-offset]

			// bottom <- right
			m[last][last-offset] = m[i][last]

			// right <- bottom
			m[i][last] = top
		}
	}
	return m
}
