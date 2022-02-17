package main

import "fmt"

const ArraySize = 7

// HashTable struct
type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	Next *bucketNode
}

// Insert for hashTable
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search for hashTable
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete for hashTable
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *bucket) insert(key string) {
	if !b.search(key) {
		n := &bucketNode{key: key}
		n.Next = b.head
		b.head = n
		return
	}

	fmt.Printf("%s is alreadt exists", key)
	return
}

func (b *bucket) search(key string) bool {
	current := b.head
	for current == nil { // !!!
		if current.key == key {
			return true
		}
		current = current.Next
	}

	return false
}

func (b *bucket) delete(key string) {
	// check match head
	if b.head.key == key {
		b.head = b.head.Next
		return
	}

	prev := b.head
	for prev.Next != nil {
		if prev.Next.key == key {
			prev.Next = prev.Next.Next
		}
		prev = prev.Next
	}

	return
}

func hash(key string) int {
	sum := 0
	for _, c := range key {
		sum += int(c)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	t := Init()
	fmt.Println(t)
	fmt.Println(hash("RANDY"))

	tb := &bucket{}
	tb.insert("RANDY")
	fmt.Println(tb.search("RANDY"))
	fmt.Println(tb.search("ERIC"))
}
