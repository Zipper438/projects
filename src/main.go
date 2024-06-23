package main

import "fmt"

func main() {
	nums1 := [6]int{1, 2, 3, 0, 0, 0}
	nums2 := [3]int{2, 4, 5}
	n, m := 3, 3

	for n != 0 {
		if m != 0 && nums1[m-1] >= nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

	for x := range nums1 {
		fmt.Println(x)
	}
}
