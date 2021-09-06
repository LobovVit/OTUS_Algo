package task05

import "math"

func swap(mas []int, a int, b int) {
	var c int
	c = mas[a]
	mas[a] = mas[b]
	mas[b] = c
}

func parent(i int) int {
	return int(math.Floor(float64((i - 1) / 2)))
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}

func drown(heap []int, i int, size int) {
	l := leftChild(i)
	r := rightChild(i)
	largest := i
	if l < size && heap[l] > heap[largest] {
		largest = l
	}
	if r < size && heap[r] > heap[largest] {
		largest = r
	}
	if largest != i {
		swap(heap, i, largest)
		drown(heap, largest, size)
	}
}

func buildHeap(mas []int) {
	for i := len(mas)/2 - 1; i >= 0; i-- {
		drown(mas, i, len(mas))
	}
}

func HeapSort(mas []int) {
	//size := len(mas)
	buildHeap(mas)
	for i := len(mas) - 1; i >= 0; i-- {
		//size = size -1
		swap(mas, 0, i)
		drown(mas, 0, i)
	}
}
