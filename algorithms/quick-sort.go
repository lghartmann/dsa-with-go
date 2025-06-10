package algorithms

func QuickSort(src []int, left, right int) {
	if left < right {
		pivot := right
		i := left - 1
		for j := left; j < pivot; j++ {
			if src[j] < src[pivot] {
				i++
				src[i], src[j] = src[j], src[i]
			}
		}
		src[i+1], src[pivot] = src[pivot], src[i+1]
		pivotIndex := i + 1

		QuickSort(src, left, pivotIndex-1)
		QuickSort(src, pivotIndex+1, right)
	}
}

func GPTQuickSort(src []int, left, right int) {
	if left < right {
		pivotIndex := partition(src, left, right)
		QuickSort(src, left, pivotIndex-1)
		QuickSort(src, pivotIndex+1, right)
	}
}

func partition(src []int, left, right int) int {
	pivot := src[right]
	i := left - 1
	for j := left; j < right; j++ {
		if src[j] < pivot {
			i++
			src[i], src[j] = src[j], src[i]
		}
	}
	src[i+1], src[right] = src[right], src[i+1]
	return i + 1
}
