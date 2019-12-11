package mos

import(
	"math/rand"
)

func exchange(arr [] int , i , j int) {
	temp := arr[j]
	arr[j] = arr[i]
	arr[i] = temp
}

func Partition(arr []int, p, r int) int {

	i := p - 1

	for j := p ; j < r; j++ {
		if arr[j] < arr[r] {
			i = i+1
			exchange(arr, i,j)
		}
	}

	i = i+1
	exchange(arr, i, r)
	return i
}


func RandomizedPartition(arr []int, p, r int) int {
	if r == p {
		return p
	}	
	n := rand.Intn(r-p) + p
	exchange(arr, n, r)
	return Partition(arr, p, r)
}
