package tasks

func LongestRepeatingCharacterReplacementRun() {
	PrintTask(
		"424. Longest Repeating Character Replacement",
		"https://leetcode.com/problems/longest-repeating-character-replacement/description/",
	)

	PrintScriptTwo(characterReplacement, "ABAB", 2)    // 4
	PrintScriptTwo(characterReplacement, "AABABBA", 1) // 4
}

func characterReplacement(s string, k int) int {
	l, r := 0, 0
	count := 0
	freq := make(map[byte]int, 0)

	for _, c := range s {
		freq[byte(c)]++

		if freq[byte(c)] > count {
			count = freq[byte(c)]
		}

		if (r-l+1)-count > k {
			freq[s[l]]--
			l++
		}
		r++
	}

	return r - l
}
