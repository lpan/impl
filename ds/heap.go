package ds

// NewHeap constructs an empty Heap
func NewHeap() Heap {
	return &heap{}
}

// Heap is a binary max-heap with integer keys
type Heap interface {
	Insert(key int)
	DeleteMax() (int, bool)
}

type heap struct {
	slice []int
}

func (h *heap) Insert(key int) {
	h.slice = append(h.slice, key)

	// fix up
	ci := h.lastIndex()
	for {
		pi, ok := h.parentIndex(ci)
		if !ok {
			break
		}

		if h.slice[pi] >= h.slice[ci] {
			break
		} else {
			h.slice[pi], h.slice[ci] = h.slice[ci], h.slice[pi]
			ci = pi
		}
	}
}

func (h *heap) DeleteMax() (int, bool) {
	if len(h.slice) == 0 {
		return -1, false
	}

	max := h.slice[0]

	// replace root with the last element
	last := h.slice[h.lastIndex()]
	h.slice[0] = last

	// fix down
	var pi, ci int
	for {
		li, ok := h.leftChildIndex(pi)
		if !ok {
			break
		}

		ri, ok := h.rightChildIndex(pi)
		if !ok {
			ci = li
		} else if h.slice[ri] > h.slice[li] {
			ci = ri
		} else {
			ci = li
		}

		h.slice[pi], h.slice[ci] = h.slice[ci], h.slice[pi]
		pi = ci
	}

	h.slice = h.slice[:len(h.slice)-1]
	return max, true
}

func (h *heap) parentIndex(i int) (int, bool) {
	// if node is root, return -1 (no parent)
	if i == 0 {
		return -1, false
	}

	if i%2 == 0 {
		return (i - 2) / 2, true
	}

	return (i - 1) / 2, true
}

func (h *heap) leftChildIndex(i int) (int, bool) {
	target := 2*i + 1
	if target < len(h.slice) {
		return target, true
	}
	return target, false
}

func (h *heap) rightChildIndex(i int) (int, bool) {
	target := 2*i + 2
	if target < len(h.slice) {
		return target, true
	}
	return target, false
}

func (h *heap) lastIndex() int {
	return len(h.slice) - 1
}
