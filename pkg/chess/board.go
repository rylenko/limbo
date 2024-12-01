package chess

import (
	"fmt"
	"strings"

	set "github.com/deckarep/golang-set/v2"
)

type Board struct {
	pieceTypeBitboardMap map[PieceType]bitboard
}

func NewBoard(pieceTypeBitboardMap map[PieceType]bitboard) *Board {
	return &Board{
		pieceTypeBitboardMap: pieceTypeBitboardMap,
	}
}

func NewBoardFromSquarePieceTypeMap(squarePieceTypeMap map[Square]PieceType) *Board {
	pieceTypeBitboardMap := make(map[PieceType]bitboard, len(PieceTypes))

	// Usually the most squares are occupied by the pawn piece type. The number of such squares is 8.
	const squaresCapacity = 8
	squares := set.NewThreadUnsafeSetWithSize[Square](squaresCapacity)

	for _, pieceType := range PieceTypes {
		squares.Clear()

		for square, squarePieceType := range squarePieceTypeMap {
			if squarePieceType == pieceType {
				squares.Add(square)
			}
		}

		pieceTypeBitboardMap[pieceType] = newBitboard(squares)
	}

	return NewBoard(pieceTypeBitboardMap)
}

// Parses board's FEN part to the structure.
//
// FEN argument example: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR".
func NewBoardFromFEN(fen string) (*Board, error) {
	// One part for one rank on the board.
	const fenPartsLenRequired = 8

	fenParts := strings.Split(strings.TrimSpace(fen), "/")
	if len(fenParts) != fenPartsLenRequired {
		return nil, fmt.Errorf("FEN parts len required %d but got %d", fenPartsLenRequired, len(fenParts))
	}

	squarePieceTypeMap := make(map[Square]PieceType, squaresCount)

	for fenPartIndex, fenPart := range fenParts {
		rank := Rank(uint8(Rank8) - uint8(fenPartIndex))
		fenPartFilesCount := 0

		for fenPartByteIndex, fenPartByte := range []byte(fenPart) {
			if '1' <= fenPartByte && fenPartByte <= '9' {
				fenPartFilesCount += int(fenPartByte - '0')
				continue
			}

			piece, err := NewPieceTypeFromFEN(fenPartByte)
			if err != nil {
				return nil, fmt.Errorf(
					"part #%d, byte #%d, NewPieceTypeFromFEN(%d): %w", fenPartIndex, fenPartByteIndex, fenPartByte, err)
			}

			squarePieceTypeMap[NewSquare(File(fenPartFilesCount), rank)] = piece
			fenPartFilesCount++
		}

		if fenPartFilesCount != int(filesCount) {
			return nil, fmt.Errorf("part #%d has invalid files count != %d", fenPartIndex, filesCount)
		}
	}

	return NewBoardFromSquarePieceTypeMap(squarePieceTypeMap), nil
}
