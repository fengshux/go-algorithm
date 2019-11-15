package sort


type MaxPriorityQueue struct {
	Heap []int
	Size int 
}

func NewMaxPriorityQueue() *MaxPriorityQueue {
	return &MaxPriorityQueue{make([]int, 100),0}
}

func (q *MaxPriorityQueue) MaxMum() int {
	if q.Size == 0 {
		panic("Queue is empty")
	}
	return q.Heap[0]
}

func (q *MaxPriorityQueue) ExtractMaxMum() int {
	if q.Size == 0 {
		panic("Queue is empty")
	}
	max := q.Heap[0]
	q.Heap[0] = q.Heap[q.Size-1]
	q.Size = q.Size - 1
	maxHeapify(q.Heap, 0, q.Size)
	return max
}

func (q *MaxPriorityQueue) Increase(i, val int) {
	if val < q.Heap[i] {
		panic("New value is smaller then current value.")
	}

	q.Heap[i] = val
	for i > 0 && q.Heap[parent(i)] < q.Heap[i] {
		exchange(q.Heap, parent(i), i)
		i = parent(i)
	}
}

func (q *MaxPriorityQueue) Insert(val int) {
	if q.Size == len(q.Heap) {
		panic("Queue is full")
	}
	q.Size = q.Size + 1
	q.Heap[q.Size -1 ] = val -1
	q.Increase(q.Size -1, val)
}
