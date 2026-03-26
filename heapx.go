// Package heapx: Generic indexed min-heap with Fix/Remove support
// Wraps container/heap with generics, each node tracks its own heap index
// Enables O(log n) priority updates via Fix, no manual heap.Interface boilerplate
//
// heapx: 泛型索引最小堆，支持 Fix/Remove 操作
// 使用泛型封装 container/heap，每个节点自动维护堆中索引
// 通过 Fix 实现 O(log n) 优先级更新，无需手写 heap.Interface 样板代码
package heapx

import (
	"container/heap"

	"github.com/yylego/must"
)

// Node is a heap node that tracks its position in the heap
// V is the value type, must be comparable
//
// Node 是堆节点，自动维护在堆中的位置
// V 是值类型
type Node[V any] struct {
	Value V   // the stored value // 存储的值
	idx   int // position in heap, maintained via swap // 堆中位置，通过 swap 维护
}

// Index returns the current position of this node in the heap
// Returns -1 if the node has been removed
//
// Index 返回节点在堆中的当前位置
// 如果节点已被移除则返回 -1
func (node *Node[V]) Index() int {
	return node.idx
}

// Heap is a generic indexed min-heap
// Less function defines the ordering (return true when a should come before b)
//
// Heap 是泛型索引最小堆
// Less 函数定义排序规则（a 应排在 b 前面时返回 true）
type Heap[V any] struct {
	data innerHeap[V]
}

// New creates a new indexed heap with the given comparison function
//
// New 使用给定的比较函数创建新的索引堆
func New[V any](less func(a, b V) bool) *Heap[V] {
	h := &Heap[V]{
		data: innerHeap[V]{
			less: less,
		},
	}
	heap.Init(&h.data)
	return h
}

// Push adds a value to the heap and returns its node (for later Fix/Remove)
//
// Push 添加值到堆并返回节点（用于后续 Fix/Remove）
func (h *Heap[V]) Push(v V) *Node[V] {
	node := &Node[V]{Value: v}
	heap.Push(&h.data, node)
	return node
}

// Pop removes and returns the minimum node from the heap
// Returns nil if the heap is empty
//
// Pop 移除并返回堆中最小的节点
// 堆为空时返回 nil
func (h *Heap[V]) Pop() *Node[V] {
	if h.data.Len() == 0 {
		return nil
	}
	return heap.Pop(&h.data).(*Node[V])
}

// Peek returns the minimum node without removing it
// Returns nil if the heap is empty
//
// Peek 返回堆中最小的节点但不移除
// 堆为空时返回 nil
func (h *Heap[V]) Peek() *Node[V] {
	if h.data.Len() == 0 {
		return nil
	}
	return h.data.nodes[0]
}

// Fix re-establishes heap ordering after the node's value has been changed
// Call this after modifying node.Value to maintain correct heap order
//
// Fix 在节点值被修改后重新建立堆排序
// 修改 node.Value 后调用此方法保持正确的堆顺序
func (h *Heap[V]) Fix(node *Node[V]) {
	must.True(node.idx >= 0 && node.idx < h.data.Len())
	heap.Fix(&h.data, node.idx)
}

// Remove removes a specific node from the heap
//
// Remove 从堆中移除指定节点
func (h *Heap[V]) Remove(node *Node[V]) {
	must.True(node.idx >= 0 && node.idx < h.data.Len())
	heap.Remove(&h.data, node.idx)
}

// Len returns the number of nodes in the heap
//
// Len 返回堆中节点数量
func (h *Heap[V]) Len() int {
	return h.data.Len()
}

// innerHeap implements heap.Interface for internal use
type innerHeap[V any] struct {
	nodes []*Node[V]
	less  func(a, b V) bool
}

func (h *innerHeap[V]) Len() int { return len(h.nodes) }

func (h *innerHeap[V]) Less(i, j int) bool {
	return h.less(h.nodes[i].Value, h.nodes[j].Value)
}

func (h *innerHeap[V]) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
	h.nodes[i].idx = i
	h.nodes[j].idx = j
}

func (h *innerHeap[V]) Push(x any) {
	node := x.(*Node[V])
	node.idx = len(h.nodes)
	h.nodes = append(h.nodes, node)
}

func (h *innerHeap[V]) Pop() any {
	old := h.nodes
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.idx = -1
	h.nodes = old[:n-1]
	return node
}
