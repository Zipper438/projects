package flipbeats

import "testing"

func Test_flip(t *testing.T) {
	tests := []struct {
		name string
		num  int64
		want uint32
	}{
		{
			name: "1",
			num:  1,
			want: 4294967294,
		},
		{
			name: "2147483647",
			num:  2147483647,
			want: 2147483648,
		},
		{
			name: "0",
			num:  0,
			want: 4294967295,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plippingBits(tt.num); got != tt.want {
				t.Fatalf("\nExpected: %d\nGot: %d", tt.want, got)
			}
		})
	}
}
