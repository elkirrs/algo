package tasks

import "fmt"

func Info() {
	fmt.Println("")
}

func PrintTask(name, link string) {
	fmt.Println(name)
	fmt.Println(link)
}

func PrintScript[T any, R any](fn func(T) R, arg T) {
	fmt.Println("")
	fmt.Print("Input: ")
	printValue(arg)

	fmt.Print("Output: ")
	printValue(fn(arg))
}

func PrintScriptTwo[A, B any, R any](fn func(A, B) R, a A, b B) {
	fmt.Println("")
	fmt.Print("Input: ")
	printValue(a)
	printValue(b)

	fmt.Print("Output: ")
	printValue(fn(a, b))
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func printValue(v any) {
	switch t := v.(type) {
	case *ListNode:
		printList(t)
	default:
		fmt.Println(v)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrepareList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}
