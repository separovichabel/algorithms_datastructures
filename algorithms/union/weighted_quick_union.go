package union

type WeightedQuickUnion struct {
	Sets   int   // Number of sets
	Parent []int // Parent of b is Parent[b]
	Size   []int // Size of b is Size[b]
}

func (ul *WeightedQuickUnion) Connected(a, b int) bool {
	return ul.RootOf(a) == ul.RootOf(b)
}

// Union connects two itens
// b will be connected to a
func (ul *WeightedQuickUnion) Union(a, b int) {
	rootA := ul.RootOf(a)
	rootB := ul.RootOf(b)

	// skip if already connected
	if rootA == rootB {
		return
	}

	// connect smaller tree to larger tree
	if ul.Size[rootA] < ul.Size[rootB] {
		ul.Parent[rootA] = rootB
		ul.Size[rootB] += ul.Size[rootA]
	} else {
		ul.Parent[rootB] = rootA
		ul.Size[rootA] += ul.Size[rootB]
	}

	// decrease number of sets
	ul.Sets--
}

// RootOf finds itens root
func (ul *WeightedQuickUnion) RootOf(iten int) int {
	for iten != ul.Parent[iten] {
		iten = ul.Parent[iten]
	}

	return iten
}

func NewWeightedQuickUnion(len int) *WeightedQuickUnion {
	ul := &WeightedQuickUnion{}
	ul.Parent = make([]int, len)
	ul.Size = make([]int, len)
	ul.Sets = len

	for i := range ul.Parent {
		ul.Parent[i] = i
		ul.Size[i] = 1
	}

	return ul
}
