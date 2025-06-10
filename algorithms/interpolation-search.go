package algorithms

func InterpolationSearch(src []int, target int) int {
	end := len(src) - 1
	start := 0

	for target >= src[start] && target <= src[end] && start <= end {
		probe := start + (end-start)*(target-src[start])/(src[end]-src[start])

		if src[probe] == target {
			return probe
		}

		if src[probe] < target {
			start = probe + 1
		} else {
			end = probe - 1
		}
	}

	return -1
}
