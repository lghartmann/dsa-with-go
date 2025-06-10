package algorithms

func BubbleSort(src []int) {

	for i := range src {
		for j := range len(src) - i - 1 {
			if src[j] > src[j+1] {
				src[j+1], src[j] = src[j], src[j+1]
			}
		}
	}
}
