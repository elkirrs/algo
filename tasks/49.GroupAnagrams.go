package tasks

import (
	"fmt"
	"slices"
)

func GroupAnagramsRun() {
	fmt.Println("Task: 49. Group Anagrams")
	fmt.Println("Link: https://leetcode.com/problems/group-anagrams/description/")
	fmt.Println("")

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("Input: ", strs)
	result := groupAnagrams(strs)
	fmt.Println("Output: ", result)	
}


func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)

    for _, world := range strs {
        chars := []rune(world)
        slices.Sort(chars)
        key := string(chars)
        groups[key] = append(groups[key], world)
    }

    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}