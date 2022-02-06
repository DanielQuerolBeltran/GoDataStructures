package binarySearchTree

import "fmt"

type Node struct {
	right *Node
	left *Node
	content interface{}
}

type BinarySearchTree struct {
	root *Node
}

func (t *BinarySearchTree) Insert(content interface{}) {
	node := Node{
		right: nil,
		left: nil,
		content: content,
	}

	if t.root == nil {
		t.root = &node
	} else {
		t.insertNode(&node, t.root)
	}
}

func (t *BinarySearchTree) insertNode(newNode *Node, currentNode *Node) {
	if (newNode.toStringContent() < currentNode.toStringContent()) {
		if currentNode.left == nil {
			currentNode.left = newNode
		} else {
			t.insertNode(newNode, currentNode.left)
		}
	} else {
		if currentNode.right == nil {
			currentNode.right = newNode
		} else {
			t.insertNode(newNode, currentNode.right)
		}
	}
}

func (n *Node) toStringContent() string {
	return fmt.Sprintf("%v", n.content)
}


func (t *BinarySearchTree) TraverseLeftTest() {
	if t.root != nil {
		currentNode := t.root
		for currentNode.left != nil {
			fmt.Printf(" %+v ->", currentNode.content)
			currentNode = currentNode.left 
		}
		fmt.Printf(" %+v", currentNode.content)
	}
	fmt.Println()
}


func (t *BinarySearchTree) PreorderTraversal() {

}
