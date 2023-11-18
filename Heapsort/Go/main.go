//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func leftChildOf(i int) int {
	return 2*i + 1
}

func rightChildOf(i int) int {
	return 2*i + 2
}

func parentOf(i int) int {
	return (i - 1) / 2
}

// Repair the heap whose root element is at index 'start', assuming the
// heaps rooted at its children are valid
func siftDown(array []int, root int, end int) {
	// While the root has at least one child
	for leftChildOf(root) < end {
		// Get the left child of root
		child := leftChildOf(root)
		// If there is a right child and that child is greater
		if child+1 < end && array[child] < array[child+1] {
			child = child + 1
		}
		if array[root] < array[child] {
			array[root], array[child] = array[child], array[root]
			// Repeat to continue sifting down the child
			root = child
		} else {
			// The root holds the largest element, done
			return
		}
	}
}

// Put elements of array in heap order, in-place
func heapify(array []int, length int) {
	// start is initialized to the first leaf node
	// The last element in a 0-based array is at index length - 1; find
	// the parent of that element
	start := parentOf(length + 1)

	for start > 0 {
		// Go to the last non-heap node
		start--
		// sift down the node at index 'start' to the proper place such
		// that all nodes below the start index are in heap order
		siftDown(array, start, length)
	}
	// After sifting down the root all nodes/elements are in heap order
}

func heapSort(array []int, length int) {
	// Build the heap in array so that largest value is at the root
	heapify(array, length)

	// The following loop maintains the invariants that array[0: end - 1]
	// is a heap, and every element array[end: length - 1] beyond end is
	// greater than everything before it
	// i.e. array[end: length - 1] is in sorted order
	end := length
	for end > 1 {
		// The heap size is reduced by one
		end--
		// array[0] is the root and largest value. Moves it in front of
		// the sorted elements.
		array[0], array[end] = array[end], array[0]
		// The heap property damaged, restore it
		siftDown(array, 0, end)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Open the input file
	file, err := os.Open("../input")
	check(err)

	// Wrap the file with Scanner
	scanner := bufio.NewScanner(file)
	// Split the stream by white spaces
	scanner.Split(bufio.ScanWords)

	// Construct a buffer to hold the input
	var array []int
	for scanner.Scan() {
		// Retrieve the content
		content := scanner.Text()
		num, err := strconv.Atoi(content)
		check(err)
		// Put in the back of the buffer
		array = append(array, num)
	}

	heapSort(array, len(array))

	fmt.Println(array)
}
