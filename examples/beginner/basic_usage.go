package main

import (
	"fmt"
	"go-programming/algorithms/searching"
	"go-programming/algorithms/sorting"
	"go-programming/data-structures/arrays"
	linkedlists "go-programming/data-structures/linked-lists"
	"go-programming/data-structures/queues"
	"go-programming/data-structures/stacks"
	io "go-programming/utils/input-output"
)

func main() {
	fmt.Println("=== Go Data Structures and Algorithms - Basic Usage ===")

	// Demonstrate Dynamic Array
	fmt.Println("\n1. Dynamic Array Example:")
	demonstrateDynamicArray()

	// Demonstrate Linked List
	fmt.Println("\n2. Linked List Example:")
	demonstrateLinkedList()

	// Demonstrate Stack
	fmt.Println("\n3. Stack Example:")
	demonstrateStack()

	// Demonstrate Queue
	fmt.Println("\n4. Queue Example:")
	demonstrateQueue()

	// Demonstrate Sorting Algorithms
	fmt.Println("\n5. Sorting Algorithms Example:")
	demonstrateSorting()

	// Demonstrate Searching Algorithms
	fmt.Println("\n6. Searching Algorithms Example:")
	demonstrateSearching()
}

func demonstrateDynamicArray() {
	da := arrays.NewDynamicArray[int](5)

	// Add elements
	for i := 1; i <= 5; i++ {
		da.Append(i * 10)
	}

	fmt.Printf("After adding elements: %v\n", da)

	// Insert element
	da.Insert(2, 25)
	fmt.Printf("After inserting 25 at index 2: %v\n", da)

	// Get element
	if val, err := da.Get(2); err == nil {
		fmt.Printf("Element at index 2: %d\n", val)
	}

	// Delete element
	da.Delete(2)
	fmt.Printf("After deleting element at index 2: %v\n", da)
}

func demonstrateLinkedList() {
	ll := linkedlists.NewLinkedList()

	// Add elements
	for i := 1; i <= 5; i++ {
		ll.Append(i * 5)
	}

	fmt.Printf("After adding elements: %s\n", ll)

	// Prepend element
	ll.Prepend(0)
	fmt.Printf("After prepending 0: %s\n", ll)

	// Find element
	index := ll.Find(15)
	fmt.Printf("Index of element 15: %d\n", index)

	// Reverse list
	ll.Reverse()
	fmt.Printf("After reversing: %s\n", ll)
}

func demonstrateStack() {
	stack := stacks.NewStack()

	// Push elements
	for i := 1; i <= 5; i++ {
		stack.Push(i)
		fmt.Printf("Pushed %d: %s\n", i, stack)
	}

	// Pop elements
	for !stack.IsEmpty() {
		if val, err := stack.Pop(); err == nil {
			fmt.Printf("Popped %d: %s\n", val, stack)
		}
	}
}

func demonstrateQueue() {
	queue := queues.NewQueue()

	// Enqueue elements
	for i := 1; i <= 5; i++ {
		queue.Enqueue(i * 2)
		fmt.Printf("Enqueued %d: %s\n", i*2, queue)
	}

	// Dequeue elements
	for !queue.IsEmpty() {
		if val, err := queue.Dequeue(); err == nil {
			fmt.Printf("Dequeued %d: %s\n", val, queue)
		}
	}
}

func demonstrateSorting() {
	// Generate random array
	arr := io.GenerateRandomArray(10, 1, 100)
	io.PrintArray(arr, "Original array")

	// Demonstrate different sorting algorithms
	algorithms := map[string]func([]int) []int{
		"Bubble Sort":    sorting.BubbleSort,
		"Selection Sort": sorting.SelectionSort,
		"Insertion Sort": sorting.InsertionSort,
		"Merge Sort":     sorting.MergeSort,
		"Quick Sort":     sorting.QuickSort,
		"Heap Sort":      sorting.HeapSort,
	}

	for name, sortFunc := range algorithms {
		sorted := sortFunc(arr)
		io.PrintArray(sorted, fmt.Sprintf("%s result", name))
	}
}

func demonstrateSearching() {
	// Create sorted array for search demonstrations
	arr := io.GenerateSortedArray(20, 1, 100, true)
	io.PrintArray(arr, "Sorted array for searching")

	target := arr[len(arr)/2] // Pick middle element as target
	fmt.Printf("Searching for target: %d\n", target)

	// Demonstrate different searching algorithms
	algorithms := map[string]func([]int, int) int{
		"Linear Search":           searching.LinearSearch,
		"Binary Search":           searching.BinarySearch,
		"Binary Search Recursive": searching.BinarySearchRecursive,
		"Interpolation Search":    searching.InterpolationSearch,
		"Exponential Search":      searching.ExponentialSearch,
		"Jump Search":             searching.JumpSearch,
		"Ternary Search":          searching.TernarySearch,
	}

	for name, searchFunc := range algorithms {
		index := searchFunc(arr, target)
		if index != -1 {
			fmt.Printf("%s: Found at index %d\n", name, index)
		} else {
			fmt.Printf("%s: Not found\n", name)
		}
	}
}
