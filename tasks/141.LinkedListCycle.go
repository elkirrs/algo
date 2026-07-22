package tasks

func LinkedListCycleRun() {
	PrintTask(
		"141. Linked List Cycle",
		"https://leetcode.com/problems/linked-list-cycle/description/",
	)

	list := []int{3,2,0,-4}
	PrintScript(hasCycle, PrepareList(list))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}