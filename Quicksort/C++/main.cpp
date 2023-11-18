#include <iostream>
#include <vector>
#include <algorithm>
#include <fstream>

using namespace std;

int partition(vector<int> &Array, int left, int right) {
    // Choose the last element as the pivot
    int pivot = Array[right];

    // Temporary pivot index
    int i = left - 1;

    for (int j = left; j < right; ++j) {
        // If the cuurent element is less than or equal to the pivot
        if (Array[j] <= pivot) {
            // Move the temporary pivot index forward
            ++i;
            // Swap the current element with the element at the temporary pivot index
            swap(Array[i], Array[j]);
        }
    }

    // Move the pivot element to the correct pivot position, between the smaller and
    // larger elements
    ++i;
    swap(Array[i], Array[right]);
    return i;
}

// Sorts a (partion of an) array, divides it into partitions, then sorts those
void quicksort(vector<int> &Array, int left, int right) {
    // Ensure indeces are in correct order
    // Return immediately otherwise
    if (left >= right || left < 0) {
        return;
    }

    // Partition array and get the pivot index
    int p = partition(Array, left, right);

    // Sort the two partitions
    quicksort(Array, left, p - 1);
    quicksort(Array, p + 1, right);
}


int main() {
    // Construct file stream
    ifstream file_stream;
    file_stream.open("../input");
    vector<int> Array = vector<int>();
    int element;
    while (file_stream >> element) {
        Array.emplace_back(element);
    }
    
    quicksort(Array, 0, Array.size() - 1);
    for (int e: Array) {
        cout << e << endl;
    }
}