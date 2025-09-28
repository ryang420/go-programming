package io

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ReadInts reads integers from standard input
func ReadInts(prompt string) ([]int, error) {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return nil, fmt.Errorf("failed to read input")
	}

	line := strings.TrimSpace(scanner.Text())
	if line == "" {
		return []int{}, nil
	}

	parts := strings.Fields(line)
	result := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid integer: %s", part)
		}
		result[i] = num
	}

	return result, nil
}

// ReadInt reads a single integer from standard input
func ReadInt(prompt string) (int, error) {
	fmt.Print(prompt)

	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		return 0, fmt.Errorf("failed to read integer: %v", err)
	}

	return num, nil
}

// ReadString reads a string from standard input
func ReadString(prompt string) (string, error) {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", fmt.Errorf("failed to read input")
	}

	return strings.TrimSpace(scanner.Text()), nil
}

// PrintArray prints an array in a formatted way
func PrintArray(arr []int, label string) {
	if label != "" {
		fmt.Printf("%s: ", label)
	}
	fmt.Print("[")
	for i, val := range arr {
		fmt.Print(val)
		if i < len(arr)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

// PrintMatrix prints a 2D matrix
func PrintMatrix(matrix [][]int, label string) {
	if label != "" {
		fmt.Printf("%s:\n", label)
	}

	for _, row := range matrix {
		fmt.Print("[")
		for i, val := range row {
			fmt.Printf("%3d", val)
			if i < len(row)-1 {
				fmt.Print(", ")
			}
		}
		fmt.Println("]")
	}
}

// GenerateRandomArray creates an array of random integers
func GenerateRandomArray(size, min, max int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(max-min+1) + min
	}

	return arr
}

// GenerateSortedArray creates a sorted array with optional duplicates
func GenerateSortedArray(size, min, max int, allowDuplicates bool) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)

	if allowDuplicates {
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(max-min+1) + min
		}
	} else {
		if size > (max - min + 1) {
			panic("Cannot generate array without duplicates: size exceeds range")
		}

		// Generate unique numbers
		used := make(map[int]bool)
		for i := 0; i < size; i++ {
			for {
				num := rand.Intn(max-min+1) + min
				if !used[num] {
					arr[i] = num
					used[num] = true
					break
				}
			}
		}
	}

	// Sort the array
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

// ReadFromFile reads integers from a file
func ReadFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	var result []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid integer in file: %s", part)
			}
			result = append(result, num)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return result, nil
}

// WriteToFile writes integers to a file
func WriteToFile(filename string, arr []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i, val := range arr {
		if i > 0 {
			writer.WriteString(" ")
		}
		writer.WriteString(strconv.Itoa(val))
	}
	writer.WriteString("\n")

	return nil
}

// Benchmark function for timing operations
func Benchmark(name string, fn func()) time.Duration {
	fmt.Printf("Running %s...\n", name)
	start := time.Now()
	fn()
	duration := time.Since(start)
	fmt.Printf("%s completed in: %v\n", name, duration)
	return duration
}
