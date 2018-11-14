package hashtable

import (
	"github.com/herobs/algorithms-in-go/datastructures"
	"github.com/herobs/algorithms-in-go/datastructures/linkedlist"
)

// Trivial hash table implementation
// just demonstrate for how hash table works
// how to design a hash function
// how to resolve collisions

const (
	defaultHashTableSize = 1024
)

// HashNode hash node store in bucket
type HashNode struct {
	key   string
	value interface{}
}

// HashTable trivial hash table
type HashTable struct {
	keys    map[string]int
	buckets [defaultHashTableSize]datastructures.List
}

// New create new hash table
func New() *HashTable {
	table := &HashTable{
		keys: make(map[string]int),
	}

	for i := 0; i < defaultHashTableSize; i++ {
		table.buckets[i] = linkedlist.New()
	}

	return table
}

// hash simple hash function
// use string as key, sum up all characters code point
// then mod hash table size
func (table *HashTable) hash(key string) int {
	sum := 0
	for c := range key {
		sum += int(c)
	}

	return sum % defaultHashTableSize
}

// Set value with key
func (table *HashTable) Set(key string, value interface{}) {
	keyHash := table.hash(key)
	table.keys[key] = keyHash

	bucket := table.buckets[keyHash]
	node := bucket.Find(nil, func(v interface{}) bool {
		return v.(*HashNode).key == key
	})

	if node == nil {
		bucket.Append(&HashNode{
			key:   key,
			value: value,
		})
	} else {
		node.SetValue(value)
	}
}

// Get value with key
func (table *HashTable) Get(key string) interface{} {
	bucket := table.buckets[table.hash(key)]
	node := bucket.Find(nil, func(v interface{}) bool {
		return v.(*HashNode).key == key
	})

	if node == nil {
		return nil
	}

	return node.Value().(*HashNode).value
}

// Has check key exists
func (table *HashTable) Has(key string) bool {
	_, ok := table.keys[key]

	return ok
}
