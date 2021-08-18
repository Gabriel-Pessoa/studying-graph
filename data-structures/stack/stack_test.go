package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	stackTest := NewStack()

	t.Run("starts empty", func(t *testing.T) {
		assert.Equal(t, 0, stackTest.Size())
		assert.Equal(t, true, stackTest.IsEmpty())
	})

	t.Run("pushes elements", func(t *testing.T) {
		stackTest.Push(1)
		assert.Equal(t, 1, stackTest.Size())

		stackTest.Push(2)
		assert.Equal(t, 2, stackTest.Size())

		stackTest.Push(3)
		assert.Equal(t, 3, stackTest.Size())

		assert.Equal(t, false, stackTest.IsEmpty())
	})
}

func TestPop(t *testing.T) {
	stackTest := NewStack()

	stackTest.Push(1)
	stackTest.Push(2)
	stackTest.Push(3)

	t.Run("pops elements", func(t *testing.T) {
		for i := 3; i > 0; i-- {
			value, err := stackTest.Pop()
			assert.Equal(t, nil, err)
			assert.Equal(t, i, value)
		}
	})
}

func TestSize(t *testing.T) {
	stackTest := NewStack()

	t.Run("returns the correct size", func(t *testing.T) {
		assert.Equal(t, 0, stackTest.Size())
		stackTest.Push(1)
		assert.Equal(t, 1, stackTest.Size())
		stackTest.Push(2)
		assert.Equal(t, 2, stackTest.Size())
		stackTest.Push(3)
		assert.Equal(t, 3, stackTest.Size())
		stackTest.Clear()
		assert.Equal(t, true, stackTest.IsEmpty())

		stackTest.Push(1)
		stackTest.Push(2)
		stackTest.Push(3)

		stackTest.Pop()
		assert.Equal(t, 2, stackTest.Size())
		stackTest.Pop()
		assert.Equal(t, 1, stackTest.Size())
		stackTest.Pop()
		assert.Equal(t, 0, stackTest.Size())
		stackTest.Pop()
		assert.Equal(t, 0, stackTest.Size())
	})
}

func TestIsEmpty(t *testing.T) {
	stackTest := NewStack()

	t.Run("returns if it is empty", func(t *testing.T) {
		assert.Equal(t, true, stackTest.IsEmpty())
		stackTest.Push(1)
		assert.Equal(t, false, stackTest.IsEmpty())
		stackTest.Push(2)
		assert.Equal(t, false, stackTest.IsEmpty())
		stackTest.Push(3)
		assert.Equal(t, false, stackTest.IsEmpty())

		stackTest.Clear()
		assert.Equal(t, true, stackTest.IsEmpty())

		stackTest.Push(1)
		stackTest.Push(2)
		stackTest.Push(3)

		stackTest.Pop()
		assert.Equal(t, false, stackTest.IsEmpty())
		stackTest.Pop()
		assert.Equal(t, false, stackTest.IsEmpty())
		stackTest.Pop()
		assert.Equal(t, true, stackTest.IsEmpty())
		stackTest.Pop()
		assert.Equal(t, true, stackTest.IsEmpty())
	})
}

func TestClear(t *testing.T) {
	stackTest := NewStack()

	t.Run("clears the stack", func(t *testing.T) {
		stackTest.Clear()
		assert.Equal(t, true, stackTest.IsEmpty())

		stackTest.Push(1)
		stackTest.Push(2)

		stackTest.Clear()
		assert.Equal(t, true, stackTest.IsEmpty())
	})
}
