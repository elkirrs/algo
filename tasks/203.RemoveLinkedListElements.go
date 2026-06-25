package tasks

import (
	"fmt"
)

func RemoveLinkedListElementsRun() {
	fmt.Println("203. Remove Linked List Elements")
	fmt.Println("https://leetcode.com/problems/remove-linked-list-elements/description/")
	fmt.Println("")

	l1 := &ListNode2{
		Val: 1, Next: &ListNode2{
			Val: 2, Next: &ListNode2{
				Val: 6, Next: &ListNode2{
					Val: 3, Next: &ListNode2{
						Val: 4, Next: &ListNode2{
							Val: 5, Next: &ListNode2{
								Val: 6}}}}}}}

	fmt.Println("Input: ")
	printList2(l1)
	result := removeElements(l1, 6)
	fmt.Println("Output: ")
	printList2(result)	

}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func removeElements(head *ListNode2, val int) *ListNode2 {
	dummy := &ListNode2{}
	dummy.Next = head
	current := dummy

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}

func printList2(head *ListNode2) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

