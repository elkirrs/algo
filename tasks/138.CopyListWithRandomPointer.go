package tasks

func CopyListWithRandomPointerRun() {
	PrintTask(
		"138. Copy List with Random Pointer",
		"https://leetcode.com/problems/copy-list-with-random-pointer/description/",
	)

	list := &Node{
		Val: 7,
		Next: &Node{
			Val: 13,
			Next: &Node{
				Val: 11,
				Next: &Node{
					Val: 10,
					Next: &Node{
						Val:    1,
						Next:   &Node{},
						Random: &Node{Val: 0},
					},
					Random: &Node{Val: 2},
				},
				Random: &Node{Val: 4},
			},
			Random: &Node{Val: 0},
		},
		Random: &Node{},
	}
	PrintScript(copyRandomList, list)
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	oldToCopy := make(map[*Node]*Node, 0)

	curr := head
	for curr != nil {
		oldToCopy[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head
	for curr != nil {
		copy := oldToCopy[curr]
		copy.Next = oldToCopy[curr.Next]
		copy.Random = oldToCopy[curr.Random]
		curr = curr.Next
	}

	return oldToCopy[head]
}
