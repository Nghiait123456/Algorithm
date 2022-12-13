package main

type MinHeap struct {
	arr []int
}

func newMinHeap(arr []int) *MinHeap {
	minHeap := &MinHeap{
		arr: arr,
	}

	return minHeap
}

func (m *MinHeap) leftChildIndex(index int) int {
	return 2*index + 1
}

func (m *MinHeap) rightChildIndex(index int) int {
	return 2*index + 2
}

func (m *MinHeap) swap(first, second int) {
	temp := m.arr[first]
	m.arr[first] = m.arr[second]
	m.arr[second] = temp
}

func (m *MinHeap) leaf(index int, size int) bool {
	if index >= (size/2) && index <= size {
		return true
	}

	return false
}
