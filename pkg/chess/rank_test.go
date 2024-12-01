package chess

import "testing"

func TestRankOrder(t *testing.T) {
	t.Parallel()

	if Rank1 >= Rank2 {
		t.Fatalf("Rank1=%d >= Rank2=%d", Rank1, Rank2)
	}
}
