package datastructures

import (
	"testing"
	// "fmt"
)

func TestHashTable(t *testing.T) {
	table := CreateHashTable()
	
	// Testing a hashtable set
	table.set("A", 2)
	table.set("B", -1)

	if table.size != 2 {
		t.Errorf(`
Incorrect hashtable size: %v
Should be: %v`, table.size, 2)
	}

	// Testing a hashtable get
	thingy, err := table.get("A")
	if len(err) != 0 {
		t.Errorf(err)
	}
	if thingy != 2 {
		t.Errorf(`
Failed to get "%v" from the hashtable.
Instead got "%v"`, "A", thingy)
	}

	// Test resizing (will break if too many)
	table.set("C", 4)
	table.set("D", 0)
	table.set("E", -5)
	table.set("F", 1)
	table.set("G", 0)

	// Test verious settings
	thingy, err = table.get("a")
	if err == "" {
		t.Errorf(`
Retrieved an item that wasn't set: %v
Item key: %v`, thingy, "a")
	}

	// Test hashtable setting an already set value
	preSize := table.size
	table.set("A", 0)
	thingy, err = table.get("A")
	if err != "" {
		t.Errorf("")
	}
	if table.size != preSize {
		t.Errorf(`
Table size increased for already set value: %v
Expected: %v`, table.size, preSize)
	}
	if thingy != 0 {
		t.Errorf(`
Failed to reset a current value: %v
Instead got: %v`, 0, thingy)
	}

	// Use the print function to test hash distribution
	// table.print()
}