package data_structure_test

import (
	"testing"

	ds "nico/data_structure"
)

func TestBinSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 9, 10}
	if ds.BinSearch(nums, 9) != 5 {
		t.Fatalf("bin search result mismatch")
	}
}

func TestMooreVoted(t *testing.T) {
	nums := []int{1, 1, 2, 2, 1, 1, 10}
	if ds.MooreVoted(nums) != 1 {
		t.Fatalf("moore voted result mismatch")
	}
}

func TestMaxHeap(t *testing.T) {
	h := ds.NewMaxHeap()
	h.Push(4)
	h.Push(10)
	max, _ := h.Pop()
	if max != 10 {
		t.Fatalf("max heap pop mismatch")
	}
}

func TestMisra(t *testing.T) {
	stream := []string{"a", "b", "a", "a", "c", "a"}
	if len(ds.MisraGries(stream, 3)) == 0 {
		t.Fatalf("expected non-empty result")
	}
}
