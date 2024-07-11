package mock

func max(i, j int32) int32 {
	if i > j {
		return i
	}
	return j
}

func flippingMatrix(matrix [][]int32) int32 {
	n := len(matrix) / 2
	var maxSum int32

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			maxValue := max(
				max(matrix[i][j], matrix[i][2*n-j-1]),
				max(matrix[2*n-i-1][j], matrix[2*n-i-1][2*n-j-1]),
			)
			maxSum += maxValue
		}
	}
	return maxSum
}
