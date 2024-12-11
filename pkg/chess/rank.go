package chess

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

var (
	// Array of all valid ranks.
	ranks = [8]Rank{Rank1, Rank2, Rank3, Rank4, Rank5, Rank6, Rank7, Rank8}

	// Mapping of all Rank variants to strings.
	rankStrings = map[Rank]string{
		RankNil: "RankNil",
		Rank1: "Rank1",
		Rank2: "Rank2",
		Rank3: "Rank3",
		Rank4: "Rank4",
		Rank5: "Rank5",
		Rank6: "Rank6",
		Rank7: "Rank7",
		Rank8: "Rank8",
	}
)

func (rank Rank) String() string {
	str, ok := rankStrings[rank]
	if !ok {
		return "<unknown Rank>"
	}
	return str
}
