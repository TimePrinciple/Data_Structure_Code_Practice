# Data_Structure_Code_Practice

## Get the Code Running

For C++:
1. Enter the `C++` directory under each argorithm folder.
2. Run `gcc <file_to_compile> -lstdc++` to compile the code. 
3. Run `./a.out` to run the program.

For Go:
1. Enter the `Go` directory under each argorithm folder.
2. Run `go run .` to run the program.

For Rust:
1. Enter the `Rust` directory under each argorithm folder.
2. Run `cargo run` to run the program.

## File Processing Snippet

The snippet is served to read the content of a file named `input`, put the elements into the according container.

### C++

```C++
#include <iostream>
#include <fstream>
#include <vector>

using namespace std;

int main() {
    // Construct the input file stream
    ifstream input_file_stream;
    // Connect the stream to the file
    input_file_stream.open("../input");

    // Construct the buffer to hold the input
    vector<int> array = vector<int>();
    int element;
    while (input_file_stream >> element) {
        array.emplace_back(element);
    }

    // Process the input
    // ...

    // Return result
}
```

### Go

```Go
//lint:file-ignore U1000 Ignore all unused code
package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // Open the file
    file, err := os.Open("../input")
    check(err)

    // Construct the scanner
    scanner := bufio.NewScanner(file)
    // Separate by white spaces
    scanner.Split(bufio.ScanWords)

    // Construct the buffer to hold the input (slice)
    var array []int
    for scanner.Scan() {
        // Retrieve the content
        content := scanner.Text()
        // Convert to integer
        num, err := strconv.Atoi(content)
        check(err)
        array = append(array, num)
    }

    // Process the input
    // ...

    // Return result
    return
}
```

### Rust

```Rust
!#[allow(non_snake_case)]
!#[allow(dead_code)]



```