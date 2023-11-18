//lint:file-ignore U1000 Ignore all unused code
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func partition(array []int, left int, right int) int {
	// Choose the last element as the pivot
	pivot := array[right]

	// Set temporary pivot index
	i := left - 1

	for j := left; j < right; j++ {
		// If the current element is less than or equal to the pivot
		if array[j] <= pivot {
			// Move the temporary pivot forward
			i++
			// Swap the current with the element at the temporary pivot index
			array[i], array[j] = array[j], array[i]
		}
	}

	// Move the pivot element to the correct pivot position, between the smaller
	// elements and larger elements
	i++
	array[i], array[right] = array[right], array[i]

	// Return the pivot index
	return i
}

// Sorts a (partition of an) array, divides it into partitions, then sort those
func quickSort(array []int, left int, right int) {
	// Ensure indeces are in correct order
	// Return immediately otherwise
	if left >= right || left < 0 {
		return
	}

	// Partition array and get the pivot index
	p := partition(array, left, right)

	quickSort(array, left, p-1)
	quickSort(array, p+1, right)
}

func main() {
	file, err := os.Open("../input")
	check(err)

	reader := bufio.NewScanner(file)
	reader.Split(bufio.ScanWords)

	var array []int
	for reader.Scan() {
		num, err := strconv.Atoi(reader.Text())
		check(err)
		array = append(array, num)
	}

	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
