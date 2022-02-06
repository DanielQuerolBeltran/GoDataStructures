package queue

import "testing"

func TestQueue(t *testing.T) {
	values := []int{1,2,3,4,6,7,8,9,10}

	queue := Queue{}

	empty := queue.IsEmpty()
	if empty != true {
		t.Errorf("Empty operation error")
	}


	value := queue.Peek()
	if value != nil {
		t.Errorf("Not nil value when queue is empty")
	}

	for _, value := range values {
		queue.Enqueue(value)
	}

	for i, value := range values {
		if queue[i] != value {
			t.Errorf("Value expected not equal real value")
		}
	}

	value = queue.Peek()
	if value != 1 || len(values) != len(queue) {
		t.Errorf("Peek operation failure")
	}

	value = queue.Dequeue()
	if value != 1 || len(values)-1 != len(queue) {
		t.Errorf("Peek operation failure")
	}

	for i := 0; i < len(queue); i++{
		queue.Dequeue()
	}

	empty = queue.IsEmpty()
	if empty != false {
		t.Errorf("Empty operation error")
	}
}