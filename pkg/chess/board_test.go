package chess

import (
	"reflect"
	"testing"

	set "github.com/deckarep/golang-set/v2"
)

func TestNewBoard(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		boardMap map[Square]PieceType
		board    *Board
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
			},
			&Board{
				pieceTypeBitboards: map[PieceType]bitboard{
					PieceTypeWhiteKing:   newBitboard(set.NewThreadUnsafeSet[Square](SquareE1)),
					PieceTypeWhiteQueen:  newBitboard(set.NewThreadUnsafeSet[Square]()),
					PieceTypeWhiteRook:   newBitboard(set.NewThreadUnsafeSet[Square](SquareA1, SquareH1)),
					PieceTypeWhiteKnight: newBitboard(set.NewThreadUnsafeSet[Square]()),
					PieceTypeWhiteBishop: newBitboard(set.NewThreadUnsafeSet[Square]()),
					PieceTypeWhitePawn:   newBitboard(set.NewThreadUnsafeSet[Square]()),
					PieceTypeBlackKing:   newBitboard(set.NewThreadUnsafeSet[Square](SquareE8)),
					PieceTypeBlackQueen:  newBitboard(set.NewThreadUnsafeSet[Square]()),
					PieceTypeBlackRook:   newBitboard(set.NewThreadUnsafeSet[Square](SquareA8, SquareH8)),
					PieceTypeBlackKnight: newBitboard(set.NewThreadUnsafeSet[Square]()),
					PieceTypeBlackBishop: newBitboard(set.NewThreadUnsafeSet[Square]()),
					PieceTypeBlackPawn:   newBitboard(set.NewThreadUnsafeSet[Square]()),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			gotBoard := NewBoard(test.boardMap)

			if !reflect.DeepEqual(gotBoard, test.board) {
				t.Fatalf("Board of map %v expected %v but got %v", test.boardMap, test.board, gotBoard)
			}
		})
	}
}
