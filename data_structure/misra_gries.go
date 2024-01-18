package data_structure

import (
	"sort"
)

type Counter struct {
	Element string
	Count   int
}

func misraGries(stream []string, k int) []Counter {
	counters := make([]Counter, 0)
	for _, element := range stream {
		found := false
		for i := range counters {
			if counters[i].Element == element {
				counters[i].Count++
				found = true
				break
			}
		}

		if !found {
			if len(counters) < k-1 {
				counters = append(counters, Counter{Element: element, Count: 1})
			} else {
				for i := range counters {
					counters[i].Count--
					if counters[i].Count == 0 {
						counters = append(counters[:i], counters[i+1:]...)
						break
					}
				}
			}
		}
	}

	// 对计数器按照计数值降序排序
	sort.Slice(counters, func(i, j int) bool {
		return counters[i].Count > counters[j].Count
	})

	return counters
}

type LogCount struct {
	Log   string
	Count int64
}

func MisraGries(log []string, n int) map[string]int {
	counts := make(map[string]int, 0)
	for _, s := range log {
		if val, ok := counts[s]; ok {
			val += 1
			counts[s] = val
		} else {
			if len(counts) < n-1 {
				counts[s] = 1
			} else {
				for key, val := range counts {
					val -= 1
					if val == 0 {
						delete(counts, key)
					} else {
						counts[key] = val
					}
				}
			}
		}
	}
	return counts
}
