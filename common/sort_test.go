package common

import (
	"reflect"
	"testing"
)

func TestCeshi(t *testing.T) {

}

func TestOrder(t *testing.T) {

}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minNumberOfValidStrings(t *testing.T) {
	type args struct {
		words  []string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfValidStrings(tt.args.words, tt.args.target); got != tt.want {
				t.Errorf("minNumberOfValidStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minValidStrings(t *testing.T) {
	type args struct {
		words  []string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minValidStrings(tt.args.words, tt.args.target); got != tt.want {
				t.Errorf("minValidStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	freq := make(map[byte]int) // 使用哈希表记录字符出现次数
	left, right := 0, 0
	maxLength := 0

	for right < n {
		freq[s[right]]++         // 右指针右移，更新频率
		for freq[s[right]] > 1 { // 如果出现重复字符
			freq[s[left]]-- // 左指针右移，直到移除重复字符
			left++
		}
		maxLength = max(maxLength, right-left+1) // 更新最大长度
		right++
	}

	return maxLength
}
