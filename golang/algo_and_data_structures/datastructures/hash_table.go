package datastructures

import (
	"fmt"
)

// HashTable is the defining struct for the object
type HashTable struct {
	size int
	table []*HashTableEntry
}

// HashTableEntry is a single entry in the hash table
type HashTableEntry struct {
	key string
	value int
}

// Prints the hashtable
func (hashTable *HashTable) print() {
	fmt.Println(hashTable.table)
}

// Gets the value of a give key
func (hashTable *HashTable) get(key string) (int, string) {
	index := hash(key, hashTable) 

	// While our hash index doesn't give us a nil value, move along
	for hashTable.table[index] != nil {

		// If we don't find the key, keep moving.
		if hashTable.table[index].key != key {
			index = (index + 1) % len(hashTable.table)
		} else {

			// If we do find the key, then we're done
			return hashTable.table[index].value, ""
		}
	}

	return 0, "Current key not set"
}

// Sets a given value with a set key
func (hashTable *HashTable) set(key string, value int) {
	
	// Resize and get the current index
	resizeIfNeeded(hashTable)
	index := hash(key, hashTable)

	// Continue till either you find an empty spot...
	for hashTable.table[index] != nil {

		// Or you find the entry...
		if hashTable.table[index].key == key {
			hashTable.table[index].value = value
			return	// And we're done
		}
		
		// Move along to the next index
		index = (index + 1) % len(hashTable.table)
	}

	// At this point we need to create a new entry
	hashTable.table[index] = &HashTableEntry{
		key: key,
		value: value,
	}

	hashTable.size++
}

// Spits out an index based on the string key going in
func hash(key string, hashTable *HashTable) int {
	var result int = 1

	// Simply multiply the ASCII values of a string together
	for _, char := range key {
		result = (result * (int(char) * 5831))	// 5831? djb2
	}
	
	return result % len(hashTable.table)
}

// Resizes the table when we reach a significant size
func resizeIfNeeded(hashTable *HashTable) {
	if hashTable.size + 1 > (len(hashTable.table) / 2 ){

		// Create a new set of memory for the next table
		// We create a new table to exploit it's set method
		newTable := HashTable{
			size: hashTable.size,
			table: make([]*HashTableEntry, len(hashTable.table) * 2),
		}

		// Get each of the entries from the original table and set them in the new one
		for _, entry := range hashTable.table {
			if entry != nil {
				newTable.set(entry.key, entry.value)
			}
		}
		
		// Finally reassign the new table over
		hashTable.table = newTable.table
	}
}

// CreateHashTable returns a hashtable of values
func CreateHashTable(size ...int) (HashTable) {
	var hashTable HashTable
	if (len(size) > 0) {
		hashTable = HashTable{
			size: 0,
			table: make([]*HashTableEntry, size[0]),
		}
	} else {
		hashTable = HashTable{
			size: 0,
			table: make([]*HashTableEntry, 8),
		}
	}

	return hashTable
}