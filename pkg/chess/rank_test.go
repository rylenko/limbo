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

func TestRankIsPawnLongMove(t *testing.T) {
	t.Parallel()

	tests := []struct {
		rank           Rank
		color          Color
		isPawnLongMove bool
	}{
		{RankNil, ColorNil, false},
		{Rank1, ColorWhite, false},
		{Rank2, ColorWhite, true},
		{Rank2, ColorBlack, false},
		{Rank3, ColorWhite, false},
		{Rank4, ColorBlack, false},
		{Rank5, ColorBlack, false},
		{Rank6, ColorWhite, false},
		{Rank7, ColorBlack, true},
		{Rank7, ColorWhite, false},
		{Rank8, ColorBlack, false},
		{Rank(123), Color(111), false},
	}

	for _, test := range tests {
		t.Run(test.rank.String(), func(t *testing.T) {
			t.Parallel()

			isPawnLongMove := test.rank.IsPawnLongMove(test.color)
			if isPawnLongMove != test.isPawnLongMove {
				t.Fatalf("%s.IsPawnLongMove(%s) expected %t but got %t", test.rank, test.color, test.isPawnLongMove, isPawnLongMove)
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
