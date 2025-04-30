package ht

import "errors"

var (
	errKeyNotFound = errors.New("key not found")
)

type node struct {
	key   int
	value string
	next  *node
}

type HashTable struct {
	size    int
	buckets []*node
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([]*node, size),
	}
}

func (h *HashTable) hash(key int) int {
	return key % h.size
}

func (h *HashTable) Get(key int) (returnVal string, err error) {
	hashedKey := h.hash(key)
	val := h.buckets[hashedKey]

	if val == nil {
		return "", errKeyNotFound
	}

	for val != nil {
		if val.key == key {
			return val.value, nil
		}
		val = val.next
	}
	return "", errKeyNotFound
}
