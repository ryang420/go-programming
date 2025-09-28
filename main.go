package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("🎉 Welcome to Go Programming: Data Structures and Algorithms!")
	fmt.Println("===========================================================")

	fmt.Println("\n📚 Available Examples:")
	fmt.Println("1. Basic Usage Examples:")
	fmt.Println("   go run examples/beginner/basic_usage.go")

	fmt.Println("\n2. Intermediate Problems:")
	fmt.Println("   go run examples/intermediate/problems.go")

	fmt.Println("\n🧪 Run Tests:")
	fmt.Println("   go test ./...                    # Test all packages")
	fmt.Println("   go test ./data-structures/...    # Test all data structures")
	fmt.Println("   go test ./algorithms/...         # Test all algorithms")

	fmt.Println("\n📊 Run Benchmarks:")
	fmt.Println("   go test -bench=. ./data-structures/arrays/")
	fmt.Println("   go test -bench=. ./algorithms/sorting/")

	fmt.Println("\n📖 Learning Path:")
	fmt.Println("1. Start with data structures (arrays, linked lists, stacks, queues)")
	fmt.Println("2. Learn sorting algorithms (bubble, merge, quick sort)")
	fmt.Println("3. Practice searching algorithms (linear, binary search)")
	fmt.Println("4. Solve problems in examples/intermediate/")

	fmt.Println("\n📁 Project Structure:")
	printProjectStructure()

	// Check if any arguments were provided to run specific examples
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "basic":
			fmt.Println("\n🚀 Running basic examples...")
			// This would require importing and running the basic examples
			fmt.Println("Run: go run examples/beginner/basic_usage.go")
		case "problems":
			fmt.Println("\n🚀 Running intermediate problems...")
			fmt.Println("Run: go run examples/intermediate/problems.go")
		case "test":
			fmt.Println("\n🧪 To run tests, use: go test ./...")
		default:
			fmt.Printf("\nUnknown command: %s\n", os.Args[1])
			fmt.Println("Available commands: basic, problems, test")
		}
	}

	fmt.Println("\n✨ Happy Learning! Start with: go run examples/beginner/basic_usage.go")
}

func printProjectStructure() {
	structure := `
data-structures/
├── arrays/          # Dynamic arrays with auto-resize
├── linked-lists/    # Singly linked list implementation  
├── stacks/          # LIFO stack operations
├── queues/          # FIFO queue (regular & circular)
├── trees/           # Binary search trees
└── hash-tables/     # Hash table with chaining

algorithms/
├── sorting/         # Bubble, merge, quick, heap sort
├── searching/       # Linear, binary, interpolation search
├── dynamic-programming/  # (Coming soon)
├── greedy/          # (Coming soon)
├── backtracking/    # (Coming soon)
└── graph-algorithms/# (Coming soon)

examples/
├── beginner/        # Basic usage demonstrations
├── intermediate/    # Common interview problems
└── advanced/        # (Coming soon)

utils/
├── input-output/    # I/O helpers and benchmarking
├── testing/         # Testing utilities (Coming soon)
└── benchmarks/      # Performance tools (Coming soon)`

	fmt.Println(structure)
}
