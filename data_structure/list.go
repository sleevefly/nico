package data_structure

type ListNode struct {
	Val  int
	Next *ListNode
}

func reserveList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	pre = nil
	cur := head
	for cur != nil {
		//先保存下个节点
		tmp := cur.Next
		//转向
		cur.Next = pre
		//下个节点的指针域需要指向的节点
		pre = cur
		//下个阶段
		cur = tmp

	}
	// 返回反转后的链表头节点
	return pre
}

func lengthOfLongestSubstring1(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	slowIndex, subLength := -1, 0
	m := make(map[string]int)
	for fastIndex := 0; fastIndex < len(s); fastIndex++ {
		if val, ok := m[string(s[fastIndex])]; ok {
			slowIndex = val + 1
			m[string(s[fastIndex])] = fastIndex
			if fastIndex-slowIndex > subLength {
				subLength = fastIndex - slowIndex + 1
			}
		} else {
			m[string(s[fastIndex])] = fastIndex
			if fastIndex-slowIndex > subLength {
				subLength = fastIndex - slowIndex + 1
			}
		}
	}
	if subLength > 0 {
		return subLength
	}
	return 0
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		//判断是否相等
		if cur.Next.Val == cur.Next.Next.Val {
			tmp := cur.Next.Next.Val
			//删除操作
			for cur.Next != nil && cur.Next.Val == tmp {
				cur.Next = cur.Next.Next
			}
		} else {
			//移动
			cur = cur.Next
		}
	}
	return dummy.Next
}

func deleteItem(head *ListNode, target int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == target {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
