package mos


import (
	"testing"
)


func TestMinAndMax( t *testing.T) {

	arr := []int{3,7,8,2,6,9,1,5}
	min, max := MinAndMax(arr)

	if min != 1 || max != 9 {
		t.Error("worn")
	}
}
