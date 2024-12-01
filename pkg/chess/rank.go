package chess

const ranksCount = 8

// Rank1 has the lowest value, Rank8 has the highest value.
type Rank uint8

const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)
