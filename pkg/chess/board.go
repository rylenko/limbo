package chess

import (
	"fmt"
	"strings"
)

type Board struct {
	pieceTypeBitboardMap map[PieceType]Bitboard
}

func NewBoard(pieceTypeBitboardMap map[PieceType]Bitboard) *Board {
	return &Board{
		pieceTypeBitboardMap: pieceTypeBitboardMap,
	}
}

// Parses board's FEN part to the Board structure.
//
// FEN argument example: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR".
func NewBoardFromFEN(fen string) (*Board, error) {
	fenParts := strings.Split(strings.TrimSpace(fen), "/")
	if len(fenParts) != ranksCount {
		return nil, fmt.Errorf("required %d parts but got %d", ranksCount, len(fenParts))
	}

	pieceTypeBitboardMap := make(map[PieceType]Bitboard)

	for fenPartIndex, fenPart := range fenParts {
		fenPartRank := Rank(uint8(Rank8) - uint8(fenPartIndex))
		fenPartFilesCount := uint8(0)

		for fenPartByteIndex, fenPartByte := range []byte(fenPart) {
			if '1' <= fenPartByte && fenPartByte <= '9' {
				fenPartFilesCount += fenPartByte - '0'
				continue
			}

			pieceType, err := NewPieceTypeFromFEN(fenPartByte)
			if err != nil {
				return nil, fmt.Errorf(
					"part #%d, byte #%d, NewPieceTypeFromFEN(%d): %w", fenPartIndex, fenPartByteIndex, fenPartByte, err)
			}

			square := NewSquare(File(fenPartFilesCount), fenPartRank)
			pieceTypeBitboardMap[pieceType] = pieceTypeBitboardMap[pieceType].Set(square)

			fenPartFilesCount++
		}

		if fenPartFilesCount != filesCount {
			return nil, fmt.Errorf("required %d files but got %d in part #%d", filesCount, fenPartFilesCount, fenPartIndex)
		}
	}

	return NewBoard(pieceTypeBitboardMap), nil
}
