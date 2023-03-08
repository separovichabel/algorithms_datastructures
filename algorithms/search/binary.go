package search

func Binary(l []int, t int) int {
	// Adding 2 cursors defining the limits of the search
	left, right := 0, len(l)-1

	// interacting until de cursors get inverted
	for left <= right {

		// get the middle between two cursors
		mid := (left + right) / 2

		// if the middle is the desired, than, returns it
		if l[mid] == t {
			return mid
		} else if t > l[mid] {

			// if the targuet is greather than the middle, than we can ignore
			// the elements on the left of the middle. So the left cursor will
			// be the middle + 1.
			left = mid + 1

		} else {
			// otherwise, the targuet is smaller than the middle. So we will
			// ignore de right side.
			right = mid - 1
		}
	}
	// if the cursors get inverted it means that there is no value iquals to the target
	// so -1 is returned to indicate the absence of the target.
	return -1
}
