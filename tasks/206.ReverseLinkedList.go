package tasks

func ReverseLinkedListRun() {
	PrintTask(
		"206. Reverse Linked List",
		"https://leetcode.com/problems/reverse-linked-list/description/",
	)

	list := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3, Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	PrintScript(reverseList, list)
}


// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
