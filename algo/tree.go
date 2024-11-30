package algo

import (
	"fmt"
	"math/rand"
	"strings"
)

type TreeNode struct {
	Value    int
	Children []TreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func Tree() {
	node := generateRandomTree(3, 4)
	printTree(node, 0, true)

	sum := sumTreeRecursive(node)
	fmt.Printf("sumTreeRecursive: The sum of the values of the tree nodes: %d\n", sum)

	sum = sumTreeIteration(node)
	fmt.Printf("sumTreeIteration: The sum of the values of the tree nodes: %d\n", sum)

	binaryTree := &BinaryTree{}

	iterations := rand.Intn(20)
	for i := 0; i < iterations; i++ {
		binaryTree.Insert(rand.Intn(20) + i)
	}

	binaryTree.printBinaryTree(binaryTree.Root, "", false)

}

func sumTreeRecursive(node TreeNode) int {
	sum := node.Value
	for _, child := range node.Children {
		sum += sumTreeRecursive(child)
	}
	return sum
}

func sumTreeIteration(node TreeNode) int {
	stack := []TreeNode{node}
	sum := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += current.Value

		if len(current.Children) > 0 {
			stack = append(stack, current.Children...)
		}
	}

	return sum
}

func generateRandomTree(depth, maxChildren int) TreeNode {
	node := TreeNode{
		Value: rand.Intn(50),
	}

	if depth <= 0 {
		return node
	}

	numChildren := rand.Intn(maxChildren + 1)
	for i := 0; i < numChildren; i++ {
		child := generateRandomTree(depth-1, maxChildren)
		node.Children = append(node.Children, child)
	}

	return node
}

func printTree(node TreeNode, depth int, isLast bool) {
	prefix := strings.Repeat("    ", depth)
	if depth > 0 {
		if isLast {
			prefix += "└── "
		} else {
			prefix += "├── "
		}
	}
	fmt.Printf("%s%d\n", prefix, node.Value)

	for i, child := range node.Children {
		printTree(child, depth+1, i == len(node.Children)-1)
	}
}

func (tree *BinaryTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &BinaryTreeNode{Value: value}
		return
	} else {
		node := tree.Root
		newNode := &BinaryTreeNode{Value: value}
		for {
			if value < node.Value {
				if node.Left == nil {
					node.Left = newNode
					return
				}
				node = node.Left
			} else {
				if node.Right == nil {
					node.Right = newNode
					return
				}
				node = node.Right
			}
		}
	}
}

func (tree *BinaryTree) printBinaryTree(node *BinaryTreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("└── ")
		prefix += "    "
	} else {
		fmt.Print("├── ")
		prefix += "│   "
	}
	fmt.Println(node.Value)

	tree.printBinaryTree(node.Left, prefix, false)
	tree.printBinaryTree(node.Right, prefix, true)
}
