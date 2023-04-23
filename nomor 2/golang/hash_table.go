package main

import (
	"fmt"
)

type HashTable struct {
	data [][]string
}

func HashTableGo(size int) *HashTable {
	return &HashTable{make([][]string, size)}
}
func (t *HashTable) _hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])*i) % len(t.data)
	}
	return hash
}
func (t *HashTable) Set(key, value string) {
	address := t._hash(key)
	if t.data[address] == nil {
		t.data[address] = []string{}
	}
	t.data[address] = append(t.data[address], key, value)
}
func (t *HashTable) Get(key string) string {
	address := t._hash(key)
	currentBucket := t.data[address]
	if currentBucket != nil {
		for i := 0; i < len(currentBucket); i += 2 {
			if currentBucket[i] == key {
				return currentBucket[i+1]
			}
		}
	}
	return ""
}
func main() {
	myHashTable := HashTableGo(100)
	myHashTable.Set("082124606606", "Rony Setyawan")
	fmt.Println(myHashTable.Get("082124606606"))
}
