package chess

import "testing"

func TestNewPieceTypeFromFEN(t *testing.T) {
	t.Parallel()

	const expectedErrString = "unknown byte"
	pieceTypeFENMap := map[byte]PieceType{
		'k': PieceTypeBlackKing,
		'q': PieceTypeBlackQueen,
		'r': PieceTypeBlackRook,
		'b': PieceTypeBlackBishop,
		'n': PieceTypeBlackKnight,
		'p': PieceTypeBlackPawn,
		'K': PieceTypeWhiteKing,
		'Q': PieceTypeWhiteQueen,
		'R': PieceTypeWhiteRook,
		'B': PieceTypeWhiteBishop,
		'N': PieceTypeWhiteKnight,
		'P': PieceTypeWhitePawn,
	}

	for i := 0; i < 256; i++ {
		if expectedPieceType, ok := pieceTypeFENMap[byte(i)]; ok {
			gotPieceType, err := NewPieceTypeFromFEN(byte(i))
			if err != nil {
				t.Fatalf("PieceTypeFromFEN(%d): %v", i, err)
			}

			if gotPieceType != expectedPieceType {
				t.Fatalf("PieceTypeFromFEN(%d) expected %d but got %d", i, expectedPieceType, gotPieceType)
			}

			continue
		}

		if _, err := NewPieceTypeFromFEN(byte(i)); err == nil || err.Error() != expectedErrString {
			t.Fatalf("PieceTypeFromFEN(%d) expected error %q but got %q", i, expectedErrString, err)
		}
	}
}
