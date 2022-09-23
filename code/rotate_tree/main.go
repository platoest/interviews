package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	value int
}

func rotate(node *Node) {
	node.left, node.right = node.right, node.left
	if node.left != nil {
		rotate(node.left)
	}
	if node.right != nil {
		rotate(node.right)
	}
}

func printTree(root *Node, index int) {
	fmt.Println(root.value)
	if root.right != nil {
		for i := 0; i < index; i++ {
			fmt.Print("│       ")
		}
		if root.left == nil {
			fmt.Print("└── right: ")
		} else {
			fmt.Print("├── right: ")
		}
		printTree(root.right, index+1)
	}
	if root.left != nil {
		for i := 0; i < index; i++ {
			fmt.Print("│       ")
		}
		fmt.Print("└── left: ")
		printTree(root.left, index+1)
	}
}

func main() {
	root := &Node{
		left: &Node{
			right: &Node{
				value: 6,
			},
			value: 2,
		},
		right: &Node{
			left: &Node{
				value: 4,
			},
			right: &Node{
				value: 5,
			},
			value: 3,
		},
		value: 1,
	}
	printTree(root, 0)
	rotate(root)
	printTree(root, 0)
}
