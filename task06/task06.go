package task06

//мердж сорт и квик сорт
func MergeSort(list []int) []int {
	count := len(list)

	switch {
	case count > 2:
		lb := MergeSort(list[:count/2])
		rb := MergeSort(list[count/2:])
		list = append(lb, rb...)
		lastIndex := len(list) - 1

		for i := 0; i < lastIndex; i++ {
			mv := list[i]
			mi := i

			for j := i + 1; j < lastIndex+1; j++ {
				if mv > list[j] {
					mv = list[j]
					mi = j
				}
			}

			if mi != i {
				list[i], list[mi] = list[mi], list[i]
			}
		}

	case len(list) > 1 && list[0] > list[1]:
		list[0], list[1] = list[1], list[0]
	}

	return list
}
