# Go Programming: Data Structures and Algorithms

A comprehensive learning repository for Go programming with a focus on data structures and algorithms implementation.

## 📚 Project Structure

```
go-programming/
├── data-structures/           # Core data structure implementations
│   ├── arrays/               # Dynamic arrays
│   ├── linked-lists/         # Singly linked lists
│   ├── stacks/              # LIFO stack implementation
│   ├── queues/              # FIFO queue (regular and circular)
│   ├── trees/               # Binary trees and BST
│   ├── graphs/              # Graph representations and algorithms
│   ├── heaps/               # Min/Max heap implementations
│   └── hash-tables/         # Hash table with collision handling
├── algorithms/               # Algorithm implementations
│   ├── sorting/             # Various sorting algorithms
│   ├── searching/           # Search algorithms
│   ├── dynamic-programming/ # DP problems and solutions
│   ├── greedy/              # Greedy algorithm examples
│   ├── backtracking/        # Backtracking problems
│   └── graph-algorithms/    # Graph traversal and pathfinding
├── examples/                 # Practical examples and demos
│   ├── beginner/            # Basic usage examples
│   ├── intermediate/        # Common interview problems
│   └── advanced/            # Complex algorithm implementations
└── utils/                   # Helper utilities
    ├── input-output/        # I/O helpers and benchmarking
    ├── testing/             # Testing utilities
    └── benchmarks/          # Performance benchmarking tools
```

## 🚀 Getting Started

### Prerequisites

- Go 1.24.2 or later
- Basic understanding of Go syntax

### Installation

1. Clone this repository:
```bash
git clone <your-repo-url>
cd go-programming
```

2. Initialize Go module (already done):
```bash
go mod init go-programming
```

3. Run the basic example:
```bash
go run examples/beginner/basic_usage.go
```

## 📖 Learning Path

### 1. **Data Structures** (Start Here)
Learn the fundamental building blocks:

- **Arrays** (`data-structures/arrays/`)
  - Dynamic arrays with automatic resizing
  - Time complexities: Access O(1), Insert/Delete O(n)

- **Linked Lists** (`data-structures/linked-lists/`)
  - Singly linked list implementation
  - Operations: Insert, Delete, Search, Reverse

- **Stacks** (`data-structures/stacks/`)
  - LIFO (Last In, First Out) operations
  - Applications: Expression evaluation, backtracking

- **Queues** (`data-structures/queues/`)
  - FIFO (First In, First Out) operations
  - Both regular and circular queue implementations

- **Trees** (`data-structures/trees/`)
  - Binary Search Trees (BST)
  - Traversals: Inorder, Preorder, Postorder

- **Hash Tables** (`data-structures/hash-tables/`)
  - Collision handling with chaining
  - O(1) average case for insert/search/delete

### 2. **Sorting Algorithms** (`algorithms/sorting/`)
Master different sorting techniques:

- **Simple Sorts**: Bubble, Selection, Insertion Sort
- **Efficient Sorts**: Merge Sort, Quick Sort, Heap Sort  
- **Special Cases**: Counting Sort for integers

**Time Complexities:**
- Bubble/Selection/Insertion: O(n²)
- Merge/Heap: O(n log n)
- Quick Sort: O(n log n) average, O(n²) worst
- Counting Sort: O(n + k)

### 3. **Searching Algorithms** (`algorithms/searching/`)
Learn efficient search techniques:

- **Linear Search**: O(n) - works on unsorted arrays
- **Binary Search**: O(log n) - requires sorted array
- **Advanced**: Interpolation, Exponential, Jump, Ternary Search

### 4. **Problem Solving** (`examples/`)
Apply your knowledge to real problems:

- **Beginner**: Basic usage of all data structures
- **Intermediate**: Common interview questions
- **Advanced**: Complex algorithmic challenges

## 🏃‍♂️ Quick Examples

### Using a Dynamic Array
```go
package main

import (
    "fmt"
    "go-programming/data-structures/arrays"
)

func main() {
    da := arrays.NewDynamicArray(5)
    da.Append(10)
    da.Append(20)
    da.Insert(1, 15)
    fmt.Println(da) // [10, 15, 20]
}
```

### Sorting an Array
```go
package main

import (
    "fmt"
    "go-programming/algorithms/sorting"
)

func main() {
    arr := []int{64, 34, 25, 12, 22}
    sorted := sorting.MergeSort(arr)
    fmt.Println(sorted) // [12, 22, 25, 34, 64]
}
```

### Binary Search
```go
package main

import (
    "fmt"
    "go-programming/algorithms/searching"
)

func main() {
    arr := []int{1, 3, 5, 7, 9, 11}
    index := searching.BinarySearch(arr, 7)
    fmt.Println(index) // 3
}
```

## 🧪 Testing

Each data structure and algorithm includes comprehensive tests:

```bash
# Test specific package
go test ./data-structures/arrays/
go test ./algorithms/sorting/

# Test all packages
go test ./...

# Run with verbose output
go test -v ./data-structures/arrays/

# Run benchmarks
go test -bench=. ./data-structures/arrays/

# Test with coverage
go test ./data-structures/stacks -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## 📊 Performance Analysis

Use the built-in benchmarking utilities:

```go
package main

import (
    "go-programming/utils/input-output"
    "go-programming/algorithms/sorting"
)

func main() {
    arr := io.GenerateRandomArray(10000, 1, 10000)
    
    io.Benchmark("Quick Sort", func() {
        sorting.QuickSort(arr)
    })
    
    io.Benchmark("Merge Sort", func() {
        sorting.MergeSort(arr)
    })
}
```

## 🎯 Learning Tips

1. **Start with Basics**: Begin with arrays and linked lists
2. **Understand Complexity**: Always consider time and space complexity
3. **Practice Implementation**: Write code from scratch, don't just read
4. **Test Your Code**: Use the provided tests and write your own
5. **Benchmark Performance**: Compare different algorithms on various inputs
6. **Solve Problems**: Work through the examples to apply your knowledge

## 🔄 Common Patterns

### Iterator Pattern
```go
// Iterate through linked list
current := ll.Head
for current != nil {
    fmt.Println(current.Data)
    current = current.Next
}
```

### Two Pointers
```go
// Used in many algorithms
left, right := 0, len(arr)-1
for left < right {
    // Process elements
    left++
    right--
}
```

### Divide and Conquer
```go
// Pattern used in merge sort, binary search
func divideAndConquer(arr []int) {
    if len(arr) <= 1 {
        return // Base case
    }
    mid := len(arr) / 2
    left := arr[:mid]
    right := arr[mid:]
    // Recursively process both halves
}
```

## 🚀 Advanced Topics (Coming Soon)

- **Graphs**: Adjacency lists, DFS, BFS, Dijkstra's algorithm
- **Dynamic Programming**: Fibonacci, knapsack, longest common subsequence
- **Advanced Trees**: AVL trees, Red-Black trees, Segment trees
- **String Algorithms**: Pattern matching, trie data structure

## 📝 Contributing

Feel free to contribute by:
- Adding new algorithms
- Improving existing implementations
- Adding more test cases
- Creating additional examples
- Fixing bugs or improving documentation

## 📄 License

This project is for educational purposes. Feel free to use and modify as needed for your learning journey.

---

Happy coding! 🎉 Start with `examples/beginner/basic_usage.go` and work your way through the data structures and algorithms.
