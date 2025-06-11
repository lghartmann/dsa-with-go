package algorithms

func merge(leftArr, rightArr, mainArr []int) {
	leftSize := len(mainArr) / 2
	rightSize := len(mainArr) - leftSize

	i, l, r := 0, 0, 0

	for l < leftSize && r < rightSize {
		if leftArr[l] < rightArr[r] {
			mainArr[i] = leftArr[l]
			i++
			l++
		} else {
			mainArr[i] = rightArr[r]
			i++
			r++
		}
	}

	for l < leftSize {
		mainArr[i] = leftArr[l]
		i++
		l++
	}

	for r < rightSize {
		mainArr[i] = rightArr[r]
		i++
		r++
	}

}

func MergeSort(src []int) {
	length := len(src)
	if length <= 1 {
		return
	}

	middle := length / 2

	leftArr := make([]int, middle)
	rightArr := make([]int, length-middle)

	leftIndex := 0
	rightIndex := 0

	for ; leftIndex < length; leftIndex++ {
		if leftIndex < middle {
			leftArr[leftIndex] = src[leftIndex]
		} else {
			rightArr[rightIndex] = src[leftIndex]
			rightIndex++
		}
	}
	MergeSort(leftArr)
	MergeSort(rightArr)
	merge(leftArr, rightArr, src)
}
