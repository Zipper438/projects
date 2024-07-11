package simple

import "testing"

func Test_sum(t *testing.T) {
	a, b := 3, 4
	want := 7
	got := sum(a, b)

	if got != want {
		t.Fatalf("ожидалось: %d; получено: %d", want, got)
	}
	t.Log("OK")
}
