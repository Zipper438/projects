package countingvalleys

func countingValleys(path string) int32 {
	//runes := []rune(path)
	var count int32
	lvl := 0
	for _, step := range path {
		switch step {
		case 'U':
			lvl++
			if lvl == 0 {
				count++
			}
		case 'D':
			lvl--
		}
	}
	return count
}
