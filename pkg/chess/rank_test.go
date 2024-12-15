package chess

import "testing"

func TestRankIsEnPassant(t *testing.T) {
	t.Parallel()

	tests := []struct {
		rank        Rank
		isEnPassant bool
	}{
		{RankNil, false},
		{Rank1, false},
		{Rank2, false},
		{Rank3, true},
		{Rank4, false},
		{Rank5, false},
		{Rank6, true},
		{Rank7, false},
		{Rank8, false},
		{Rank(123), false},
	}

	for _, test := range tests {
		t.Run(test.rank.String(), func(t *testing.T) {
			t.Parallel()

			isEnPassant := test.rank.IsEnPassant()
			if isEnPassant != test.isEnPassant {
				t.Fatalf("%s.IsEnPassant() expected %t but got %t", test.rank, test.isEnPassant, isEnPassant)
			}
		})
	}
}

func TestRankString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		rank Rank
		str  string
	}{
		{RankNil, "RankNil"},
		{Rank1, "Rank1"},
		{Rank2, "Rank2"},
		{Rank3, "Rank3"},
		{Rank4, "Rank4"},
		{Rank5, "Rank5"},
		{Rank6, "Rank6"},
		{Rank7, "Rank7"},
		{Rank8, "Rank8"},
		{Rank(123), "<unknown Rank=123>"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			t.Parallel()

			str := test.rank.String()
			if str != test.str {
				t.Fatalf("%s.String() expected %q but got %q", test.str, test.str, str)
			}
		})
	}
}
