#include <iostream>
#include <fstream>
#include <algorithm>
#include <vector>

using namespace std;

const int ESTIMATE_INPUT_SIZE = 16;

int left_child_of(int node_index) {
    return 2 * node_index + 1;
}

int parent_of(int node_index) {
    return (node_index - 1) / 2;
}

// Repair the heap whose root element is at index `start`, assuming the
// heaps rooted at its children are valid
void sift_down(vector<int> &array, int root, int end) {
    // While the root has at least one child
    while (left_child_of(root) < end) {
        // Get the left child of root
        int child = left_child_of(root);
        // If there is a right child and that child is greater
        if (child + 1 < end && array[child] < array[child + 1]) {
            child = child + 1;
        }
        if (array[root] < array[child]) {
            swap(array[root], array[child]);
            // Repeat to continue sifting down the child
            root = child;
        } else {
            // The root holds the largest element, done
            return;
        }
    }
}

// Put elements of array in heap order, in-place operation
void heapify(vector<int> &array, int length) {
    // `start` is initialized to the first leaf node
    // The last element in a 0-base array is at index length - 1; find
    // the parent of that element
    int start = parent_of(length + 1);

    while (start > 0) {
        // Go to the last non-heap node
        --start;
        // sift down the node at index `start` to the proper place such
        // that all nodes below the `start` index are in heap order
        sift_down(array, start, length);
    }
    // After sifting down the root all nodes/elements are in heap order
}

void heapsort(vector<int> &array, int length) {
    // Build the heap in array so that largest value is at the root
    heapify(array, length);

    // The following loop maintains the invariants that array[0: end - 1]
    // is a heap, and every element array[end: length - 1] beyond end is
    // greater than everything before it
    // i.e. array[end: length - 1] is in sorted order
    int end = length;
    while (end > 1) {
        // Reduce the heap size
        --end;
        // array[0] si the root and largest value. Move it in front of
        // the sorted elements
        swap(array[0], array[end]);
        // The heap property damaged, restore it
        sift_down(array, 0, end);
    }
}

int main() {
    // Construct the file stream
    ifstream input_file_stream;
    // Connect the file stream to the actual file
    input_file_stream.open("../input");

    // Construct a vector of int to hold the input
    vector<int> input = vector<int>();
    // Set the capacity of this vector to avoid reallocation
    // during the pushing process
    input.reserve(ESTIMATE_INPUT_SIZE);

    int element;
    while (input_file_stream >> element) {
        input.emplace_back(element);
    }

    // Process the input
    heapsort(input, input.size());

    for (int e: input) {
        cout << e << endl;
    }

    // Print the result
    return 0;
}