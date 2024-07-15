package stack_test

import (
	"testing"

	"github.com/anirudhology/leetcode-go/problems/stack"
)

func TestMinStack(t *testing.T) {
	t.Run("Push and Top", func(t *testing.T) {
		minStack := stack.Constructor()
		minStack.Push(1)
		if got := minStack.Top(); got != 1 {
			t.Errorf("Top() = %v, want %v", got, 1)
		}

		minStack.Push(2)
		if got := minStack.Top(); got != 2 {
			t.Errorf("Top() = %v, want %v", got, 2)
		}

		minStack.Push(-1)
		if got := minStack.Top(); got != -1 {
			t.Errorf("Top() = %v, want %v", got, -1)
		}
	})

	t.Run("Pop", func(t *testing.T) {
		minStack := stack.Constructor()
		minStack.Push(1)
		minStack.Push(2)
		minStack.Push(3)

		minStack.Pop()
		if got := minStack.Top(); got != 2 {
			t.Errorf("Top() = %v, want %v", got, 2)
		}

		minStack.Pop()
		if got := minStack.Top(); got != 1 {
			t.Errorf("Top() = %v, want %v", got, 1)
		}
	})

	t.Run("GetMin", func(t *testing.T) {
		minStack := stack.Constructor()
		minStack.Push(1)
		if got := minStack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want %v", got, 1)
		}

		minStack.Push(2)
		if got := minStack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want %v", got, 1)
		}

		minStack.Push(-1)
		if got := minStack.GetMin(); got != -1 {
			t.Errorf("GetMin() = %v, want %v", got, -1)
		}

		minStack.Pop()
		if got := minStack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want %v", got, 1)
		}
	})

	t.Run("Mixed Operations", func(t *testing.T) {
		minStack := stack.Constructor()
		minStack.Push(3)
		minStack.Push(5)
		if got := minStack.GetMin(); got != 3 {
			t.Errorf("GetMin() = %v, want %v", got, 3)
		}

		minStack.Push(2)
		minStack.Push(1)
		if got := minStack.GetMin(); got != 1 {
			t.Errorf("GetMin() = %v, want %v", got, 1)
		}

		minStack.Pop()
		if got := minStack.GetMin(); got != 2 {
			t.Errorf("GetMin() = %v, want %v", got, 2)
		}

		minStack.Pop()
		if got := minStack.GetMin(); got != 3 {
			t.Errorf("GetMin() = %v, want %v", got, 3)
		}

		if got := minStack.Top(); got != 5 {
			t.Errorf("Top() = %v, want %v", got, 5)
		}
	})
}
