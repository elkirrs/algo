package tasks

func AddTwoNumbersRun() {
	PrintTask(
		"2. Add Two Numbers",
		"https://leetcode.com/problems/add-two-numbers/description/",
	)

	PrintScriptTwo(addTwoNumbers, PrepareList([]int{2, 4, 3}), PrepareList([]int{5, 6, 4}))                // [7,0,8]
	PrintScriptTwo(addTwoNumbers, PrepareList([]int{9, 9, 9, 9, 9, 9, 9}), PrepareList([]int{9, 9, 9, 9})) // [7,0,8]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	demmy := &ListNode{}
	curr := demmy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return demmy.Next
}
