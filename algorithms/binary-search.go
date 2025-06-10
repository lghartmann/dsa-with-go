package algorithms

func GptBinarySearch(src []int, target int) int {
	left, right := 0, len(src)-1
	for left <= right {
		mid := left + (right-left)/2
		if src[mid] == target {
			return mid
		} else if src[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearch(src []int, target int) int {
	length := len(src)
	pos := length / 2
	pivot := src[pos]

	if pivot == target {
		return target
	}

	if length == 1 {
		return -1
	}

	if pivot > target {
		newArr := src[:pos]
		return BinarySearch(newArr, target)
	}

	newArr := src[pos:]
	return BinarySearch(newArr, target)
}
