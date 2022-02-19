package doublelinkedlist

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrInvalidDate = errors.New("invalid content type")


type Node struct {
	prev *Node
	next *Node
	content interface{}
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (l *DoubleLinkedList) InsertFirst(content interface{}) (error) {
	if l.head != nil {
		err := l.checkContentType(content)

		if err != nil {
			return err
		}
	}

	node := Node{
		next: l.head,
		prev: nil,
		content: content,
	}

	if l.head != nil {
		l.head.prev = &node
	}

	l.head = &node
	if l.tail == nil {
		l.tail = &node
	}

	return nil
}

func (l *DoubleLinkedList) checkContentType(content interface{}) error {
	typeNewContent := reflect.TypeOf(content)
		typeFirstContent := reflect.TypeOf(l.head.content)

		if (typeFirstContent != typeNewContent) {
			return ErrInvalidDate
		}

		return nil
}

func (l *DoubleLinkedList) ShowHead() interface{} {
	if (l.head == nil) {
		return nil
	}
	return l.head.content
}

func (l *DoubleLinkedList) ShowTail()  interface{} {
	if (l.head == nil) {
		return nil
	}

	return l.tail.content
}

func (l *DoubleLinkedList) InsertLast(content interface{}) error {
	if l.head != nil {
		err := l.checkContentType(content)

		if err != nil {
			return err
		}
	}

	node := Node{
		next: nil,
		prev: l.tail,
		content: content,
	}

	if l.tail != nil {
		l.tail.next = &node
	}

	l.tail = &node
	if l.head == nil {
		l.head = &node
	}

	return nil
}

func (l *DoubleLinkedList) DeleteFirst() {
	if l.head != nil {
		if l.head.next != nil {
			l.head.next.prev = nil
		}
		l.head = l.head.next
	}
}

func (l *DoubleLinkedList) DeleteLast() {
	if l.tail != nil {
		if l.tail.prev != nil {
			l.tail.prev.next = nil
		}
		l.tail = l.tail.prev
	}
}

func (l *DoubleLinkedList) Sort() {
	if l.head != l.tail {
		var anyChanges bool = true;
		for anyChanges {
			anyChanges = false
			current := l.head
			for current != nil && current.next != nil {
				if current.toStringContent() > current.next.toStringContent() {
					anyChanges = true
					savedNext := current.next
					savedPrev := current.prev
					if current.next == l.tail {
						current.next = nil
						current.prev = savedNext
						l.tail = current
					} else {
						current.prev = current.next
						current.next = current.next.next
						current.next.prev = current
					}
	
					if current == l.head {
						savedNext.prev = nil
						savedNext.next = current
						l.head = savedNext
					} else {
						savedPrev.next = savedNext
						savedNext.next = current
						savedNext.prev = savedPrev
					}
				}
	
				current = current.next
			}
		}
	}
}

func (n *Node) toStringContent() string {
	return fmt.Sprintf("%v", n.content)
}

func (l *DoubleLinkedList) ToArray() []interface{} {
	var arrayContent []interface{}
	
	if l.head != nil {
		currentNode := l.head
		for currentNode != nil {
			arrayContent = append(arrayContent, currentNode.content)
			currentNode = currentNode.next
		}
	}

	return arrayContent
}