package task04

import (
	"math/rand"
	"otus_algo/task05"
	"sort"
	"testing"
)

func generateMass(len int) []int {
	ret := make([]int, len)
	for i := 0; i < len; i++ {
		ret[i] = rand.Intn(len * 10)
	}
	return ret
}

func BenchmarkSortBubble(b *testing.B) {
	var a = generateMass(20000)
	SortBubble(a)
}

func BenchmarkSortSelection(b *testing.B) {
	var a = generateMass(20000)
	SortSelection(a)
}

func BenchmarkSortInsertion(b *testing.B) {
	var a = generateMass(20000)
	SortInsertion(a)
}

func BenchmarkSortShell(b *testing.B) {
	var a = generateMass(20000)
	SortShell(a)
}

func BenchmarkGo(b *testing.B) {
	var a = generateMass(20000)
	sort.Ints(a)
}

func BenchmarkHeap(b *testing.B) {
	var a = generateMass(20000)
	task05.HeapSort(a)
}

func TestAll(t *testing.T) {
	var a = generateMass(20000)
	var b = make([]int, len(a))
	copy(b, a)
	SortShell(a)
	sort.Ints(b)
	for i, _ := range a {
		if a[i] != b[i] {
			t.Errorf("не сошлись %d = %d ", a[i], b[i])
		}
	}
}
