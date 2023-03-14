package union

type QuickUnion struct {
	array []int
}

func (ul *QuickUnion) Len() int {
	return len(ul.array)
}

func (ul *QuickUnion) Connected(a, b int) bool {
	return ul.RootOf(a) == ul.RootOf(b)
}

func (ul *QuickUnion) ConnectedDirect(a, b int) bool {
	return a == ul.array[b]
}

// Union connects two itens
// b will be connected to a
func (ul *QuickUnion) Union(a, b int) {
	if ul.Connected(a, b) {
		return
	}

	ul.array[b] = ul.RootOf(a)
}

// RootOf finds itens root
func (ul *QuickUnion) RootOf(iten int) int {
	for iten != ul.array[iten] {
		iten = ul.array[iten]
	}

	return iten
}

func NewQuickUnion(len int) *QuickUnion {
	ul := &QuickUnion{}
	ul.array = make([]int, len)

	for i := range ul.array {
		ul.array[i] = i
	}

	return ul
}
