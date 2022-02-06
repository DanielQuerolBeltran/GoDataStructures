package linkedlist

import (
	"testing"
)


func TestInsertLast(t *testing.T) {
	var expected []interface{}
	expected = append(expected, 5, 4, 3, 2, 1)

	ll := LinkedList{}
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)
	ll.InsertLast(4)
	ll.InsertLast(5)
	result := ll.ToArray()

	if equal(expected, result) {
       t.Errorf("Expected and Result not equal")
    }
}

func TestInsertLastWhenOnlyOneNode(t *testing.T) {
	var expected []interface{}
	expected = append(expected, 5)

	ll := LinkedList{}
	ll.InsertLast(5)
	result := ll.ToArray()

	if equal(expected, result) {
       t.Errorf("Expected and Result not equal")
    }
}

func TestInsertLastWhenEmptyList(t *testing.T) {
	ll := LinkedList{}
	ll.Display()
	result := ll.ToArray()


	if len(result) == 0 {
       t.Errorf("Expected and Result not equal")
    }
}

func equal(a, b []interface{}) bool {
    if len(a) != len(b) {
        return false
    }

    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}