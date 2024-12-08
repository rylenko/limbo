package chess

import (
	"fmt"
	"strings"
)

// Board is the collection of Bitboards, representing chess board.
type Board struct {
	bitboards map[Piece]Bitboard
}

// NewBoard creates new Board with passed parameters.
func NewBoard(bitboards map[Piece]Bitboard) *Board {
	return &Board{
		bitboards: bitboards,
	}
}

// NewBoardFromFEN parses board's FEN part to the Board structure.
//
// FEN argument example: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR".
func NewBoardFromFEN(fen string) (*Board, error) {
	fenParts := strings.Split(fen, "/")
	if len(fenParts) != len(ranks) {
		return nil, fmt.Errorf("required %d parts but got %d", len(ranks), len(fenParts))
	}

	bitboards := make(map[Piece]Bitboard)

	for fenPartIndex, fenPart := range fenParts {
		fenPartRank := Rank(uint8(Rank8) - uint8(fenPartIndex))
		fenPartFilesCount := uint8(0)

		for fenPartByteIndex, fenPartByte := range []byte(fenPart) {
			if '1' <= fenPartByte && fenPartByte <= '9' {
				fenPartFilesCount += fenPartByte - '0'
				continue
			}

			fenPartByteString := string(fenPartByte)

			piece, err := NewPieceFromFEN(fenPartByteString)
			if err != nil {
				return nil, fmt.Errorf(
					"part #%d, byte #%d, NewPieceFromFEN(%q): %w", fenPartIndex, fenPartByteIndex, fenPartByteString, err)
			}

			square := NewSquare(File(fenPartFilesCount), fenPartRank)
			bitboards[piece] = bitboards[piece].SetSquares(square)

			fenPartFilesCount++
		}

		if int(fenPartFilesCount) != len(files) {
			return nil, fmt.Errorf("required %d files but got %d in part #%d", len(files), fenPartFilesCount, fenPartIndex)
		}
	}

	return NewBoard(bitboards), nil
}

// GetColorBitboard returns bitboard of occupied squares by pieces of passed color.
func (board *Board) GetColorBitboard(color Color) Bitboard {
	var bitboard Bitboard

	for _, piece := range NewPiecesOfColor(color) {
		bitboard |= board.bitboards[piece]
	}

	return bitboard
}

// GetOccupiedBitboard returns bitboard of unoccupied squares.
func (board *Board) GetOccupiedBitboard() Bitboard {
	return board.GetColorBitboard(ColorWhite) | board.GetColorBitboard(ColorBlack)
}
