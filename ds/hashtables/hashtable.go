package hashtables

import "errors"

var (
	ErrKeyNotFound  = errors.New("key not found")
	ErrDuplicateKey = errors.New("key is duplicate")
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
	current := h.buckets[hashedKey]

	if current == nil {
		return "", ErrKeyNotFound
	}

	for current != nil {
		if current.key == key {
			return current.value, nil
		}
		current = current.next
	}
	return "", ErrKeyNotFound
}

func (h *HashTable) Set(key int, value string) {
	hashedKey := h.hash(key)
	current := h.buckets[hashedKey]

	newNode := &node{
		key:   key,
		value: value,
		next:  nil,
	}

	if current == nil {
		h.buckets[hashedKey] = newNode
		return
	}

	for current.next != nil { //if there are other nodes
		if current.key == key {
			current.value = value
			return
		}
		current = current.next
	}

	if current.key == key {
		current.value = value
		return
	}

	current.next = newNode
}
