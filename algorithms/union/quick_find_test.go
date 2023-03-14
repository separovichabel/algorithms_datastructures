package union_test

import (
	"testing"

	quickFind "github.com/separovichabel/algorithms_datastructures/algorithms/union"
)

// TestNewQuickFind tests if NewQuickFind returns the correct list len
func TestNewQuickFind(t *testing.T) {
	ul := quickFind.NewQuickFind(3)
	l := ul.Len()

	if l != 3 {
		t.Errorf("NewQuickFind(3) should len 3 but got %v", l)
	}
}

// TestNewQuickFindDefaultValues verifies if NewQuickFind returns non unioned values only
func TestNewQuickFindDefaultValues(t *testing.T) {
	ul := quickFind.NewQuickFind(3)
	connected1, connected2, connected3 := ul.Connected(0, 1), ul.Connected(0, 2), ul.Connected(1, 2)

	if connected1 || connected2 || connected3 {
		t.Errorf("NewQuickFind(3) should not return any unioned value")
	}
}

func TestQuickFindList_Union(t *testing.T) {
	ul := quickFind.NewQuickFind(10)
	ul.Union(5, 6)

	resp := ul.Connected(5, 6)
	if resp == false {
		t.Errorf("Union should connect elements")
	}
}

func TestQuickFindList_UnionRepeat(t *testing.T) {
	ul := quickFind.NewQuickFind(10)
	ul.Union(5, 6)
	ul.Union(6, 5)

	resp := ul.Connected(5, 6)
	if resp == false {
		t.Errorf("Union called multiple times with the same values should should not change its behaviour")
	}
}

// Union Union should no unite non asked values
func TestQuickFindListUnionNonConnecteds(t *testing.T) {
	ul := quickFind.NewQuickFind(3)
	ul.Union(0, 1)

	resp := ul.Connected(1, 2)
	if resp == true {
		t.Errorf("Union should not connect elements that where not united")
	}
}

func TestQuickFindListUnionMultiple(t *testing.T) {
	ul := quickFind.NewQuickFind(3)
	ul.Union(0, 1)
	ul.Union(1, 2)

	resp := ul.Connected(0, 2)
	if resp == false {
		t.Errorf("Union should connect elements that has common unions")
	}
}
