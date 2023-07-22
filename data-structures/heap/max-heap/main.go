package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds a value to the heap
func (h *MaxHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest value and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	lastIndex := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("Heap is empty")
		return -1
	}
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.maxHeapifyDown(0)
	return extracted
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	left, right := leftChild(index), rightChild(index)
	childToCompare := 0
	for left < lastIndex {
		if left == lastIndex { // when left child is the only child
			childToCompare = left
		} else if h.array[left] > h.array[right] { // when left child is greater
			childToCompare = left
		} else { // when right child is greater
			childToCompare = right
		}
		// compare array value of current index to largest child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			left, right = leftChild(index), rightChild(index)
		} else {
			return
		}
	}
}

// maxHeapifyUp will heapify the array up
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Get the parent index
func parent(index int) int {
	return (index - 1) / 2
}

// Get the left child index
func leftChild(index int) int {
	return 2*index + 1
}

// Get the right child index
func rightChild(index int) int {
	return 2*index + 2
}

// Swap keys in the array
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func main() {
	maxHeap := MaxHeap{}
	arr := []int{1, 2, 3, 4, 5, 6, 8, 9, 11, 7}
	for _, v := range arr {
		maxHeap.Insert(v)
		fmt.Println(maxHeap)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(maxHeap.Extract())
		fmt.Println(maxHeap)
	}
}
