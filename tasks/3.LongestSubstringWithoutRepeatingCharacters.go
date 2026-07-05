package tasks

func LongestSubstringWithoutRepeatingCharactersRun() {
	PrintTask(
		"3. Longest Substring Without Repeating Characters",
		"https://leetcode.com/problems/longest-substring-without-repeating-characters/description/",
	)

	PrintScript(lengthOfLongestSubstring, "abcabcbb") // 3
	PrintScript(lengthOfLongestSubstring, "bbbbb")    // 1
	PrintScript(lengthOfLongestSubstring, "pwwkew") //3
	PrintScript(lengthOfLongestSubstring, "dvdf") // 3
	PrintScript(lengthOfLongestSubstring, "") // 0
	PrintScript(lengthOfLongestSubstring, " ") // 1
}

func lengthOfLongestSubstring(s string) int {
	count := 0
	left := 0
	last := make(map[byte]int, 0)

	for right := 0; right < len(s); right++ {
		if idx, exist := last[s[right]]; exist && idx >= left {
			left = idx + 1
		}

		last[s[right]] = right
		if (right - left + 1) > count {
			count = right - left + 1 
		}
	}

	return count
}
