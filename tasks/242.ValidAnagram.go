package tasks

import "fmt"

func ValidAnagramRun() {
	fmt.Println("242. Valid Anagram")
	fmt.Println("Link: https://leetcode.com/problems/valid-anagram/description/")
	fmt.Println("")
	
	s := "anagram"
	t := "nagaram"
	fmt.Println("Input: ", s, ",", t)
	result := isAnagram(s, t)
	fmt.Println("Output: ", result)

}

func isAnagram(s string, t string) bool {
    if len(t) != len(s) {
        return false
    }

    m := make(map[byte]int)

    for i := 0; i < len(s); i++ {
        m[s[i]]++
        m[t[i]]--
    }
    
    for _, c := range m {
        if c != 0 {
            return false
        }
    }
    return true
}