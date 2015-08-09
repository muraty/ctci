package ctci

import (
	"fmt"
)

type ArrayList struct {
	list       []int
	size       int
	last_index int
}

// get the length of the list independently from underlying implementation
func (arr *ArrayList) Len() int {
	return arr.last_index
}

// get all the elements
func (arr *ArrayList) GetElements() []int {
	return arr.list
}

// print all the elements
func (arr *ArrayList) PrintElements() {
	for idx, elem := range arr.list {
		fmt.Println("index:", idx, "element:", elem)
	}
}

// get element of given index
func (arr *ArrayList) Get(idx int) int {
	if idx < arr.Len() {
		return arr.list[idx]
	} else {
		return -1
	}
}

// add int element to the list
func (arr *ArrayList) Set(elem int) {
	// reaches the limit of underlying array
	if int(arr.size) == arr.last_index {

		new_part := make([]int, arr.size*2)
		fmt.Println()
		copy(new_part, arr.list)
		arr.list = new_part
		arr.size = arr.size * 2

	}
	arr.list[arr.last_index] = elem
	arr.last_index++
}

// initialize the arraylist structure
func (arr *ArrayList) Init(size int) {
	// init underlying list of arraylist
	arr.size = size
	arr.list = make([]int, arr.size)
}
