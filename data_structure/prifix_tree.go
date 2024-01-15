package data_structure

func str(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	table := make([]int, len(needle))

	for j, i := 0, 1; i < m; i++ {

		for j > 0 && needle[i] != needle[j] {
			j = table[j-1]
		}
		if needle[j] == needle[i] {
			j++
		}
		table[i] = j

	}

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && needle[j] != haystack[i] {
			j = table[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func preTable(text string) []int {
	table := make([]int, len(text))
	for length, i := 0, 1; i < len(text); i++ {
		for length > 0 && text[length] != text[i] {
			length = table[length-1]
		}
		if text[length] == text[i] {
			length++
		}
		table[i] = length
	}
	return table
}

func strV1(str, text string) {
	table := preTable(text)
	for i, j := 0, 0; i < len(str); i++ {
		for j > 0 && text[j] != str[i] {
			j = table[j-1]
		}
		if text[j] == str[i] {
			i++
			j++
		}
		if i == len(str)-1 {
			return
		}
	}
}
