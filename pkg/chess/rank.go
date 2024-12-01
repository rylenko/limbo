package chess

const ranksCount uint8 = 8

// 1 has the lowest value, 8 has the highest value.
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
