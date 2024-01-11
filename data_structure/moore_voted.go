package data_structure

//摩尔投票，用于找出集合中超过一半元素

func MooreVoted(nums []int) int {
	num, count := -1, 1
	for _, i := range nums {
		if i != num {
			if count > 0 {
				count -= 1
			}
			if count == 0 {
				count = 1
				num = i
			}

		}
		if num == i {
			count++
		}
	}
	return num
}
