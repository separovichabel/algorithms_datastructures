package sum

func Three(data []int) int {
	var sum int

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				if data[i]+data[j]+data[k] == 0 {
					sum++
				}
			}
		}
	}

	return sum
}
