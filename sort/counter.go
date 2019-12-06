package sort


func CounterSort( arr []int, max int)  []int{

	counter := make([]int, max+1)
	sorted := make([]int, len(arr))

	for i := range arr {
		counter[arr[i]] = counter[arr[i]] + 1
	}

	for i:=1; i <= max; i++ {
		counter[i] = counter[i] + counter[i-1]
	}

	for i := range arr {
		item := arr[i]
		// 下标是从0开始，所以要counter[item]-1
		sorted[counter[item]-1] = item
		counter[item] = counter[item] - 1
	}
	return sorted
}
