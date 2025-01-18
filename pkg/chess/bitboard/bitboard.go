package bitboard

import (
	"fmt"
	"errors"
	"math/bits"
	"unsafe"

	"github.com/rylenko/limbo/pkg/chess/square"
)

// Bitboard represents a mask of selected squares on the board using 64 bits unsigned integer.
//
// The most significant bit denotes the square A1, the least significant bit denotes the square H8.
//
// Zero value is ready to use.
type Bitboard uint64

const BitboardNil Bitboard = iota

var bitboardBitsCount = unsafe.Sizeof(BitboardNil)

// GetSquares gets all set squares in the bitboard.
func (bitboard Bitboard) GetSquares() []square.Square {
	squares := make([]square.Square, 0, bits.OnesCount(bitboard))

	for i := range bitboardBitsCount {
		if bitboard&(1<<(bitboardBitsCount-i-1)) != BitboardNil {
			squares = append(squares, square.Square(i+1))
		}
	}

	return squares
}

// Occupied checks that passed square occupied on the bitboard.
//
// TODO test.
func (bitboard Bitboard) Occupied(square Square) (bool, error) {
	if square == SquareNil || uint8(square) > bitboardBitsCount {
		return false, errors.New("square won't fit in the bitboard")
	}

	return bitboard>>(bitboardBitsCount-square)&1 == 1, nil
}

// Reverse reverses bitboard bits.
//
// TODO: test.
func (bitboard Bitboard) Reverse() Bitboard {
	return Bitboard(bits.Reverse64(uint64(bitboard)))
}

// SetSquares sets bitboard bits corresponding to the passed squares.
func (bitboard Bitboard) SetSquares(squares ...Square) (Bitboard, error) {
	for _, square := range squares {
		if square == square.SquareNil || uint8(square) > bitboardBitsCount {
			return bitboard, fmt.Errorf("square %s won't fit in the bitboard", square)
		}

		bitboard |= 1 << (bitboardBitsCount - square)
	}

	return bitboard, nil
}

// UnsetSquares unsets bitboard bits corresponding to the passed squares.
//
// TODO: test.
func (bitboard Bitboard) UnsetSquares(squares ...Square) (Bitboard, error) {
	unsetBitboard, err := BitboardNil.SetSquares(squares...)
	if err != nil {
		return BitboardNil, fmt.Errorf("SetSquares(%+v): %w", squares, err)
	}

	return bitboard & ^unsetBitboard, nil
}
