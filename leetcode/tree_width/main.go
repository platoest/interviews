package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeWidth struct {
	treeMap [][]int
	root    *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	treeWidth := TreeWidth{
		root: root,
	}
	return treeWidth.processTree()
}

func (width *TreeWidth) processTree() int {
	root := width.root
	if root.Left == nil {
		if root.Right != nil {
			width.root = root.Right
			return width.processTree()
		}
		return 1
	}
	if root.Right == nil {
		if root.Left != nil {
			width.root = root.Left
			return width.processTree()
		}
		return 1
	}
	// now we have left and right and can start processing
	width.treeMap = [][]int{0: {0}}
	width.process(root.Left, 1, 0)
	width.process(root.Right, 1, 1)
	return width.calculateMaxWidth()
}

func (width *TreeWidth) calculateMaxWidth() int {
	result := 1
	for i := 0; i < len(width.treeMap); i++ {
		if len(width.treeMap[i]) < 2 {
			continue
		}
		newResult := width.treeMap[i][len(width.treeMap[i])-1] - width.treeMap[i][0] + 1
		if newResult > result {
			result = newResult
		}
	}
	return result
}

func (width *TreeWidth) process(node *TreeNode, level, newIndex int) {
	if len(width.treeMap)-1 < level {
		widthRow := make([]int, 2)
		widthRow[0] = newIndex
		widthRow[1] = newIndex + 1
		width.treeMap = append(width.treeMap, widthRow)
	} else {
		width.treeMap[level][1] = newIndex
	}
	if node.Left != nil {
		width.process(node.Left, level+1, newIndex*2)
	}
	if node.Right != nil {
		width.process(node.Right, level+1, newIndex*2+1)
	}
}

func main() {
	root := &TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{},
			Right: &TreeNode{},
		},
		Right: &TreeNode{
			Right: &TreeNode{},
		},
	}
	fmt.Println(widthOfBinaryTree(root))
}

func MathPow(n, m int) int {
	return int(math.Pow(float64(n), float64(m)))
}
