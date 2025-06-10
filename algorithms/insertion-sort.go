package algorithms

func InsertionSort(src []int) {

	for i := range src {
		temp := src[i]
		j := i - 1

		for j >= 0 && src[j] > temp {
			src[j+1] = src[j]
			j--
		}
		src[j+1] = temp
	}
}
