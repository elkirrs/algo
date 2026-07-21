package tasks

func MinimumWindowSubstringRun() {
	PrintTask(
		"76. Minimum Window Substring",
		"https://leetcode.com/problems/minimum-window-substring/description/",
	)

	PrintScriptTwo(minWindow, "ADOBECODEBANC", "ABC") // BANC
	PrintScriptTwo(minWindow, "a", "a")               // a
	PrintScriptTwo(minWindow, "a", "aa")              // ""
}

func minWindow(s string, t string) string {
	l, r := 0, 0
	count := len(t)
	req := make(map[byte]int)
	start := l
	minLen := len(s) + 1

	for _, v := range t {
		req[byte(v)]++
	}

	for _, v := range s {
		if req[byte(v)] > 0 {
			count--
		}
		req[byte(v)]--
		r++

		for count == 0 {
			if r-l < minLen {
				minLen = r - l
				start = l
			}
			req[s[l]]++
			if req[s[l]] > 0 {
				count++
			}
			l++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}
