package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)

	for i, num := range nums {
		if index, ok := m[target-num]; ok {
			return []int{i, index}
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 8, 10}
	target := 9
	fmt.Println(twoSum(nums, target))
}
