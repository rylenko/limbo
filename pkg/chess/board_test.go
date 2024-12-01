package chess

import (
	"reflect"
	"testing"

	set "github.com/deckarep/golang-set/v2"
)

func TestNewBoardFromSquarePieceTypeMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name               string
		squarePieceTypeMap map[Square]PieceType
		board              *Board
	}{
		{
			"rooks and kings",
			map[Square]PieceType{
				SquareA1: PieceTypeWhiteRook,
				SquareE1: PieceTypeWhiteKing,
				SquareH1: PieceTypeWhiteRook,
				SquareA8: PieceTypeBlackRook,
				SquareE8: PieceTypeBlackKing,
				SquareH8: PieceTypeBlackRook,
				// Must be ignored.
				SquareH8 + 1: PieceTypeBlackQueen,
			},
			NewBoard(map[PieceType]bitboard{
				PieceTypeWhiteKing:   newBitboard(set.NewThreadUnsafeSet[Square](SquareE1)),
				PieceTypeWhiteQueen:  newBitboard(set.NewThreadUnsafeSet[Square]()),
				PieceTypeWhiteRook:   newBitboard(set.NewThreadUnsafeSet[Square](SquareA1, SquareH1)),
				PieceTypeWhiteBishop: newBitboard(set.NewThreadUnsafeSet[Square]()),
				PieceTypeWhiteKnight: newBitboard(set.NewThreadUnsafeSet[Square]()),
				PieceTypeWhitePawn:   newBitboard(set.NewThreadUnsafeSet[Square]()),
				PieceTypeBlackKing:   newBitboard(set.NewThreadUnsafeSet[Square](SquareE8)),
				PieceTypeBlackQueen:  newBitboard(set.NewThreadUnsafeSet[Square]()),
				PieceTypeBlackRook:   newBitboard(set.NewThreadUnsafeSet[Square](SquareA8, SquareH8)),
				PieceTypeBlackBishop: newBitboard(set.NewThreadUnsafeSet[Square]()),
				PieceTypeBlackKnight: newBitboard(set.NewThreadUnsafeSet[Square]()),
				PieceTypeBlackPawn:   newBitboard(set.NewThreadUnsafeSet[Square]()),
			}),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			gotBoard := NewBoardFromSquarePieceTypeMap(test.squarePieceTypeMap)

			if !reflect.DeepEqual(gotBoard, test.board) {
				t.Fatalf("Board from %v expected %v but got %v", test.squarePieceTypeMap, test.board, gotBoard)
			}
		})
	}
}
