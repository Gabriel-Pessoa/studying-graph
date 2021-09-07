package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushStack(t *testing.T) {
	stackTest := NewStack()

	assert.Equal(t, 0, stackTest.Size())
	assert.Equal(t, true, stackTest.IsEmpty())

	stackTest.Push(1)
	assert.Equal(t, 1, stackTest.Size())

	stackTest.Push(2)
	assert.Equal(t, 2, stackTest.Size())

	stackTest.Push(3)
	assert.Equal(t, 3, stackTest.Size())

	assert.Equal(t, false, stackTest.IsEmpty())
}

func TestPopStack(t *testing.T) {
	stackTest := NewStack()

	stackTest.Push(1)
	stackTest.Push(2)
	stackTest.Push(3)

	for i := 3; i > 0; i-- {
		value, err := stackTest.Pop()
		assert.Equal(t, nil, err)
		assert.Equal(t, i, value)
	}
}

func TestSizeStack(t *testing.T) {
	stackTest := NewStack()

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
}

func TestIsEmptyStack(t *testing.T) {
	stackTest := NewStack()

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
}

func TestClearStack(t *testing.T) {
	stackTest := NewStack()

	stackTest.Clear()
	assert.Equal(t, true, stackTest.IsEmpty())

	stackTest.Push(1)
	stackTest.Push(2)

	stackTest.Clear()
	assert.Equal(t, true, stackTest.IsEmpty())
}
