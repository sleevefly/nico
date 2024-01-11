package data_structure

import (
	"fmt"
	"testing"
)

func TestBinSearch(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 9, 10}
	search := BinSearch(nums, 11)
	fmt.Println(search)
}

func TestMooreVoted(t *testing.T) {
	nums := []int{1, 1, 2, 2, 1, 1, 10}
	voted := MooreVoted(nums)
	fmt.Println(voted)
}
