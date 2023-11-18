# Quicksort

## Lomuto implementation

> Choose the last element as the pivot

```
quick_sort(&array, left, right) {
    // Ensure indeces are in correct order (condition for recursion to stop)
    // Return immediately otherwise
    if left >= right || left <0 {
        return
    }

    // Partition array and get the pivot index
    p := partition(array, left, right);
    // After partition, the array is divided by the pivot, the elements on the left of the index are smaller than the pivot, greater the otherwise

    // Sort the two partitions recursively
    quicksort(array, left, p - 1)
    quicksort(array, p + 1, right)
}

partition(&array, left, right) -> pivot_index {
    // Choose the last element as the pivot
    pivot := array[right]

    // Set temporary pivot index, right before the left
    pivot_index := left - 1

    // Iterate from left to right
    for i := left; i < right; ++i {
        // If the current element is less than or equal to the pivot
        if array[i] <= pivot {
            // Forward the pivot index before swap
            ++pivot_index
            // Swap the current element with the element at the temporary pivot index
            swap(array[i], array[pivot_index])
        }
    }

    // Forward the pivot index to the correct pivot position, betwee the smaller and larger elements
    ++pivot_index
    // Swap the element with the last element, which was chosen as the pivot
    swap(array[pivot_index], array[right])
    // Return the index of the pivot
    return pivot_index
}
```