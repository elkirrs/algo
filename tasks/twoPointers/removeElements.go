package twoPointers

//head := &ListNode{
//	Val: 1,
//	Next: &ListNode{
//		Val: 2,
//		Next: &ListNode{
//			Val: 6,
//			Next: &ListNode{
//				Val: 3,
//				Next: &ListNode{
//					Val: 4,
//					Next: &ListNode{
//						Val: 5,
//						Next: Next: &ListNode{
//							Val: 6,
//							Next: nil
//						}
//					}
//				}
//			}
//		}
//	}
//}
//val := 6

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
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
