package task04

// сортировка вставками
// сортировка шела
// бенч марки для частично упорядоченых или для очень не упрорядоченых
// для шела - поэкспериментировать с гэп сиквенсом
func swap(mas []int, a int, b int) {
	var c int
	c = mas[a]
	mas[a] = mas[b]
	mas[b] = c
}

func SortBubble(a []int) {
	var len = len(a)
	for true {
		swapped := false
		for i := 1; i < len; i++ {
			if a[i-1] > a[i] {
				swap(a, i-1, i)
				swapped = true
			}
		}
		len = len - 1
		if !swapped {
			break
		}
	}
}

func SortSelection(a []int) {
	for i, _ := range a {
		min := i
		for j := i; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if i != min {
			swap(a, i, min)
		}
	}

}

func SortInsertion(a []int) {
	for i, _ := range a {
		x := a[i]
		j := i - 1
		for true {
			if !(j >= 0 && a[j] > x) {
				break
			}
			a[j+1] = a[j]
			j = j - 1
			a[j+1] = x
		}
	}
}

func SortShell(a []int) {
	for k := len(a) / 2; k >= 1; k = k / 2 {
		for i := 0; i < len(a); i = i + k {
			x := a[i]
			j := i - 1
			for true {
				if !(j >= 0 && a[j] > x) {
					break
				}
				a[j+1] = a[j]
				j = j - 1
				a[j+1] = x
			}
		}
	}
}
