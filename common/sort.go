package common

import (
	"fmt"
	"math"
	"strings"
)

type Sort interface {
	int64 | int | int32 | uint32
}

func Order[T Sort](args []T) T {

	return args[0]
}

func minValidStrings(words []string, target string) int {
	prefixFunction := func(word, target string) []int {
		s := word + "#" + target
		n := len(s)
		pi := make([]int, n)
		for i := 1; i < n; i++ {
			j := pi[i-1]
			for j > 0 && s[i] != s[j] {
				j = pi[j-1]
			}
			if s[i] == s[j] {
				j++
			}
			pi[i] = j
		}
		return pi
	}

	n := len(target)
	back := make([]int, n)
	for _, word := range words {
		pi := prefixFunction(word, target)
		m := len(word)
		for i := 0; i < n; i++ {
			back[i] = int(math.Max(float64(back[i]), float64(pi[m+1+i])))
		}
	}

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = int(1e9)
	}
	for i := 0; i < n; i++ {
		dp[i+1] = dp[i+1-back[i]] + 1
		if dp[i+1] > n {
			return -1
		}
	}
	return dp[n]
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for ind, num := range nums {
		if val, ok := m[target-num]; ok {
			return []int{val, num}
		} else {
			m[num] = ind
		}
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 创建一个新的链表头节点，用于存放结果
	dummyHead := &ListNode{}
	// current 指针用于遍历新链表并添加节点
	current := dummyHead
	// carry 用于记录进位
	carry := 0

	// 当 l1 和 l2 都不为空时，或者 carry 不为 0 时，继续循环
	for l1 != nil || l2 != nil || carry != 0 {
		// 获取 l1 和 l2 当前节点的值，如果为空则默认为 0
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}

		// 计算当前位的和以及进位
		sum := x + y + carry
		carry = sum / 10  // 计算进位
		digit := sum % 10 // 计算当前位的数字

		// 创建新的节点并添加到结果链表中
		current.Next = &ListNode{Val: digit}
		current = current.Next

		// 移动 l1 和 l2 到下一个节点
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	// 返回结果链表的头节点（不包含 dummyHead）
	return dummyHead.Next
}

func minNumberOfValidStrings(words []string, target string) int {
	n := len(target)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for _, word := range words {
			m := len(word)
			if m <= i && strings.HasPrefix(target[:i], word) { // 正确的前缀判断
				if i == m {
					dp[i] = 1
				} else if dp[i-m] != math.MaxInt32 {
					dp[i] = min(dp[i], dp[i-m]+1)
				}
			}
		}
	}

	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Ceshi() {
	words := []string{"ab", "abc", "abcd"}
	target := "abcdab"
	fmt.Println(minNumberOfValidStrings(words, target)) // Output: 2

	words = []string{"ab", "abc", "bc"}
	target = "abcd"
	fmt.Println(minNumberOfValidStrings(words, target)) //Output:-1

	words = []string{"a", "b", "c"}
	target = "abc"
	fmt.Println(minNumberOfValidStrings(words, target)) //Output:3

}
