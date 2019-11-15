package sort

import(
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{4, 6, 1 , 3, 8, 2, 5}
	HeapSort(arr, len(arr))

	for i := 1 ; i < len(arr) ; i ++ {
		if arr[i-1] > arr[i] {
			t.Error("not sorted")
		}
	}
}
