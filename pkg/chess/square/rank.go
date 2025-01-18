package square

import "fmt"

// Rank is the enumeration of all chess board ranks.
type Rank uint8

const (
	RankNil Rank = iota
	Rank1
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

// Array of all valid ranks.
var Ranks = [8]Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}

// EnPassantIsPossible returns true if en passant square can belong to this rank.
func (rank Rank) EnPassantIsPossible() bool {
	return rank == Rank3 || rank == Rank6
}

// String returns string representation of current rank.
func (rank Rank) String() string {
	switch rank {
	case RankNil:
		return "RankNil"
	case Rank1:
		return "Rank1"
	case Rank2:
		return "Rank2"
	case Rank3:
		return "Rank3"
	case Rank4:
		return "Rank4"
	case Rank5:
		return "Rank5"
	case Rank6:
		return "Rank6"
	case Rank7:
		return "Rank7"
	case Rank8:
		return "Rank8"
	default:
		return fmt.Sprintf("<unknown Rank=%d>", rank)
	}
}
