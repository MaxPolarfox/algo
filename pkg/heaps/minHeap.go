package heaps

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(arr []int) {
	parentNode := (len(arr) - 2) / 2
	for currentIdx := parentNode + 1; currentIdx >= 0; currentIdx-- {
		h.siftDown(currentIdx, len(arr)-1)
	}
}

func (h *MinHeap) siftDown(currentIdx, endIndex int) {
	childOneIdx := currentIdx*2 + 1
	for childOneIdx <= endIndex {
		childTwoIdx := -1
		if currentIdx*2+2 <= endIndex {
			childTwoIdx = currentIdx*2 + 2
		}
		idxToSwap := childOneIdx
		if childTwoIdx > -1 && (*h)[childTwoIdx] < (*h)[childOneIdx] {
			idxToSwap = childTwoIdx
		}
		if (*h)[idxToSwap] < (*h)[currentIdx] {
			h.swap(currentIdx, idxToSwap)
			currentIdx = idxToSwap
			childOneIdx = currentIdx*2 + 1
		} else {
			return
		}
	}
}

func (h *MinHeap) siftUp() {
	currIdx := len(*h) - 1
	parentIdx := (currIdx - 1) / 2
	for currIdx > 0 {
		currentVal, parentVal := (*h)[currIdx], (*h)[parentIdx]
		if currentVal < parentVal {
			h.swap(currIdx, parentIdx)
			currIdx = parentIdx
			parentIdx = (currIdx - 1) / 2
		} else {
			return
		}
	}
}

func (h MinHeap) Peek() int {
	return h[0]
}

func (h *MinHeap) Remove() int {
	l := len(*h)
	h.swap(0, l-1)
	removed := (*h)[l-1]
	*h = (*h)[0 : l-1]
	h.siftDown(0, l-2)
	return removed
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
