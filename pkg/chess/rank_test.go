package chess

import "testing"

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
