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
