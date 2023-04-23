package main

import (
	"crud/model"
	"fmt"
)

type Pair struct {
	key   string
	value model.User
}

type HashTable struct {
	data [][]Pair
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		data: make([][]Pair, size),
	}
}

func (t *HashTable) Hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])*i) % len(t.data)
	}

	return hash
}

func (t *HashTable) Set(key string, value model.User) {
	address := t.Hash(key)
	t.data[address] = append(t.data[address], Pair{key: key, value: value})
}

func (t *HashTable) Get(key string) (model.User, error) {
	address := t.Hash(key)
	currentBucket := t.data[address]

	if len(currentBucket) == 0 {
		return model.User{}, fmt.Errorf("key not allocated")
	}

	for _, bucket := range currentBucket {
		if bucket.key == key {
			return bucket.value, nil
		}
	}

	return model.User{}, fmt.Errorf("key not found")
}
