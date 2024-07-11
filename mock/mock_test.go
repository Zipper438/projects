package mock

import "testing"

func Test_mock(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int32
		want   int32
	}{
		{
			name: "N1",
			matrix: [][]int32{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: 54,
		},
		{
			name: "N2",
			matrix: [][]int32{
				{107, 54, 128, 15},
				{12, 75, 110, 138},
				{100, 96, 34, 85},
				{75, 15, 28, 112},
			},
			want: 488,
		},
		{
			name: "N3",
			matrix: [][]int32{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{9, 10, 11, 12, 13, 14, 15, 16},
				{17, 18, 19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30, 31, 32},
				{33, 34, 35, 36, 37, 38, 39, 40},
				{41, 42, 43, 44, 45, 46, 47, 48},
				{49, 50, 51, 52, 53, 54, 55, 56},
				{57, 58, 59, 60, 61, 62, 63, 64},
			},
			want: 808,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flippingMatrix(tt.matrix); got != tt.want {
				t.Fatalf("\nExpected: %d\nGot: %d", tt.want, got)
			}
		})
	}
}
