package union

type QuickFind struct {
	array []int
}

func (ul *QuickFind) Len() int {
	return len(ul.array)
}

func (ul *QuickFind) Connected(a, b int) bool {
	return ul.array[a] == ul.array[b]
}

func (ul *QuickFind) Union(a, b int) {
	if ul.array[a] == ul.array[b] {
		return
	}

	for i, v := range ul.array {
		if v == ul.array[b] {
			ul.array[i] = ul.array[a]
		}
	}
}

func NewQuickFind(len int) *QuickFind {
	ul := &QuickFind{}
	ul.array = make([]int, len)

	for i := range ul.array {
		ul.array[i] = i
	}

	return ul
}
