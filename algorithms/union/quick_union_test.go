package union_test

import (
	"testing"

	quickUnion "github.com/separovichabel/algorithms_datastructures/algorithms/union"
)

// TestNewQuickUnion_RootOf tests if RootOf returns the correct root
func TestNewQuickUnion_RootOf(t *testing.T) {
	t.Log("RootOf()")
	t.Run("No Union", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(1)
		root := ul.RootOf(0)

		if root != 0 {
			t.Errorf("Should return 0 but got %v", root)
		}
	})

	t.Run("1 Union", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(2)
		ul.Union(0, 1)

		root := ul.RootOf(1)

		if root != 0 {
			t.Errorf("RootOf(1) should return 0 but got %v", root)
		}
	})

	// RootOf chained elements 0 <- 1 <- 2
	t.Run("Chained Union", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(3)
		ul.Union(0, 1)
		ul.Union(1, 2)

		root := ul.RootOf(2)

		if root != 0 {
			t.Errorf("Should return 0 but got %v", root)
		}
	})

	// ConnectedDirect two diferent Chains 0 <- 1 and 2 <- 3
	t.Run("Union two diferent Chains", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(4)
		ul.Union(0, 1)
		ul.Union(2, 3)

		ul.Union(1, 2)

		rootOf3, rootOf1 := ul.RootOf(3), ul.RootOf(1)

		if rootOf3 != rootOf1 {
			t.Errorf("RootOf connected chains should be the same but got %v and %v", rootOf3, rootOf1)
		}
	})

}

// TestNewQuickUnion_ConnectedDirect tests if a value is directly connected to another
func TestNewQuickUnion_ConnectedDirect(t *testing.T) {
	t.Log("ConnectedDirect()")
	t.Run("No Union", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(1)
		connected := ul.ConnectedDirect(0, 0)

		if !connected {
			t.Errorf("Should return true but got %v", connected)
		}
	})

	t.Run("1 Union", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(2)
		ul.Union(0, 1)

		connected := ul.ConnectedDirect(0, 1)

		if !connected {
			t.Errorf("Should return true but got %v", connected)
		}
	})

	// RootOf chained elements 0 <- 1 <- 2
	t.Run("Chained Union", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(3)
		ul.Union(0, 1)
		ul.Union(1, 2)

		connected := ul.ConnectedDirect(1, 2)

		if connected {
			t.Errorf("Should return false but got %v", connected)
		}
	})

}

// TestNewQuickUnion_Connected tests if a value is connected to another
func TestNewQuickUnion_Connected(t *testing.T) {
	t.Log("Connected()")
	t.Run("No Union", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(1)
		connected := ul.Connected(0, 0)

		if !connected {
			t.Errorf("Should return true but got %v", connected)
		}
	})

	t.Run("1 Union", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(2)
		ul.Union(0, 1)

		connected := ul.Connected(0, 1)

		if !connected {
			t.Errorf("Should return true but got %v", connected)
		}
	})

	// RootOf chained elements 0 <- 1 <- 2
	t.Run("Chained Union", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(3)
		ul.Union(0, 1)
		ul.Union(1, 2)

		connected := ul.Connected(1, 2)

		if !connected {
			t.Errorf("Should return true but got %v", connected)
		}
	})

	// ConnectedDirect two diferent Chains 0 <- 1 and 2 <- 3 and 0 <- 2
	t.Run("Union two diferent Chains", func(t *testing.T) {
		ul := quickUnion.NewQuickUnion(4)
		ul.Union(0, 1)
		ul.Union(2, 3)

		ul.Union(0, 2)

		connected := ul.Connected(3, 1)

		if !connected {
			t.Errorf("Should return true but got %v", connected)
		}
	})

}
