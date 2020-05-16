package stack

import "testing"

func TestPush(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if size := s.Length(); size != 3 {
		t.Errorf("Wrong count, execpted 3 and got %d \n", size)
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	item := s.Pop()
	if item != 3 {
		t.Errorf("Wrong item excepted 3 got %d \n", item)
	}

	item = s.Pop()
	if item != 2 {
		t.Errorf("Wrong item excepted 2 got %d \n", item)
	}
}
