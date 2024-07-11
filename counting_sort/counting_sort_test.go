package countingsort

import (
	"testing"
)

func Test_countingSort(t *testing.T) {
	arr := []int32{1, 1, 3, 2, 1}
	want := []int32{1, 1, 1, 2, 3}
	got := countingSort(arr)

	t.Fatalf("\nwant: %v\ngot:  %v", want, got)

}
