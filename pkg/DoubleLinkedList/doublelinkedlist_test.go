package doublelinkedlist

import (
	"testing"
)

func TestInsertFirst(t *testing.T) {
	list := DoubleLinkedList{}
	list.InsertFirst("A")

	head := list.ShowHead()
	tail := list.ShowTail()

	if head != tail || head != "A" {
		t.Errorf("Expected and Result not equal")
	}

	list.InsertFirst("B")
	list.InsertFirst("C")
	head = list.ShowHead()
	tail = list.ShowTail()

	if head != "C" || tail != "A" {
		t.Errorf("Expected and Result not equal")
	}
}

func TestInsertLast(t *testing.T) {
	list := DoubleLinkedList{}
	list.InsertLast("A")

	head := list.ShowHead()
	tail := list.ShowTail()

	if head != tail || head != "A" {
		t.Errorf("Expected and Result not equal")
	}

	list.InsertLast("B")
	list.InsertLast("C")
	head = list.ShowHead()
	tail = list.ShowTail()

	if head != "A" || tail != "C" {
		t.Errorf("Expected and Result not equal")
	}
}

func TestDeleteFirst(t *testing.T) {
	list := DoubleLinkedList{}
	list.DeleteFirst()

	list.InsertLast("A")
	list.DeleteFirst()
	head := list.ShowHead()
	if head != nil {
		t.Errorf("Expected and Result not equal")
	}

	list.InsertLast("A")
	list.InsertLast("B")
	list.InsertLast("C")
	list.DeleteFirst()

	head = list.ShowHead()
	tail := list.ShowTail()

	if head != "B" || tail != "C" {
		t.Errorf("Expected and Result not equal")
	}
}

func TestDeleteLast(t *testing.T) {
	list := DoubleLinkedList{}
	list.DeleteFirst()

	list.InsertLast("A")
	list.DeleteFirst()
	head := list.ShowHead()
	if head != nil {
		t.Errorf("Expected and Result not equal")
	}

	list.InsertLast("A")
	list.InsertLast("B")
	list.InsertLast("C")
	list.DeleteLast()

	head = list.ShowHead()
	tail := list.ShowTail()

	if head != "A" || tail != "B" {
		t.Errorf("Expected and Result not equal")
	}
}

func TestToArray(t *testing.T) {
	values := []int{1,2,3,4,5,6,7}

	list := DoubleLinkedList{}

	result := list.ToArray()
	if result != nil {
		t.Errorf("Expected nil when empty list")
	}

	for _, value := range values {
		list.InsertLast(value)
	}

	head := list.ShowHead()
	tail := list.ShowTail()
	if head != 1 || tail != 7 {
		t.Errorf("Expected and Result not equal")
	}


	result = list.ToArray()
	for i, value := range values {
		if result[i].(int) != value {
			t.Errorf("Value expected not equal real value")
		}
	}
}