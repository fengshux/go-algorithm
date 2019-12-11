package mos

import(
	"testing"
)


func TestRandomizedSelect(t *testing.T) {
	arr := []int{3,4,2,1,7,8,5,6,9}

	m := RandomizedSelect(arr, 0, len(arr) -1 , 5)
	
	if m != 5 {
		t.Error("Error", m)
	}
}
