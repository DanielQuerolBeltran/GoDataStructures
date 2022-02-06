package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	values := []int{1,2,3,4,6,7,8,9,10}

	stack := Stack{}
	value := stack.Pop()

	if value != nil {
		t.Errorf("Not nil when empty stack")
	}

	empty := stack.IsEmpty()

	if empty != true {
		t.Errorf("Empty operation error")
	}

	for _, value := range values {
		stack.Push(value)
	}

	if len(values) != len(stack) {
		t.Errorf("Length is not equal")
	}

	empty = stack.IsEmpty()

	if empty != false {
		t.Errorf("Empty operation error")
	}

	for j := len(values)-1; j >= 0; j-- {
		value = stack.Pop()
		if value != values[j] {
			t.Errorf("Pop operation error. Value expected not equal to real value")
		}
	}
	
	value = stack.Pop()

	if value != nil {
		t.Errorf("Not nil when empty stack")
	}
}