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
		actual := reserveList(tc.head)
		fmt.Println(actual)
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	substring := lengthOfLongestSubstring("bbbbb")
	fmt.Println(substring)
}
