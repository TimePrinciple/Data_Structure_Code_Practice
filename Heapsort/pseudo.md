# Heapsort

## Wikipedia implementation

> Max heap

```
left_child_of(node_index) -> left_child_index {
    return node_index * 2 + 1
}

parent_of(node_index) -> parent_index {
    return floor((node_index - 1) / 2)
}

heapsort(array, length) {
    //Build the heap in array so that largest value is at the root
    heapify(array, length)

    // The following loop maintains the invariants that array[0: end - 1]
    // is a heap, and every element array[end: length - 1] beyond end is
    // greater than everything before it
    // i.e. array[end: length - 1] is in sorted order
    end := length

    while end > 1 {
        // Reduce the heap size
        --end
        // array[0] is the root and largest value. Moves it in front of
        // the sorted elements
        swap(array[0], array[end])
        sift_down(array, 0, end)
    }
}

// Put elements of array in heap order, in-place operation
heapify(array, length) {
    // start is initialized to the first leaf node
    // The last element in a 0-based array is at index length - 1; find
    // the parent of that element
    start := parent_of(length + 1)

    while start > 0 {
        // Go to the last non-heap node
        --start
        // sift down the node at index `start` to the proper place such
        // that all nodes below the start index are in heap order
        sift_down(array, start, length)
    }
    // After sifting down the root all nodes/elements are in heap order
}

// Repair the heap whose root element is at index `start`, assuming the
// heaps rooted at its children are valid
sift_down(array, root, end) {
    // While the root has at least one child
    while left_child_of(root) < end {
        // Get the left child of root
        child := left_child_of(root)
        // If there is a right child and that child is greater
        if child + 1 < end && array[child] < array[child + 1] {
            // Choose the greater child
            child = child + 1
        }
        if array[root] < array[child] {
            swap(array[root], array[child])
            // Repeat to continue sifting down the child
            root = child
        } else {
            // The root holds the largest element, done
            return
        }
    }
}
```