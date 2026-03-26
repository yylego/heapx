package heapx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func popValue[V any](t *testing.T, h *Heap[V]) V {
	node := h.Pop()
	require.NotNil(t, node)
	return node.Value
}

func peekValue[V any](t *testing.T, h *Heap[V]) V {
	node := h.Peek()
	require.NotNil(t, node)
	return node.Value
}

func TestHeap_PushPop(t *testing.T) {
	h := New[int](func(a, b int) bool { return a < b })

	h.Push(3)
	h.Push(1)
	h.Push(2)

	require.Equal(t, 3, h.Len())
	require.Equal(t, 1, peekValue(t, h))

	require.Equal(t, 1, popValue(t, h))
	require.Equal(t, 2, popValue(t, h))
	require.Equal(t, 3, popValue(t, h))
	require.Nil(t, h.Pop())
}

func TestHeap_Fix(t *testing.T) {
	h := New[int](func(a, b int) bool { return a < b })

	n1 := h.Push(10)
	_ = h.Push(20)
	n3 := h.Push(30)

	require.Equal(t, 10, peekValue(t, h))

	n3.Value = 1
	h.Fix(n3)
	require.Equal(t, 1, peekValue(t, h))

	n1.Value = 100
	h.Fix(n1)

	require.Equal(t, 1, popValue(t, h))
	require.Equal(t, 20, popValue(t, h))
	require.Equal(t, 100, popValue(t, h))
}

func TestHeap_Remove(t *testing.T) {
	h := New[int](func(a, b int) bool { return a < b })

	n1 := h.Push(10)
	h.Push(20)
	h.Push(30)

	h.Remove(n1)
	require.Equal(t, 2, h.Len())
	require.Equal(t, 20, peekValue(t, h))
}

func TestHeap_EmptyPeek(t *testing.T) {
	h := New[int](func(a, b int) bool { return a < b })
	require.Nil(t, h.Peek())
	require.Nil(t, h.Pop())
}

type task struct {
	name     string
	deadline int64
}

func TestHeap_Struct(t *testing.T) {
	h := New[task](func(a, b task) bool { return a.deadline < b.deadline })

	h.Push(task{name: "c", deadline: 300})
	h.Push(task{name: "a", deadline: 100})
	h.Push(task{name: "b", deadline: 200})

	require.Equal(t, "a", popValue(t, h).name)
	require.Equal(t, "b", popValue(t, h).name)
	require.Equal(t, "c", popValue(t, h).name)
}

func TestHeap_FixWithStruct(t *testing.T) {
	h := New[task](func(a, b task) bool { return a.deadline < b.deadline })

	n1 := h.Push(task{name: "a", deadline: 100})
	h.Push(task{name: "b", deadline: 200})

	n1.Value.deadline = 999
	h.Fix(n1)

	require.Equal(t, "b", peekValue(t, h).name)
}

func TestNode_Index(t *testing.T) {
	h := New[int](func(a, b int) bool { return a < b })

	n := h.Push(10)
	require.True(t, n.Index() >= 0)

	h.Remove(n)
	require.Equal(t, -1, n.Index())
}
