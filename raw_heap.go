package heapx

// rawHeap implements heap.Interface for internal use
type rawHeap[V any] struct {
	nodes []*Node[V]
	vLess func(a, b V) bool
}

func newRawHeap[V any](vLess func(a, b V) bool) *rawHeap[V] {
	return &rawHeap[V]{
		nodes: make([]*Node[V], 0),
		vLess: vLess,
	}
}

func (h *rawHeap[V]) Len() int { return len(h.nodes) }

func (h *rawHeap[V]) Less(i, j int) bool {
	return h.vLess(h.nodes[i].Value, h.nodes[j].Value)
}

func (h *rawHeap[V]) Swap(i, j int) {
	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
	h.nodes[i].idx = i
	h.nodes[j].idx = j
}

func (h *rawHeap[V]) Push(x any) {
	node := x.(*Node[V])
	node.idx = len(h.nodes)
	h.nodes = append(h.nodes, node)
}

func (h *rawHeap[V]) Pop() any {
	old := h.nodes
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.idx = -1
	h.nodes = old[:n-1]
	return node
}
