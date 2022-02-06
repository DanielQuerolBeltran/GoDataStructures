package linkedlist

import (
	"fmt"
)

type Node struct {
	next *Node
	content interface{}
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertLast(content interface{}) {
	node := Node{
		next: nil,
		content: content,
	}

	if l.head == nil {
		l.head = &node
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
	
		currentNode.next = &node
	}
}

func (l *LinkedList) InsertFirst(content interface{}) {
	node := Node{
		next: nil,
		content: content,
	}

	node.next = l.head
	l.head = &node
}

func (l *LinkedList) DeleteFirst() {
	if l.head != nil {
		if l.head.next == nil {
			l.head = nil
		} else {
			l.head = l.head.next
		}
	}
}

func (l *LinkedList) DeleteLast() {
	if l.head != nil {
		if l.head.next == nil {
			l.head = nil
		} else {
			currentNode := l.head
			for currentNode.next.next != nil {
				currentNode = currentNode.next
			}
			currentNode.next = nil
		}
	}
}

func (l *LinkedList) Sort() {
	if l.head != nil && l.head.next != nil {
		var anyChanges bool = true;
		for anyChanges {
			anyChanges = false
			var last *Node
			prev := l.head
			current := l.head.next
			for current != nil {
				if prev.toStringContent() > current.toStringContent() {
					if last == nil { 
						l.head = current
					} else {
						last.next = current
					}
	
					saved := current.next
					current.next = prev
					if current.next != nil {
						prev.next = saved
					} else {
						prev.next = nil
					}
					anyChanges = true
				}
	
				last = prev
				prev = current
				current = current.next
			}
		}
	}
}

func (n *Node) toStringContent() string {
	return fmt.Sprintf("%v", n.content)
}

func (l *LinkedList) Display() {
	if l.head != nil {
		currentNode := l.head
		for currentNode != nil {
			fmt.Printf(" %+v ->", currentNode.content)
			currentNode = currentNode.next 
		}
	}
	fmt.Println()

}

func (l *LinkedList) ToArray() []interface{} {
	var arrayContent []interface{}

	if l.head != nil {
		currentNode := l.head
		for currentNode != nil {
			arrayContent = append(arrayContent, currentNode.toStringContent())
			currentNode = currentNode.next
		}
	}

	return arrayContent
}