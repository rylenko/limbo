package chess

import (
	set "github.com/deckarep/golang-set/v2"
)

type Board struct {
	pieceTypeBitboards map[PieceType]bitboard
}

func NewBoard(boardMap map[Square]PieceType) *Board {
	board := &Board{
		pieceTypeBitboards: make(map[PieceType]bitboard, len(PieceTypes)),
	}

	// Usually the most squares are occupied by the pawn piece type. The number of such squares is 8.
	const squaresCapacity int = 8
	squares := set.NewThreadUnsafeSetWithSize[Square](squaresCapacity)

	for _, pieceType := range PieceTypes {
		squares.Clear()

		for square, squarePieceType := range boardMap {
			if squarePieceType == pieceType {
				squares.Add(square)
			}
		}

		board.pieceTypeBitboards[pieceType] = newBitboard(squares)
	}

	return board
}
