package main

import "fmt"

func romanToInt(s string) int {
	dic := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var res int
	for i := 0; i < len(s); {
		if i != len(s)-1 && dic[s[i]] < dic[s[i+1]] {
			res += dic[s[i+1]] - dic[s[i]]
			i += 2
		} else {
			res += dic[s[i]]
			i++
		}
	}
	return res
}

func main() {
	s := "IX"
	fmt.Print(romanToInt(s))
}
