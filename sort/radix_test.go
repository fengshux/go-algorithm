package sort

import(
	"testing"
	"fmt"
)

func TestRadixSort(t *testing.T) {
	arr := []int{228, 306, 444, 103, 429}

	RadixSort(arr, 3)
	fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Error("Not sorted.")
		}
	}
}
