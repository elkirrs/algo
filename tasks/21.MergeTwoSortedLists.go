package tasks

import (
	"fmt"
)

func MergeTwoSortedListsRun() {
	fmt.Println("21. Merge Two Sorted Lists")
	fmt.Println("https://leetcode.com/problems/merge-two-sorted-lists/")
	fmt.Println("")

	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	fmt.Print("Input: l1 ")
	printList(l1)
	fmt.Print("Input: l2 ")
	printList(l2)
	result := mergeTwoLists(l1, l2)
	fmt.Print("Output: ")
	printList(result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	tmp := &ListNode{}
	current := tmp

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return tmp.Next
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
