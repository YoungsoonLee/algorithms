package main

import "fmt"

// Heap is a really complete tree
// it means the parent have 2 childs or null
// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// NewMaxHeap ...
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(k int) {
	h.array = append(h.array, k)
	h.heapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and remove it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.heapifyDown(0)
	return extracted
}

func (h *MaxHeap) heapifyUp(index int) {
	for h.array[parentIndex(index)] < h.array[index] {
		h.swap(parentIndex(index), index)
		index = parentIndex(index)
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := leftIndex(index), rightIndex(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = leftIndex(index), rightIndex(index)
		} else {
			return
		}

	}
}

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftIndex(i int) int {
	return (i * 2) + 1
}

func rightIndex(i int) int {
	return (i * 2) + 2
}

func main() {
	test := NewMaxHeap()
	fmt.Println(test)

	test.Insert(5)
	test.Insert(6)
	test.Insert(1)
	test.Insert(10)
	test.Insert(8)
	test.Insert(9)
	fmt.Println(test)

	for i := 0; i < 3; i++ {
		test.Extract()
		fmt.Println(test)
	}
}
