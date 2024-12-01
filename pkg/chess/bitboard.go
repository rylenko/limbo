package chess

// Bitboard represents chess board using 64 bit integer.
//
// A1 is the most significant bit and H8 is the least significant bit.
//
// Zero value is ready to use.
type Bitboard uint64

// Set sets the bit corresponding to the passed square in the bitboard.
func (bitboard Bitboard) Set(square Square) Bitboard {
	// Need to shift square lower values to more significant bits.
	const bitMaxShift = 63

	return bitboard | 1<<(bitMaxShift-uint8(square))
}
