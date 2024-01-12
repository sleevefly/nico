package data_structure

func reserveList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	pre = nil
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// leetcode 11
func maxArea(height []int) int {
	result, left, right := 0, 0, len(height)-1
	for left < right {
		lower := max(right, left) - min(right, left)
		h := min(height[left], height[right])
		result = max(result, lower*h)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return result
}

// 2085
func countWords(words1 []string, words2 []string) int {
	wordsCount1 := make(map[string]int)
	wordsCount2 := make(map[string]int)
	for _, s := range words2 {
		if val, ok := wordsCount2[s]; ok {
			val += 1
			wordsCount2[s] = val
		} else {
			wordsCount2[s] = 1
		}
	}
	for _, s := range words1 {
		if val, ok := wordsCount1[s]; ok {
			val += 1
			wordsCount1[s] = val
		} else {
			wordsCount1[s] = 1
		}
	}
	count := 0
	for key, val := range wordsCount2 {
		if val == 1 && wordsCount1[key] == 1 {
			count += 1
		}
	}

	return count
}
