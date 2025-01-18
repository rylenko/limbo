package chess

import "testing"

func TestCanBeInRank(t *testing.T) {
	t.Parallel()

	tests := []struct {
		role  Role
		ranks []Rank
		can   bool
	}{
		{RoleKing, []Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}, true},
		{RoleQueen, []Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}, true},
		{RoleRook, []Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}, true},
		{RoleBishop, []Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}, true},
		{RoleKnight, []Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}, true},
		{RolePawn, ranks[1:7], true},
		{RolePawn, []Rank{Rank1, Rank8}, false},
	}

	for _, test := range tests {
		t.Run(test.role.String(), func(t *testing.T) {
			t.Parallel()

			for _, rank := range test.ranks {
				can := test.role.CanBeInRank(rank)
				if can != test.can {
					t.Fatalf("%s.CanBeInRank(%s) expected %t but got %t", test.role, rank, test.can, can)
				}
			}
		})
	}
}
