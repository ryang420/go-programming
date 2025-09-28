package searching

// LinearSearch searches for a target value in an array using linear search
// Time Complexity: O(n), Space Complexity: O(1)
func LinearSearch(arr []int, target int) int {
	for i, val := range arr {
		if val == target {
			return i
		}
	}
	return -1 // Not found
}

// BinarySearch searches for a target value in a sorted array using binary search
// Time Complexity: O(log n), Space Complexity: O(1)
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // Not found
}

// BinarySearchRecursive implements binary search recursively
// Time Complexity: O(log n), Space Complexity: O(log n)
func BinarySearchRecursive(arr []int, target int) int {
	return binarySearchHelper(arr, target, 0, len(arr)-1)
}

func binarySearchHelper(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return binarySearchHelper(arr, target, mid+1, right)
	} else {
		return binarySearchHelper(arr, target, left, mid-1)
	}
}

// InterpolationSearch searches for a target in a uniformly distributed sorted array
// Time Complexity: O(log log n) average, O(n) worst case, Space Complexity: O(1)
func InterpolationSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right && target >= arr[left] && target <= arr[right] {
		if left == right {
			if arr[left] == target {
				return left
			}
			return -1
		}

		// Calculate position using interpolation formula
		pos := left + ((target-arr[left])*(right-left))/(arr[right]-arr[left])

		if arr[pos] == target {
			return pos
		} else if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}

	return -1
}

// ExponentialSearch combines binary search with exponential growth
// Time Complexity: O(log n), Space Complexity: O(1)
func ExponentialSearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	if arr[0] == target {
		return 0
	}

	// Find range for binary search
	bound := 1
	for bound < len(arr) && arr[bound] < target {
		bound *= 2
	}

	// Perform binary search in the found range
	left := bound / 2
	right := bound
	if right >= len(arr) {
		right = len(arr) - 1
	}

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// JumpSearch searches by jumping through the array in fixed steps
// Time Complexity: O(√n), Space Complexity: O(1)
func JumpSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	// Calculate optimal jump size
	step := int(float64(n) * 0.5) // sqrt(n) approximation
	if step == 0 {
		step = 1
	}

	prev := 0

	// Jump through array
	for arr[min(step, n)-1] < target {
		prev = step
		step += int(float64(n) * 0.5)
		if prev >= n {
			return -1
		}
	}

	// Linear search in the identified block
	for i := prev; i < min(step, n); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TernarySearch searches in a sorted array by dividing into three parts
// Time Complexity: O(log₃ n), Space Complexity: O(1)
func TernarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid1 := left + (right-left)/3
		mid2 := right - (right-left)/3

		if arr[mid1] == target {
			return mid1
		}
		if arr[mid2] == target {
			return mid2
		}

		if target < arr[mid1] {
			right = mid1 - 1
		} else if target > arr[mid2] {
			left = mid2 + 1
		} else {
			left = mid1 + 1
			right = mid2 - 1
		}
	}

	return -1
}

// FindFirst finds the first occurrence of target in a sorted array with duplicates
func FindFirst(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			right = mid - 1 // Continue searching in left half
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// FindLast finds the last occurrence of target in a sorted array with duplicates
func FindLast(arr []int, target int) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			left = mid + 1 // Continue searching in right half
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
