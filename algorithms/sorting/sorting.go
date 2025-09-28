package sorting

// BubbleSort implements the bubble sort algorithm
// Time Complexity: O(n²), Space Complexity: O(1)
func BubbleSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		// If no swapping occurred, array is already sorted
		if !swapped {
			break
		}
	}

	return result
}

// SelectionSort implements the selection sort algorithm
// Time Complexity: O(n²), Space Complexity: O(1)
func SelectionSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if result[j] < result[minIdx] {
				minIdx = j
			}
		}
		result[i], result[minIdx] = result[minIdx], result[i]
	}

	return result
}

// InsertionSort implements the insertion sort algorithm
// Time Complexity: O(n²), Space Complexity: O(1)
func InsertionSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1

		// Move elements greater than key one position ahead
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}

	return result
}

// MergeSort implements the merge sort algorithm
// Time Complexity: O(n log n), Space Complexity: O(n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		result := make([]int, len(arr))
		copy(result, arr)
		return result
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Add remaining elements
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

// QuickSort implements the quick sort algorithm
// Time Complexity: O(n log n) average, O(n²) worst case, Space Complexity: O(log n)
func QuickSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// HeapSort implements the heap sort algorithm
// Time Complexity: O(n log n), Space Complexity: O(1)
func HeapSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	n := len(result)

	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(result, n, i)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		result[0], result[i] = result[i], result[0]
		heapify(result, i, 0)
	}

	return result
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// CountingSort implements counting sort (for non-negative integers)
// Time Complexity: O(n + k), Space Complexity: O(k) where k is the range
func CountingSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	// Find the maximum value to determine range
	max := arr[0]
	for _, val := range arr {
		if val > max {
			max = val
		}
		if val < 0 {
			// Counting sort works with non-negative integers
			panic("CountingSort works only with non-negative integers")
		}
	}

	// Create counting array
	count := make([]int, max+1)
	for _, val := range arr {
		count[val]++
	}

	// Build result array
	result := make([]int, 0, len(arr))
	for value := 0; value <= max; value++ {
		for count[value] > 0 {
			result = append(result, value)
			count[value]--
		}
	}

	return result
}
