package tasks


import (
	"fmt"
	"strconv"
)


func EncodeAndDecodeStringsRun() {
	fmt.Println("Task: 271. Encode and Decode Strings")
	fmt.Println("Link: https://leetcode.com/problems/encode-and-decode-strings/description/")
	fmt.Println("")

	arrStr := []string{"Hello", "World"}
	fmt.Println("Input: ", arrStr)
	s := Solution{}
	encoded := s.Encode(arrStr)
	fmt.Println("Encoded: ", encoded)

	decoded := s.Decode(encoded)
	fmt.Println("Decoded: ", decoded)
}

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded string

	for _, str := range strs {
		encoded += strconv.Itoa(len(str)) + "#" + str
	}

	return encoded
}

func (s *Solution) Decode(encoded string) []string {
	var res []string

	i := 0
	for i < len(encoded) {

		j := i
		for encoded[j] != '#' {
			j++
		}

		l, _ := strconv.Atoi(encoded[i:j])
		res = append(res, encoded[j+1:j+1+l])
		i = j+1+l

	}

	return res
}
