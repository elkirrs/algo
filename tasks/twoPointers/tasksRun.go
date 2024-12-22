package twoPointers

func TasksRun() {
	//removeDuplicates()
	//removeElement()
	//merge()
	//mergeTwoLists()
	//sortedSquares()
	head := &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 6, Next: &ListNode{
					Val: 3, Next: &ListNode{
						Val: 4, Next: &ListNode{
							Val: 5, Next: &ListNode{
								Val: 6, Next: nil}}}}}}}
	val := 6
	removeElements(head, val)
}
