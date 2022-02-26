package main

import "fmt"

// MaHeap
type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(k int) {
	h.array = append(h.array, k)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	if len(h.array) == 0 {
		return -1
	}

	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	// because heap is compeletly binary tree
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
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main() {
	m := &MaxHeap{}
	h := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range h {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

}
