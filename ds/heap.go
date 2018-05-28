package ds

// NewHeapBad naively constructs a heap based on the given list of integers
// Worst-case running time: Theta(nlogn)
func NewHeapBad(ns []int) Heap {
	h := &heap{}
	for _, n := range ns {
		h.Insert(n)
	}
	return h
}

// NewHeap constructs a heap using heapify
// Worst-case running time: Theta(n)
func NewHeap(ns []int) Heap {
	bt := bintree(ns)
	h := &heap{bt: bt}

	for i := len(bt) - 1; i >= 0; i-- {
		h.fixDown(i)
	}

	return h
}

// DefaultHeap constructs an empty Heap
func DefaultHeap() Heap {
	return &heap{}
}

// Heap is a binary max-heap with integer keys
type Heap interface {
	Insert(key int)
	DeleteMax() (int, bool)
}

type heap struct {
	bt bintree
}

func (h *heap) Insert(key int) {
	h.bt = append(h.bt, key)

	// fix up
	ci := h.bt.lastIndex()
	for {
		pi, ok := h.bt.parentIndex(ci)
		if !ok {
			break
		}

		if h.bt[pi] >= h.bt[ci] {
			break
		} else {
			h.bt[pi], h.bt[ci] = h.bt[ci], h.bt[pi]
			ci = pi
		}
	}
}

func (h *heap) DeleteMax() (int, bool) {
	if len(h.bt) == 0 {
		return -1, false
	}

	max := h.bt[0]

	// replace root with the last element
	last := h.bt[h.bt.lastIndex()]
	h.bt[0] = last

	// start fix down at root
	h.fixDown(0)

	h.bt = h.bt[:len(h.bt)-1]
	return max, true
}

func (h *heap) fixDown(start int) {
	var pi, ci int
	pi = start
	for {
		li, ok := h.bt.leftChildIndex(pi)

		// if no child, do nothing
		if !ok {
			break
		}

		ri, ok := h.bt.rightChildIndex(pi)

		// locate the largest child
		if !ok {
			ci = li
		} else if h.bt[ri] > h.bt[li] {
			ci = ri
		} else {
			ci = li
		}

		// if parent is larger than the largest child, do nothing
		if h.bt[pi] >= h.bt[ci] {
			break
		}

		h.bt[pi], h.bt[ci] = h.bt[ci], h.bt[pi]
		pi = ci
	}
}

type bintree []int

func (bintree) parentIndex(i int) (int, bool) {
	// if node is root, return -1 (no parent)
	if i == 0 {
		return -1, false
	}

	if i%2 == 0 {
		return (i - 2) / 2, true
	}

	return (i - 1) / 2, true
}

func (bt bintree) leftChildIndex(i int) (int, bool) {
	target := 2*i + 1
	if target < len(bt) {
		return target, true
	}
	return target, false
}

func (bt bintree) rightChildIndex(i int) (int, bool) {
	target := 2*i + 2
	if target < len(bt) {
		return target, true
	}
	return target, false
}

func (bt bintree) lastIndex() int {
	return len(bt) - 1
}
