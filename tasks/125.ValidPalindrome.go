package tasks

import (
	"fmt"
)

func ValidPalindromeRun() {
	fmt.Println("125. Valid Palindrome")
	fmt.Println("https://leetcode.com/problems/valid-palindrome/description/")
	fmt.Println("")

	validPalindromePrint("A man, a plan, a canal: Panama")
	validPalindromePrint("race a car")
	validPalindromePrint("a.")
	validPalindromePrint("0P")
}

func isPalindrome(s string) bool {
	if s == " " || len(s) == 1 {
		return true
	}

	start := 0
	end := len(s) - 1

	for start <= end {
		left := s[start]
		right := s[end]

		if !isAlphaNum(left) {
			start++
			continue
		}

		if !isAlphaNum(right) {
			end--
			continue
		}

		if toLower(left) != toLower(right) {
			return false
		}

		start++
		end--
	}

	return true
}

func isAlphaNum(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func validPalindromePrint(in string) {
	fmt.Println("Input:", in)
	output := isPalindrome(in)
	fmt.Println("Output:", output)
	fmt.Println("")
}
