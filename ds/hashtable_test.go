package ds_test

import (
	"strconv"
	"testing"

	"github.com/hum/ds-algo/ds"
)

func TestGetAndSetInHashTable(t *testing.T) {
	table := ds.NewHashTable(50)

	table.Set("grapes", 1000)
	table.Set(1, 2000)
	table.Set("abc", 3000)
	table.Set("bbb", "15")

	tests := []struct {
		key   interface{}
		value interface{}
	}{
		{"grapes", 1000},
		{1, 2000},
		{"abc", 3000},
		{"bbb", "15"},
	}

	for _, tt := range tests {
		if value, ok := table.Get(tt.key); ok {
			if value != tt.value {
				t.Fatalf("Wrong value recieved from HashTable. Expected=%v, got=%v.", tt.value, value)
			}
		}
	}
}

func TestNonExistentValueInHashTable(t *testing.T) {
	table := ds.NewHashTable(40)

	tests := []struct {
		expected bool
	}{
		{false},
		{false},
		{false},
		{false},
		{false},
		{false},
		{false},
		{false},
	}

	for i, tt := range tests {
		if value, ok := table.Get(i); ok {
			t.Fatalf("Recieved non-existent value. Expected %v, but got=%v", tt.expected, value)
		}
	}
}

func TestKeysListHashTable(t *testing.T) {
	table := ds.NewHashTable(40)
	tries := 100

	for key := 0; key < tries; key++ {
		value := strconv.Itoa(key)
		table.Set(key, value)
	}

	if length := len(table.Keys()); length != tries {
		t.Fatalf("Expected to get %v different key values, but got %v instead.. ", tries, length)
	}
}
