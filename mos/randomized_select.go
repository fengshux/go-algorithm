package mos

func RandomizedSelect(arr []int, p,r, i int) int {
	if p == i {
		return arr[p]
	}
	
	q := RandomizedPartition(arr,p, r)

	k := q-p+1
	if i == k {
		return arr[q]
	} else if i < k {
		return RandomizedSelect(arr,p,q-1, i)
	} else {
		return RandomizedSelect(arr, q+1, r, i-k)
	}
}
