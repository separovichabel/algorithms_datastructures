package union

type QuickUnion struct {
	Parent []int // Parent of b is Parent[b]
}

func (ul *QuickUnion) Connected(a, b int) bool {
	return ul.RootOf(a) == ul.RootOf(b)
}

// Union connects two itens
// b will be connected to a
func (ul *QuickUnion) Union(a, b int) {
	if ul.Connected(a, b) {
		return
	}

	ul.Parent[b] = a
}

// RootOf finds itens root
func (ul *QuickUnion) RootOf(iten int) int {
	for iten != ul.Parent[iten] {
		iten = ul.Parent[iten]
	}

	return iten
}

func NewQuickUnion(len int) *QuickUnion {
	ul := &QuickUnion{}
	ul.Parent = make([]int, len)

	for i := range ul.Parent {
		ul.Parent[i] = i
	}

	return ul
}
