package chess

type Square struct {
	file File
	rank Rank
}

func (square *Square) Color() Color {
	if uint8(square.file) % 2 == uint8(square.rank) % 2 {
		return ColorBlack
	}
	return ColorWhite
}

func NewSquare(file File, rank Rank) Square {
	return Square{
		file: file,
		rank: rank,
	}
}
