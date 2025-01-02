package chess

import (
	"fmt"
	"math/bits"
)

// Need to shift square lower values to more significant bits.
const bitboardBitsCount = 64

// Bitboard represents chess board using 64 bit integer.
//
// SquareA1 is the most significant bit and SquareH8 is the least significant bit. Make sure that SquareA1 has 1 value
// and SquareH8 has 64 value.
//
// Zero value is ready to use.
type Bitboard uint64

const BitboardNil Bitboard = iota

// GetSquares gets all set squares in the bitboard.
func (bitboard Bitboard) GetSquares() []Square {
	var squares []Square

	for i := range bitboardBitsCount {
		if bitboard&(1<<(bitboardBitsCount-i-1)) != 0 {
			squares = append(squares, Square(i+1))
		}
	}

	return squares
}

// Occupied checks that passed square occupied on the bitboard.
//
// TODO test.
func (bitboard Bitboard) Occupied(square Square) (bool, error) {
	if uint8(square) == 0 || uint8(square) > bitboardBitsCount {
		return false, fmt.Errorf("invalid square %s", square)
	}

	return bitboard>>(bitboardBitsCount-square)&1 == 1, nil
}

// Reverse reverses bitboard bits.
//
// TODO: test
func (bitboard Bitboard) Reverse() Bitboard {
	return Bitboard(bits.Reverse64(uint64(bitboard)))
}

// SetSquares sets bits corresponding to the passed squares in the bitboard.
func (bitboard Bitboard) SetSquares(squares ...Square) (Bitboard, error) {
	for _, square := range squares {
		if uint8(square) == 0 || uint8(square) > bitboardBitsCount {
			return bitboard, fmt.Errorf("invalid square %s", square)
		}

		bitboard |= 1 << (bitboardBitsCount - uint8(square))
	}

	return bitboard, nil
}

// UnsetSquares unsets bits corresponding to the passed squares in the bitboard.
//
// TODO: test
func (bitboard Bitboard) UnsetSquares(squares ...Squares) (Bitboard, error) {
	unsetBitboard, err := BitboardNil.SetSquares(squares...)
	if err != nil {
		return BitboardNil, fmt.Errorf("SetSquares(%+v): %w", squares, err)
	}

	return bitboard & ^unsetBitboard, nil
}
