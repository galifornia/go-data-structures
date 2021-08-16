package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		return -1
	}

	root := h.array[0]
	last := len(h.array) - 1

	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)

	return root
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while has at least one child
	for l <= lastIndex {
		// when only left child
		if l == lastIndex {
			childToCompare = l
			// when left child is larger
		} else if h.array[l] > h.array[r] {
			childToCompare = l
			// when right child is larger
		} else {
			childToCompare = r
		}

		// swap if bigger child is bigger than root
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// get te parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return i*2 + 1
}

// get the right child index
func right(i int) int {
	return i*2 + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	maxHeap := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 6, 7, 43, 35}

	fmt.Println("Creating heap from array of ints:", buildHeap)
	for _, v := range buildHeap {
		maxHeap.Insert(v)
		fmt.Println((maxHeap))
	}

	fmt.Println("-------Extracting test-------")
	for i := 0; i < 5; i++ {
		max := maxHeap.Extract()
		fmt.Println(max, "extracted, heap is now:", maxHeap)
	}
}
