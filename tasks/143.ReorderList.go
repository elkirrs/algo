package tasks

func ReorderListRun() {
	PrintTask(
		"143. Reorder List",
		"https://leetcode.com/problems/reorder-list/description/",
	)

	PrintScript(reorderList, PrepareList([]int{1, 2, 3, 4}))
	PrintScript(reorderList, PrepareList([]int{1, 2, 3, 4, 5}))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) *ListNode {
	slow, fast := head, head

	// find mid list node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverce rigth list node
	var second *ListNode
	for slow.Next != nil {
		next := slow.Next.Next
		slow.Next.Next = second
		second = slow.Next
		slow.Next = next
	}
	slow.Next = nil

	first := head

	// join both list in queue l1(1), l2(1), l1(2), l2(2), l1(3), l2(3) .....
	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next

		first.Next = second
		second.Next = tmp1

		first = tmp1
		second = tmp2
	}

	return head
}
