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
func TestReverseList(t *testing.T) {
	// 测试用例
	testCases := []struct {
		head     *ListNode
		expected *ListNode
	}{
		{
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			expected: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
		},
	}

	for _, tc := range testCases {
		// 执行测试用例
		actual := reserveList1(tc.head)
		fmt.Println(actual)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	substring := lengthOfLongestSubstring("bbbbb")
	fmt.Println(substring)
}

func TestMax(t *testing.T) {
	water := maxArea([]int{1, 2, 12, 1, 1, 1, 1, 2, 5})
	fmt.Println(water)
}

func TestWinner(t *testing.T) {
	winner := isWinner([]int{4, 10, 7, 9}, []int{6, 5, 2, 3})
	fmt.Println(winner)
}

func TestNext(t *testing.T) {

}

func TestMaxHeap(t *testing.T) {
	h := NewMaxHeap()
	h.Push(4)
	h.Push(10)
	h.Push(8)
	h.Push(5)
	h.Push(1)

	for h.Len() > 0 {
		max, _ := h.Pop()
		fmt.Printf("%d ", max)
	}
}
func TestStr(t *testing.T) {
	i := str("aabdaabaaf", "aabaaf")
	fmt.Println(i)
}
