package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
	fmt.Printf("%v", nums)
}
