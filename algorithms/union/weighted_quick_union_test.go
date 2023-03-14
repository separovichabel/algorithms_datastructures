package union_test

import (
	"testing"

	"github.com/separovichabel/algorithms_datastructures/algorithms/union"
)

// TestNewWeightedQuickUnion_RootOf tests if RootOf returns the correct root
func TestNewWeightedQuickUnion_RootOf(t *testing.T) {
	t.Log("RootOf()")
	t.Run("No Union", func(t *testing.T) {
		ul := union.NewWeightedQuickUnion(1)
		root := ul.RootOf(0)

		if root != 0 {
			t.Errorf("Should return 0 but got %v", root)
		}
	})

	t.Run("1 Union", func(t *testing.T) {
		ul := union.NewWeightedQuickUnion(2)
		ul.Union(0, 1)

		root := ul.RootOf(1)

		if root != 0 {
			t.Errorf("RootOf(1) should return 0 but got %v", root)
		}
	})

	// RootOf chained elements 0 <- 1 <- 2
	t.Run("Chained Union", func(t *testing.T) {
		ul := union.NewWeightedQuickUnion(3)
		ul.Union(0, 1)
		ul.Union(1, 2)

		root := ul.RootOf(2)

		if root != 0 {
			t.Errorf("Should return 0 but got %v", root)
		}
	})

	// ConnectedDirect two diferent Chains 0 <- 1 and 2 <- 3
	t.Run("Union two diferent Chains", func(t *testing.T) {
		ul := union.NewWeightedQuickUnion(4)
		ul.Union(0, 1)
		ul.Union(2, 3)

		ul.Union(1, 2)

		rootOf3, rootOf1 := ul.RootOf(3), ul.RootOf(1)

		if rootOf3 != rootOf1 {
			t.Errorf("RootOf connected chains should be the same but got %v and %v", rootOf3, rootOf1)
		}
	})

	// ConnectedDirect two diferent weigh Chains 0 and 1 <- 2
	t.Run("Union two diferent weigh Chains", func(t *testing.T) {
		ul := union.NewWeightedQuickUnion(3)
		ul.Union(1, 2)

		ul.Union(0, 2)

		rootOf2 := ul.RootOf(2)

		if rootOf2 != 1 {
			t.Errorf("RootOf smaller chain should points to the larger, got 2 pointing to %v, expected 1", rootOf2)
		}
	})

	// ConnectedDirect two diferent weigh Chains 1 <- 2 and 0
	t.Run("Union two diferent weigh Chains", func(t *testing.T) {
		ul := union.NewWeightedQuickUnion(3)
		ul.Union(1, 2)

		ul.Union(2, 0)

		rootOf2 := ul.RootOf(2)

		if rootOf2 != 1 {
			t.Errorf("RootOf smaller chain should points to the larger, got 2 pointing to %v, expected 1", rootOf2)
		}
	})

}
