package srot

import(
	"testing"
)

func TestHeepSort(t *testing.T) {
	ary := HeepSort([]int{4, 6, 1 , 3, 8, 2, 5})

	for i := 1 ; i < len(arr) ; i ++ {
		if arr[i-1] > arr[i] {
			t.Error("not sorted")
		}
	}
}


