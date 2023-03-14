package main

import (
	arrayvector "array/structure"
	"fmt"
)

func main() {
	// Create a new array list
	myArrayList := &arrayvector.Array{}
	myVector := &arrayvector.Vector{}
	myVector.Add(10)
	// Add elements to the array list
	myArrayList.Add(10)
	myArrayList.Add(20)
	myArrayList.Add(30)

	// Get the size of the array list
	size := myArrayList.Size()
	fmt.Println("Size:", size)

	// Get the element at index 1
	element := myArrayList.Get(1)
	fmt.Println("Element at index 1:", element)

	// Remove the element at index 1
	myArrayList.Remove(1)

	// Get the size of the array list
	size = myArrayList.Size()
	fmt.Println("Size:", size)
}
