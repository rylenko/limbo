package chess

import "testing"

func TestFileOrder(t *testing.T) {
	t.Parallel()

	fileANumber := uint8(FileA)
	fileBNumber := uint8(FileB)

	if fileANumber >= fileBNumber {
		t.Fatalf("FileA=%d must be less than FileB=%d", fileANumber, fileBNumber)
	}
}
