package chess

import "testing"

func TestFileOrder(t *testing.T) {
	t.Parallel()

	if FileA >= FileB {
		t.Fatalf("FileA=%d >= FileB=%d", FileA, FileB)
	}
}
