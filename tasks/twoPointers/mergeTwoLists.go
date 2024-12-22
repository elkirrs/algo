package twoPointers

//list1 = [1,2,4]
//list2 = [1,3,4]
//Output[1,1,2,3,4,4]

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
