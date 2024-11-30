package chess

import "testing"

func TestRankOrder(t *testing.T) {
	t.Parallel()

	rank1Number := uint8(Rank1)
	rank2Number := uint8(Rank2)

	if rank1Number >= rank2Number {
		t.Fatalf("Rank1=%d must be less than Rank2=%d", rank1Number, rank2Number)
	}
}
