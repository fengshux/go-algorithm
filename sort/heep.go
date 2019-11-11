package srot

import (
	"fmt"
)

func HeapSort(ary []int, length int) {
	
	for i := length/2; i > 0; i--  {
		justHeap(ary, i, length)		
	}
	for i := length-1; i > 0; i-- {
		swap(ary, 0, i)
		justHeap(ary, 0, i)
	}
}

func justHeap (ary []int, i, length int) {	
	k := 2*i+1
	if k + 1 < length {
		max := k+1
		if ary[k+1] < ary[k] {
			max = k
		}
		if ary[i] < ary[max] {
			swap(ary, i , max)
			justHeap(ary, max, length)
		}
	}

	if k < length {
		if ary[i] < ary[k] {
			swap(ary, i, k)
			justHeap(ary, k, length)
		}
	}
	
}

func swap( ary []int, a, b int ) {
	temp := ary[a]
	ary[a] = ary[b]
	ary[b] = temp
}
