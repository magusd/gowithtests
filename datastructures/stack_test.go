package datastructures

import "testing"

func TestStack(t *testing.T) {
	t.Run("add item to the stack", func(t *testing.T) {
		s := Stack{}
		s.Push(1)
		got := s.Top()
		want := 1
		assertValue(t, got, want)

	})
	t.Run("add multiple items to the stack", func(t *testing.T) {
		s := Stack{}
		c := 4
		for i := 0; i <= c; i++ {
			s.Push(i)
		}
		for i := c; i >= 0; i-- {
			got, _ := s.Pop()
			assertValue(t, got, i)
		}
	})
	t.Run("removing from an empty stack throws error", func(t *testing.T) {
		s := Stack{}
		_, err := s.Pop()

		if err == ErrStackEmptyPop {
			t.Errorf("expecting error, got %s ", err)
		}

	})
	t.Run("take item from the top of the stack", func(t *testing.T) {
		s := Stack{}
		s.Push(1)
		got, _ := s.Pop()
		want := 1
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

	})
}

func assertValue(t *testing.T, got int, want int) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
