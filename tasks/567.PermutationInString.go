package tasks

func PermutationInStringRun() {
	PrintTask(
		"567. Permutation in String",
		"https://leetcode.com/problems/permutation-in-string/",
	)

	PrintScriptTwo(checkInclusion, "ab", "eidbaooo")        // true
	PrintScriptTwo(checkInclusion, "ab", "eidboaoo")        // false
	PrintScriptTwo(checkInclusion, "hello", "ooolleoooleh") // false
	PrintScriptTwo(checkInclusion, "ab", "ab")              // true
	PrintScriptTwo(checkInclusion, "tor", "actor")          // true
}

func checkInclusion(s1 string, s2 string) bool {
	l, r := 0, 0

	freqS2 := [26]int{}
	freqS1 := [26]int{}

	for _, c := range s1 {
		freqS1[byte(c)-'a']++
	}

	for _, s := range s2 {
		freqS2[byte(s)-'a']++

		if r-l+1 > len(s1) {
			freqS2[s2[l]-'a']--
			l++
		}
		r++

		if freqS1 == freqS2 {
			return true
		}
	}

	return false
}
