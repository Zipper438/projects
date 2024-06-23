package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	myMap := map[int]struct{}{}

	for _, n := range nums1 {
		myMap[n] = struct{}{}
	}

	var res []int
	for _, num := range nums2 {
		if _, ok := myMap[num]; ok {
			res = append(res, num)
			delete(myMap, num)
		}
	}
	return res
}

func main() {
	n1 := []int{1, 2, 3}
	n2 := []int{2, 3, 4}
	fmt.Println(intersection(n1, n2))
}
