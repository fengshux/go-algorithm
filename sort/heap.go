package sort

import (
)

func HeapSort(ary []int, length int) {

	buildMaxHeap(ary, length)
	for i := length-1; i > 0; i-- {
		exchange(ary, 0, i)
		maxHeapify(ary, 0, i)
	}
}

func buildMaxHeap(ary []int, length int) {
	for i := parent(length -1 ); i >= 0; i-- {
		maxHeapify(ary, i, length)
	}
}

func parent(i int) int {
	return (i-1)/2
}

func leftChaild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}

func maxHeapify (ary []int, i, length int) {
	k := leftChaild(i)
	if k + 1 < length {
		max := k+1
		if ary[k+1] < ary[k] {
			max = k
		}
		if ary[i] < ary[max] {
			exchange(ary, i , max)
			maxHeapify(ary, max, length)
		}
	}

	if k < length {
		if ary[i] < ary[k] {
			exchange(ary, i, k)
			maxHeapify(ary, k, length)
		}
	}
	
}

func exchange( ary []int, a, b int ) {
	temp := ary[a]
	ary[a] = ary[b]
	ary[b] = temp
}
 
