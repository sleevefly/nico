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

func repeatLimitedString(s string, repeatLimit int) string {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	var ans []byte
	for i, j := 25, 24; i >= 0; i-- {
		j = min(j, i-1)
		for {
			for k := min(cnt[i], repeatLimit); k > 0; k-- {
				ans = append(ans, byte(i+'a'))
				cnt[i]--
			}
			if cnt[i] == 0 {
				break
			}
			for j >= 0 && cnt[j] == 0 {
				j--
			}
			if j < 0 {
				break
			}
			ans = append(ans, byte(j+'a'))
			cnt[j]--
		}
	}
	return string(ans)

}

func isWinner(player1 []int, player2 []int) int {
	if getSum(player1) == getSum(player2) {
		return 0
	}
	if getSum(player1) > getSum(player2) {
		return 1
	}
	return 2
}

func getSum(player1 []int) int {
	sum := 0
	for index, val := range player1 {
		if index == 0 {
			sum += val
			continue
		}
		if index == 1 {
			if player1[index-1] == 10 {
				sum += val * 2
			} else {
				sum += val
			}
			continue
		}
		if player1[index-1] == 10 || player1[index-2] == 10 {
			sum += val * 2
		} else {
			sum += val
		}

	}
	return sum

}

func isAnagram(s string, t string) bool {
	nums := [26]int{}
	for _, b := range []byte(s) {
		nums[b-'a']++
	}
	for _, b := range []byte(t) {
		nums[b-'a']--
	}
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	return true
}

func intersection(nums1 []int, nums2 []int) []int {
	hash := make(map[int]bool)
	for _, num := range nums1 {
		hash[num] = true
	}

	ans := []int{}
	for _, num := range nums2 {
		if _, ok := hash[num]; ok {
			ans = append(ans, num)
			delete(hash, num)
		}
	}

	return ans

}

/*func strStr(haystack string, needle string) int {
	// 边界情况处理
	if needle == "" {
		return 0
	}

	// 构建前缀表
	prefixTable := GetNext(needle)

	i, j := 0, 0 // 分别是haystack和needle的指针

	for i < len(haystack) {
		if haystack[i] == needle[j] {
			// 当前字符匹配，继续比较下一个字符
			i++
			j++

			if j == len(needle) {
				// 找到匹配的子串
				return i - j
			}
		} else if j > 0 {
			// 当前字符不匹配，回溯到前缀表中的位置
			j = prefixTable[j-1]
		} else {
			// 当前字符不匹配，且无法回溯，移动haystack指针
			i++
		}
	}

	return -1 // 未找到匹配的子串
}*/

func prefixTable(world string) []int {
	table := make([]int, len(world))
	length, i := 0, 1
	for i < len(world) {
		if world[length] == world[i] {
			length++
			table[i] = length
			i++
		} else if length > 0 {
			length = table[i-1]
		} else {
			table[i] = 0
			i++
		}
	}
	return table
}

func strStr(str, word string) int {
	if word == "" {
		return 0
	}
	table := prefixTable(word)
	j := 0
	for i := 0; i < len(str); {
		if str[i] == word[j] {
			i++
			j++
			if j == len(word)-1 {
				return i - j
			}
		} else if j > 0 {
			j = table[j-1]
		} else {
			i++
		}
	}

	return -1
}
