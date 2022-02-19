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

	tail := list.ShowTail()

	if tail != nil {
		t.Errorf("Expected tail equals nil")
	}

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
	tail = list.ShowTail()

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


func TestSortIntValues(t *testing.T) {
	initialVelues := []int{8,4,6,9,1,0}
	expectedValues := []int{0,1,4,6,8,9}

	list := DoubleLinkedList{}

	for _, value := range initialVelues {
		list.InsertLast(value)
	}

	if list.ShowHead().(int) != 8 || list.ShowTail().(int) != 0 {
		t.Errorf("Initial venues not loaded correctly")
	}

	list.Sort()
	result := list.ToArray()
	for i, value := range expectedValues {
		if result[i].(int) != value {
			t.Errorf("Value expected not equal real value")
		}
	}
}

func TestSortFloatValues(t *testing.T) {
	initialVelues := []float32{8.4,4.1,6.003,6.001,0,0.01}
	expectedValues := []float32{0,0.01,4.1,6.001, 6.003,8.4}

	list := DoubleLinkedList{}

	for _, value := range initialVelues {
		list.InsertLast(value)
	}

	if list.ShowHead().(float32) != 8.4 || list.ShowTail().(float32) != 0.01 {
		t.Errorf("Initial venues not loaded correctly")
	}

	list.Sort()
	result := list.ToArray()
	for i, value := range expectedValues {
		if result[i].(float32) != value {
			t.Errorf("Value expected not equal real value")
		}
	}
}

func TestSortStringValues(t *testing.T) {
	initialVelues := []string{"Hector", "Antonio", "Daniel", "Manolo"}
	expectedValues := []string{"Antonio", "Daniel", "Hector", "Manolo"}

	list := DoubleLinkedList{}

	for _, value := range initialVelues {
		list.InsertLast(value)
	}

	if list.ShowHead().(string) != "Hector" || list.ShowTail().(string) != "Manolo" {
		t.Errorf("Initial venues not loaded correctly")
	}

	list.Sort()
	result := list.ToArray()
	for i, value := range expectedValues {
		if result[i].(string) != value {
			t.Errorf("Value expected not equal real value")
		}
	}
}


func TestErrorInsertMixedTypes(t *testing.T) {
	list := DoubleLinkedList{}

	err := list.InsertLast("Daniel")

	if err != nil {
		t.Errorf("Not expected error inserting value")
	}

	err = list.InsertFirst(10)
	if err == nil {
		t.Errorf("Expected error inserting value with different type")
	}

	err = list.InsertLast(10.4)
	if err == nil {
		t.Errorf("Expected error inserting value with different type")
	}

	err = list.InsertFirst(true)
	if err == nil {
		t.Errorf("Expected error inserting value with different type")
	}
}