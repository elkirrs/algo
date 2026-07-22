package tasks

func RemoveNthNodeFromEndOfListRun() {
	PrintTask(
		"19. Remove Nth Node From End of List",
		"https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/",
	)

	PrintScriptTwo(removeNthFromEnd, PrepareList([]int{1, 2, 3, 4, 5}), 2) // [1,2,3,5]
	PrintScriptTwo(removeNthFromEnd, PrepareList([]int{1}), 1)             // []
	// PrintScriptTwo(removeNthFromEnd, PrepareList([]int{1, 2}), 1)          // [1]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	slow, fast := dummy, dummy

	n++
   	for n != 0 {
		n--
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	
	return dummy.Next
}
