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
	parts := strings.Split(fen, "/")
	if len(parts) != len(ranks) {
		return nil, fmt.Errorf("required %d parts separated by \"/\" but got %d", len(ranks), len(parts))
	}

	bitboards := make(map[Piece]Bitboard)

	rank := Rank8

	for partIndex, part := range parts {
		file := FileA

		for byteIndex, bytee := range []byte(part) {
			if '1' <= bytee && bytee <= '9' {
				file = File(uint8(file) + bytee - '0')
				continue
			}

			piece, err := NewPieceFromFEN(string(bytee))
			if err != nil {
				return nil, fmt.Errorf("part #%d, byte #%d, NewPieceFromFEN(%q): %w", partIndex, byteIndex, bytee, err)
			}

			square, err := NewSquare(rank, file)
			if err != nil {
				return nil, fmt.Errorf("part #%d, byte #%d, NewSquare(%s, %s): %w", partIndex, byteIndex, rank, file, err)
			}

			bitboards[piece], err = bitboards[piece].SetSquares(square)
			if err != nil {
				return nil, fmt.Errorf("part #%d, byte #%d, SetSquares(%s): %w", partIndex, byteIndex, square, err)
			}

			file++
		}

		if file-1 != FileH {
			return nil, fmt.Errorf("invalid files count in part #%d", partIndex)
		}

		rank--
	}

	return NewBoard(bitboards), nil
}

// GetColorBitboard returns bitboard of occupied squares by pieces of passed color.
func (board *Board) GetColorBitboard(color Color) (Bitboard, error) {
	var bitboard Bitboard

	pieces, err := NewPiecesOfColor(color)
	if err != nil {
		return BitboardNil, fmt.Errorf("NewPiecesOfColor(%s): %w", color, err)
	}

	for _, piece := range pieces {
		bitboard |= board.bitboards[piece]
	}

	return bitboard, nil
}

// GetOccupiedBitboard returns bitboard of unoccupied squares.
func (board *Board) GetOccupiedBitboard() (Bitboard, error) {
	whites, err := board.GetColorBitboard(ColorWhite)
	if err != nil {
		return BitboardNil, fmt.Errorf("GetColorBitboard(%s): %w", ColorWhite, err)
	}

	blacks, err := board.GetColorBitboard(ColorBlack)
	if err != nil {
		return BitboardNil, fmt.Errorf("GetColorBitboard(%s): %w", ColorBlack, err)
	}

	return whites | blacks, nil
}

// GetPieceFromSquare returns a piece that is on the passed square or PieceNil if the square is not occupied.
//
// Please note that even if no error has occurred, a piece may be PieceNil if there is no piece on the passed square.
//
// TODO test.
func (board *Board) GetPieceFromSquare(square Square) (Piece, error) {
	for piece, bitboard := range board.bitboards {
		occupied, err := bitboard.Occupied(square)
		if err != nil {
			return PieceNil, fmt.Errorf("0x%X.Occupied(%s): %w", bitboard, square, err)
		}

		if occupied {
			return piece, nil
		}
	}

	return PieceNil, nil
}

// MoveRaw makes a raw move on the current board.
//
// Note that the move is raw, so it was not validated.
//
// TODO: test.
func (board *Board) MoveRaw(move Move) error {
	originPiece, err := board.removePieceFromSquare(move.origin)
	if err != nil {
		return fmt.Errorf("removePieceFromSquare(%s): %w", move.origin, err)
	}

	if _, err := board.removePieceFromSquare(move.dest); err != nil {
		return fmt.Errorf("removeFromSquare(%s): %w", move.dest, err)
	}

	destNewPiece := originPiece

	if move.promo != RoleNil {
		originPieceColor, err := originPiece.Color()
		if err != nil {
			return fmt.Errorf("%s.Color(): %w", originPiece, err)
		}

		destNewPiece, err = NewPiece(move.promoRole, originPieceColor)
		if err != nil {
			return fmt.Errorf("NewPiece(%s, %s): %w", move.promoRole, originPieceColor, err)
		}
	}

	if err := board.setPieceToSquare(destNewPiece, move.dest); err != nil {
		return fmt.Errorf("setPieceToOrigin(%s, %s): %w",
	}

	if move.tags.Contains(MoveTagEnPassantCapture) {
		TODO EN PASSANT
	}

	return nil
}

// OccupiedByColor checks that passed square is occupied by valid piece of passed color.
//
// TODO: test.
func (board *Board) OccupiedByColor(square Square, color Color) (bool, error) {
	occupiedBitboard, err := board.GetColorBitboard(color)
	if err != nil {
		return false, fmt.Errorf("GetOccupiedBitboard(%s): %w", color, err)
	}

	occupied, err := occupiedBitboard.Occupied(square)
	if err != nil {
		return false, fmt.Errorf("0x%X.Occupied(%s): %w", occupiedBitboard, square, err)
	}

	return occupied, nil
}

// removePieceFromSquare removes piece from the passed square if exists.
//
// Please note that each figure has its own set of squares. If the squares of some figures intersect, firstly, this is an erroneous behavior, and secondly, only the first one that comes across will be removed.
//
// TODO: test
func (board *Board) removePieceFromSquare(square Square) (Piece, error) {
	piece, err := board.GetPieceFromSquare(square)
	if err != nil {
		return PieceNil, fmt.Errorf("GetPieceFromSquare(%s): %w", square, err)
	}

	board.bitboards[piece], err = board.bitboards[piece].UnsetSquares(square)
	if err != nil {
		return PieceNil, fmt.Errorf("0x%X.UnsetSquares(%s): %w", board.bitboards[piece], square, err)
	}

	return piece, nil
}

func (board *Board) setPieceToSquare(piece Piece, square Square) error {
	newBitboard, err := board.bitboards[piece].SetSquares(square)
	if err != nil {
		return fmt.Errorf("0x%X.SetSquares(%s): %w", board.bitboards[piece], square, err)
	}

	board.bitboards[piece] = newBitboard

	return nil
}
