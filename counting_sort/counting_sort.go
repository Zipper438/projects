package countingsort

func countingSort(arr []int32) []int32 {
	max := arr[0]
	for i := range arr {
		if arr[i] > max {
			max = arr[i]
		}
	}
	count := make([]int, max+1)
	for _, num := range arr {
		count[num]++
	}
	j := 0
	for i := range count {
		for count[i] > 0 {
			arr[j] = int32(i)
			j++
			count[i]--
		}
	}
	return arr
}
