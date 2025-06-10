package algorithms

func SelectionSort(src []int) []int {
	for i := range src {
		min := i
		for j := i + 1; j < len(src); j++ {
			if src[j] < src[min] {
				min = j
			}
		}
		if min != i {
			tmp := src[i]
			src[i] = src[min]
			src[min] = tmp
		}
	}

	return src
}
