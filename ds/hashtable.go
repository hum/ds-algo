package ds

import (
	"fmt"
)

type HashTable struct {
	size int
	data []interface{}
}

type HashNode struct {
	key   interface{}
	value interface{}
}

func NewHashTable(size int) *HashTable {
	table := make([]interface{}, size)

	for i := 0; i < size; i++ {
		table[i] = make([]HashNode, 0, size)
	}
	return &HashTable{size: size, data: table}
}

func (h *HashTable) Keys() []interface{} {
	keys := make([]interface{}, 0, h.size)

	for i := 0; i < len(h.data); i++ {
		nodes := h.data[i].([]HashNode)

		if len(nodes) > 0 {
			for _, v := range nodes {
				keys = append(keys, v.key)
			}
		}
	}
	return keys
}

func (h *HashTable) Set(key, value interface{}) {
	hashedKey := h.hashKey(key)
	nodes := h.data[hashedKey].([]HashNode)
	h.data[hashedKey] = append(nodes, HashNode{key, value})
}

func (h *HashTable) Get(key interface{}) (interface{}, bool) {
	nodes := h.data[h.hashKey(key)].([]HashNode)

	if len(nodes) > 0 {
		for _, node := range nodes {
			if node.key == key {
				return node.value, true
			}
		}
	}
	return nil, false
}

func (h *HashTable) hashKey(key interface{}) int {
	hash := 0
	runes := []rune(fmt.Sprint(key))

	for i := 0; i < len(runes); i++ {
		hash = (hash + int(runes[i])*i) % h.size
	}
	return hash
}
