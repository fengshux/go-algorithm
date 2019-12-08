package mos


func MinAndMax( arr []int ) (min , max int) {

	var start int
	
	if len(arr)%2 == 0 {
		start = 2

		if compare(arr[0], arr[1]) > 0 {
			min = arr[1]
			max = arr[0]
		} else {
			min = arr[0]
			max = arr[1]
		}
		
	} else {
		min = arr[0]
		max = arr[0]
	}

	for i := start ; i < len(arr); i = i + 2 {

		if compare( arr[i], arr[i+1] ) > 0 {
			if compare(arr[i], max) > 0 {
				max = arr[i]
			}
			if compare(arr[i+1],  min) < 0 {
				min = arr[i+1]
			}			
			
		} else {
			if compare(arr[i+1], max) > 0 {
				max = arr[i+1]
			}
			if compare(arr[i],  min) < 0 {
				min = arr[i]
			}
		}
		
	}
	return	
}


func compare(a, b int) int {
	if a > b {
		return 1
	}

	if a < b {
		return -1
	}

	return 0
}
