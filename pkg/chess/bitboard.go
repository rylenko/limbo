package chess

import (
	set "github.com/deckarep/golang-set/v2"
)

type bitboard uint64

func newBitboard(squares set.Set[Square]) bitboard {
	bits := uint64(0)

	squares.Each(func(square Square) bool {
		bits |= 1 << uint64(square)
		return false
	})

	return bitboard(bits)
}
