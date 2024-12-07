package chess

// Need to shift square lower values to more significant bits.
const bitboardSquareShift = 63

// Bitboard represents chess board using 64 bit integer.
//
// A1 is the most significant bit and H8 is the least significant bit.
//
// Zero value is ready to use.
type Bitboard uint64

// Get gets all set squares in the bitboard.
func (bitboard Bitboard) GetSquares() []Square {
	var squares []Square

	for i := 0; i < 64; i++ {
		if bitboard&(1<<(bitboardSquareShift-i)) != 0 {
			squares = append(squares, Square(i))
		}
	}

	return squares
}

// Set sets bits corresponding to the passed squares in the bitboard.
func (bitboard Bitboard) SetSquares(squares ...Square) Bitboard {
	for _, square := range squares {
		bitboard |= 1 << (bitboardSquareShift - uint8(square))
	}

	return bitboard
}
