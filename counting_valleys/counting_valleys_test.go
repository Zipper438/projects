package countingvalleys

import "testing"

func Test_countingValleys(t *testing.T) {
	path := "UDDDUDUU"
	want := int32(1)
	got := countingValleys(path)

	if got != want {
		t.Fatalf("\nexpected: %d\ngot: %d", want, got)
	}

}
