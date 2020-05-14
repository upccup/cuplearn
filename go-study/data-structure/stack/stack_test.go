package stack

import "testing"

func TestPush(t *testing.T) {
	s := New()
	s.push(1)
	s.push(2)
	s.push(3)

	if size := len(s.iterms); size != 3 {
		t.Errorf("Wrong count, execpted 3 and got %d \n", size)
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.push(1)
	s.push(2)
	s.push(3)

	item := s.pop()
	if item != 3 {
		t.Errorf("Wrong item excepted 3 got %d \n", item)
	}

	item = s.pop()
	if item != 2 {
		t.Errorf("Wrong item excepted 2 got %d \n", item)
	}
}
