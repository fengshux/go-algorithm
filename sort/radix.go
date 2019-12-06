package sort


func RadixSort(arr []int, d int) {

	divisor := 1
	for k := 1; k <= d ; k++ {

		for i:=0; i<len(arr) -1; i++ {
			for j:=1; j < len(arr) - i; j++ {
				if compare(arr[j-1], arr[j], divisor) > 0 {
					temp := arr[j-1]
					arr[j-1] = arr[j]
					arr[j] = temp
				}
			}
		}
		divisor = divisor*10
	}
	
}

func compare(a, b, divisor int) int {
	if a/divisor%10 > b/divisor%10 {
		return 1
	}

	if a/divisor%10 < b/divisor%10 {
		return -1
	}

	return 0
}
