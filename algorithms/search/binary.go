package search

func Binary(l []int, t int) int {
	left, right := 0, len(l)-1

	for left <= right {
		mid := (left + right) / 2

		if l[mid] == t {
			return mid
		} else if t > l[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
