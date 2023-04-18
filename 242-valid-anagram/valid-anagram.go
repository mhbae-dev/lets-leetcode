package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freq := make(map[rune]int)

	for _, char := range s {
		freq[char]++
	}

	for _, char := range t {
		if freq[char] == 0 {
			return false
		}
		freq[char]--
	}
	return true
}
