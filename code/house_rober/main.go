package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeFromArray(array []int) *TreeNode {
	root := &TreeNode{
		Val: array[0],
	}

	currentNode := root
	direction := false
	for i := 1; i < len(array)/2; i++ {
		currentNode.Left = array[i]
		currentNode.Right = array[i+1]
	}
	return root
}

func insertToTree(node *TreeNode, array []int, i, level int) {
	index = math.Pow(2, level) + i
	if index > len(array) {
		return
	}
	currentNode.Left = array[index]
	if index+1 > len(array) {
		return
	}
	currentNode.Right = array[indexi+1]
	insertToTree()
}

func main() {
	a := []int{}
	fmt.Println(rob(treeFromArray(a)))
}

func rob(root *TreeNode) int {
	max := startRobbery(root, 0)
	left, right := 0, 0
	if root.Left != nil {
		left = startRobbery(root.Left, 0)
	}
	if root.Right != nil {
		right = startRobbery(root.Right, 0)
	}
	if max < left+right {
		max = left + right
	}
	return max
}

func startRobbery(node *TreeNode, stolen int) int {
	stolen = stolen + node.Val
	left, right := 0, 0
	if node.Left != nil {
		left = skipRobbery(node.Left, stolen)
	}
	if node.Right != nil {
		right = skipRobbery(node.Right, stolen)
	}
	return stolen + left + right
}

func skipRobbery(node *TreeNode, stolen int) int {
	var ll, lr int
	if node.Left != nil {
		ll = startRobbery(node.Left, stolen)
	}
	if node.Left.Right != nil {
		lr = startRobbery(node.Right, stolen)
	}
	if ll > lr {
		return ll
	}
	return lr
}
