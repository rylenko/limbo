package chess

import (
	set "github.com/deckarep/golang-set/v2"
)

// A1 is the most significant bit and H8 is the least significant bit.
type bitboard uint64

func newBitboard(squares set.Set[Square]) bitboard {
	bits := uint64(0)

	squares.Each(func(square Square) bool {
		const bitMaxShift = 63
		bits |= 1 << (bitMaxShift - uint8(square))

		return false
	})

	return bitboard(bits)
}
