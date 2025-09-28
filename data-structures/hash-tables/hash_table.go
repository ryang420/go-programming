package hashtables

import (
	"fmt"
	"strings"
)

// Entry represents a key-value pair in the hash table
type Entry struct {
	Key   string
	Value int
	Next  *Entry // For chaining collision resolution
}

// HashTable represents a hash table with separate chaining
type HashTable struct {
	buckets []*Entry
	size    int
	count   int
}

// NewHashTable creates a new hash table with the specified size
func NewHashTable(size int) *HashTable {
	if size <= 0 {
		size = 16 // Default size
	}
	return &HashTable{
		buckets: make([]*Entry, size),
		size:    size,
		count:   0,
	}
}

// hash computes the hash value for a given key
func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash*31 + int(char)) % ht.size
	}
	if hash < 0 {
		hash += ht.size
	}
	return hash
}

// Put inserts or updates a key-value pair in the hash table
func (ht *HashTable) Put(key string, value int) {
	index := ht.hash(key)

	if ht.buckets[index] == nil {
		ht.buckets[index] = &Entry{Key: key, Value: value}
		ht.count++
		return
	}

	// Handle collision using chaining
	current := ht.buckets[index]
	for current != nil {
		if current.Key == key {
			// Update existing key
			current.Value = value
			return
		}
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	// Add new entry at the end of the chain
	current.Next = &Entry{Key: key, Value: value}
	ht.count++
}

// Get retrieves the value associated with the given key
func (ht *HashTable) Get(key string) (int, bool) {
	index := ht.hash(key)
	current := ht.buckets[index]

	for current != nil {
		if current.Key == key {
			return current.Value, true
		}
		current = current.Next
	}

	return 0, false
}

// Delete removes the key-value pair with the given key
func (ht *HashTable) Delete(key string) bool {
	index := ht.hash(key)

	if ht.buckets[index] == nil {
		return false
	}

	// If the first entry matches
	if ht.buckets[index].Key == key {
		ht.buckets[index] = ht.buckets[index].Next
		ht.count--
		return true
	}

	// Search in the chain
	current := ht.buckets[index]
	for current.Next != nil {
		if current.Next.Key == key {
			current.Next = current.Next.Next
			ht.count--
			return true
		}
		current = current.Next
	}

	return false
}

// Contains checks if the hash table contains the given key
func (ht *HashTable) Contains(key string) bool {
	_, found := ht.Get(key)
	return found
}

// Size returns the number of key-value pairs in the hash table
func (ht *HashTable) Count() int {
	return ht.count
}

// IsEmpty checks if the hash table is empty
func (ht *HashTable) IsEmpty() bool {
	return ht.count == 0
}

// Keys returns all keys in the hash table
func (ht *HashTable) Keys() []string {
	keys := make([]string, 0, ht.count)

	for _, entry := range ht.buckets {
		current := entry
		for current != nil {
			keys = append(keys, current.Key)
			current = current.Next
		}
	}

	return keys
}

// Values returns all values in the hash table
func (ht *HashTable) Values() []int {
	values := make([]int, 0, ht.count)

	for _, entry := range ht.buckets {
		current := entry
		for current != nil {
			values = append(values, current.Value)
			current = current.Next
		}
	}

	return values
}

// Clear removes all entries from the hash table
func (ht *HashTable) Clear() {
	ht.buckets = make([]*Entry, ht.size)
	ht.count = 0
}

// LoadFactor returns the load factor of the hash table
func (ht *HashTable) LoadFactor() float64 {
	return float64(ht.count) / float64(ht.size)
}

// String returns a string representation of the hash table
func (ht *HashTable) String() string {
	if ht.IsEmpty() {
		return "HashTable: {}"
	}

	var result strings.Builder
	result.WriteString("HashTable: {\n")

	for i, entry := range ht.buckets {
		if entry != nil {
			result.WriteString(fmt.Sprintf("  [%d]: ", i))
			current := entry
			for current != nil {
				result.WriteString(fmt.Sprintf("(%s: %d)", current.Key, current.Value))
				if current.Next != nil {
					result.WriteString(" -> ")
				}
				current = current.Next
			}
			result.WriteString("\n")
		}
	}

	result.WriteString("}")
	return result.String()
}

// SimpleHashMap represents a simple hash map using Go's built-in map
type SimpleHashMap struct {
	data map[string]int
}

// NewSimpleHashMap creates a new simple hash map
func NewSimpleHashMap() *SimpleHashMap {
	return &SimpleHashMap{
		data: make(map[string]int),
	}
}

// Put inserts or updates a key-value pair
func (shm *SimpleHashMap) Put(key string, value int) {
	shm.data[key] = value
}

// Get retrieves the value associated with the given key
func (shm *SimpleHashMap) Get(key string) (int, bool) {
	value, found := shm.data[key]
	return value, found
}

// Delete removes the key-value pair with the given key
func (shm *SimpleHashMap) Delete(key string) bool {
	if _, found := shm.data[key]; found {
		delete(shm.data, key)
		return true
	}
	return false
}

// Contains checks if the map contains the given key
func (shm *SimpleHashMap) Contains(key string) bool {
	_, found := shm.data[key]
	return found
}

// Size returns the number of key-value pairs
func (shm *SimpleHashMap) Size() int {
	return len(shm.data)
}

// Keys returns all keys in the map
func (shm *SimpleHashMap) Keys() []string {
	keys := make([]string, 0, len(shm.data))
	for key := range shm.data {
		keys = append(keys, key)
	}
	return keys
}
