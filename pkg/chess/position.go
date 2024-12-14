package chess

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	// Required FEN parts: board, active color, castling rights, en passant, half move clock and full move number.
	positionFENPartsCount = 6
	positionStartFEN      = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
)

// Position represents the state of the game at a certain point in time.
type Position struct {
	board           *Board
	activeColor     Color
	castlingRights  CastlingRights
	enPassantSquare Square
	halfMoveClock   uint8
	fullMoveNumber  uint16
}

// NewPosition creates a new position with passed parameters.
func NewPosition(
	board *Board,
	activeColor Color,
	castlingRights CastlingRights,
	enPassantSquare Square,
	halfMoveClock uint8,
	fullMoveNumber uint16,
) *Position {
	return &Position{
		board:           board,
		activeColor:     activeColor,
		castlingRights:  castlingRights,
		enPassantSquare: enPassantSquare,
		halfMoveClock:   halfMoveClock,
		fullMoveNumber:  fullMoveNumber,
	}
}

// NewPositionStart creates game start position.
func NewPositionStart() (*Position, error) {
	position, err := NewPositionFromFEN(positionStartFEN)
	if err != nil {
		return nil, fmt.Errorf("NewPositionFromFEN(%q): %w", positionStartFEN, err)
	}

	return position, nil
}

// NewPositionFromFEN parses FEN to the Position structure.
//
// FEN argument example: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1".
func NewPositionFromFEN(fen string) (*Position, error) {
	parts := strings.Split(fen, " ")
	if len(parts) != positionFENPartsCount {
		return nil, fmt.Errorf("FEN parts required %d but got %d", positionFENPartsCount, len(parts))
	}

	board, err := NewBoardFromFEN(parts[0])
	if err != nil {
		return nil, fmt.Errorf("NewBoardFromFEN(%q): %w", parts[0], err)
	}

	activeColor, err := NewColorFromFEN(parts[1])
	if err != nil {
		return nil, fmt.Errorf("NewColorFromFEN(%q): %w", parts[1], err)
	}

	castlingRights, err := NewCastlingRightsFromFEN(parts[2])
	if err != nil {
		return nil, fmt.Errorf("NewCastlingRightsFromFEN(%q): %w", parts[2], err)
	}

	enPassantSquare, err := NewSquareEnPassantFromFEN(parts[3])
	if err != nil {
		return nil, fmt.Errorf("NewSquareEnPassantFromFEN(%q): %w", parts[3], err)
	}

	halfMoveClockUint64, err := strconv.ParseUint(parts[4], 10, 8)
	if err != nil {
		return nil, fmt.Errorf("ParseUint(%s, 10, 8): %w", parts[4], err)
	}
	halfMoveClock := uint8(halfMoveClockUint64)

	fullMoveNumberUint64, err := strconv.ParseUint(parts[5], 10, 16)
	if err != nil {
		return nil, fmt.Errorf("ParseUint(%s, 10, 16): %w", parts[5], err)
	}
	fullMoveNumber := uint16(fullMoveNumberUint64)

	return NewPosition(board, activeColor, castlingRights, enPassantSquare, halfMoveClock, fullMoveNumber), nil
}

// CalcMoves calculates all possible moves in the current position.
//
// TODO: test.
func (position *Position) CalcMoves() ([]Move, error) {
	// TODO: generate default moves and castlings.

	var moves []Move

	pieces, err := NewPiecesOfColor(position.activeColor)
	if err != nil {
		return nil, fmt.Errorf("NewPiecesOfColor(%s): %w", position.activeColor, err)
	}

	for _, piece := range pieces {
		pieceMoves, err := position.CalcPieceMoves(piece)
		if err != nil {
			return nil, fmt.Errorf("CalcPieceMoves(%s): %w", piece, err)
		}

		moves = append(moves, pieceMoves...)
	}

	return moves, nil
}

// CalcPieceMoves calculates all possible piece moves in the current position from passed origin.
//
// TODO: test.
func (position *Position) CalcPieceMoves(piece Piece) ([]Move, error) {
	// TODO: generate default moves and castlings.

	color, err := piece.Color()
	if err != nil {
		return nil, fmt.Errorf("%s.Color(): %w", piece, err)
	}

	if color != position.activeColor {
		return nil, nil
	}

	var moves []Move

	for _, origin := range position.board.bitboards[piece].GetSquares() {
		rawDests, err := position.calcPieceRawMoveDests(piece, origin)
		if err != nil {
			return nil, fmt.Errorf("calcPieceRawMoveDests(%s, %s): %w", piece, origin, err)
		}

		for _, rawDest := range rawDests {
			rank, err := rawDest.Rank()
			if err != nil {
				return nil, fmt.Errorf("%s.Rank(): %w", rawDest, err)
			}

			isPromo := (piece == PieceWhitePawn && rank == Rank8) || (piece == PieceBlackPawn && rank == Rank1)

			moves = append(moves, NewMove(origin, rawDest, isPromo))
		}
	}

	return moves, nil
}

// calcHorVertRawMoveDestsBitboard calculates horizontal and vertical raw move Bitboard from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put them in checkmate.
//
// TODO: test.
func (position *Position) calcHorVertRawMoveDestsBitboard(origin Square) (Bitboard, error) {
	rank, err := origin.Rank()
	if err != nil {
		return BitboardNil, fmt.Errorf("%s.Rank(): %w", origin, err)
	}

	file, err := origin.File()
	if err != nil {
		return BitboardNil, fmt.Errorf("%s.File(): %w", origin, err)
	}

	rankSquares, err := NewSquaresOfRank(rank)
	if err != nil {
		return BitboardNil, fmt.Errorf("NewSquaresOfRank(%s): %w", rank, err)
	}

	fileSquares, err := NewSquaresOfFile(file)
	if err != nil {
		return BitboardNil, fmt.Errorf("NewSquaresOfFile(%s): %w", file, err)
	}

	rankBitboard, err := BitboardNil.SetSquares(rankSquares...)
	if err != nil {
		return BitboardNil, fmt.Errorf("SetSquares(%v): %w", rankSquares, err)
	}

	fileBitboard, err := BitboardNil.SetSquares(fileSquares...)
	if err != nil {
		return BitboardNil, fmt.Errorf("SetSquares(%v): %w", fileSquares, err)
	}

	rankDests, err := position.calcLinearRawMoveDestsBitboard(origin, rankBitboard)
	if err != nil {
		return BitboardNil, fmt.Errorf("calcLinearRawMoveDestsBitboard(%s, 0x%X): %w", origin, rankBitboard, err)
	}

	fileDests, err := position.calcLinearRawMoveDestsBitboard(origin, fileBitboard)
	if err != nil {
		return BitboardNil, fmt.Errorf("calcLinearRawMoveDestsBitboard(%s, 0x%X): %w", origin, fileBitboard, err)
	}

	return rankDests | fileDests, nil
}

// calcKingRawMoveDests calculates king raw move destinations from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put him in checkmate.
//
// TODO: test.
func (position *Position) calcKingRawMoveDests(origin Square) ([]Square, error) {
	colorBitboard, err := position.board.GetColorBitboard(position.activeColor)
	if err != nil {
		return nil, fmt.Errorf("GetColorBitboard(%s): %w", position.activeColor, err)
	}

	bitboard := roleKingRawMoveDestBitboards[origin] & ^colorBitboard

	return bitboard.GetSquares(), nil
}

// calcKnightRawMoveDests calculates knight raw move destinations from passed origin.
//
// Note that the moves are raw, that is, for example, the knight moves can put their king in checkmate.
//
// TODO: test.
func (position *Position) calcKnightRawMoveDests(origin Square) ([]Square, error) {
	colorBitboard, err := position.board.GetColorBitboard(position.activeColor)
	if err != nil {
		return nil, fmt.Errorf("GetColorBitboard(%s): %w", position.activeColor, err)
	}

	bitboard := roleKnightRawMoveDestBitboards[origin] & ^colorBitboard

	return bitboard.GetSquares(), nil
}

// calcLinearRawMoveDestsBitboard calculates linear raw move Bitboard from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put them in checkmate.
//
// TODO: test.
func (position *Position) calcLinearRawMoveDestsBitboard(origin Square, line Bitboard) (Bitboard, error) {
	originBitboard, err := BitboardNil.SetSquares(origin)
	if err != nil {
		return BitboardNil, fmt.Errorf("SetSquares(%s): %w", origin, err)
	}

	occupiedBitboard, err := position.board.GetOccupiedBitboard()
	if err != nil {
		return BitboardNil, fmt.Errorf("GetOccupiedBitboard(): %w", err)
	}

	occupiedLineBitboard := occupiedBitboard & line

	movesToBlockerBitboard := line & ((occupiedLineBitboard - 2*originBitboard) ^
		(occupiedLineBitboard.Reverse() - 2*originBitboard.Reverse()).Reverse())

	colorBitboard, err := position.board.GetColorBitboard(position.activeColor)
	if err != nil {
		return BitboardNil, fmt.Errorf("GetColorBitboard(%s): %w", position.activeColor, err)
	}

	return movesToBlockerBitboard & ^colorBitboard, nil
}

// calcPawnRawMoveDests calculates pawn raw move destinations from passed origin.
//
// Note that the moves are raw, that is, for example, the pawn moves can put their king in checkmate.
//
// TODO: test.
func (position *Position) calcPawnRawMoveDests(origin Square) ([]Square, error) {
	rank, err := origin.Rank()
	if err != nil {
		return nil, fmt.Errorf("%s.Rank(): %w", origin, err)
	}

	if (position.activeColor == ColorBlack && rank == Rank1) || (position.activeColor == ColorWhite && rank == Rank8) {
		return nil, nil
	}

	activeColorOpposite, err := position.activeColor.Opposite()
	if err != nil {
		return nil, fmt.Errorf("%s.Opposite(): %w", position.activeColor, err)
	}

	allCapturesBitboard, err := position.board.GetColorBitboard(activeColorOpposite)
	if err != nil {
		return nil, fmt.Errorf("GetColorBitboard(%s): %w", activeColorOpposite, err)
	}

	if position.enPassantSquare != SquareNil {
		allCapturesBitboard, err = allCapturesBitboard.SetSquares(position.enPassantSquare)
		if err != nil {
			return nil, fmt.Errorf("SetSquares(%s): %w", position.enPassantSquare, err)
		}
	}

	originBitboard, err := BitboardNil.SetSquares(origin)
	if err != nil {
		return nil, fmt.Errorf("SetSquares(%s): %w", origin, err)
	}

	occupiedBitboard, err := position.board.GetOccupiedBitboard()
	if err != nil {
		return nil, fmt.Errorf("GetOccupiedBitboard(): %w", err)
	}
	unoccupiedBitboard := ^occupiedBitboard

	file, err := origin.File()
	if err != nil {
		return nil, fmt.Errorf("%s.File(): %w", origin, err)
	}

	var bitboard Bitboard

	switch position.activeColor {
	case ColorBlack:
		bitboard |= originBitboard << len(files) & unoccupiedBitboard

		if rank == Rank7 {
			bitboard |= originBitboard << (2 * len(files)) & unoccupiedBitboard //nolint:mnd // Skip all files twice.
		}
		if file != FileA {
			bitboard |= originBitboard << (len(files) + 1) & allCapturesBitboard
		}
		if file != FileH {
			bitboard |= originBitboard << (len(files) - 1) & allCapturesBitboard
		}
	case ColorWhite:
		bitboard |= originBitboard >> len(files) & unoccupiedBitboard

		if rank == Rank2 {
			bitboard |= originBitboard >> (2 * len(files)) & unoccupiedBitboard //nolint:mnd // Skip all files twice.
		}
		if file != FileA {
			bitboard |= originBitboard >> (len(files) - 1) & allCapturesBitboard
		}
		if file != FileH {
			bitboard |= originBitboard >> (len(files) + 1) & allCapturesBitboard
		}
	case ColorNil:
		return nil, errors.New("no moves for ColorNil")
	default:
		return nil, fmt.Errorf("unknown color %s", position.activeColor)
	}

	return bitboard.GetSquares(), nil
}

// calcPieceRawMoveDests calculates piece raw move destinations from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their king in checkmate.
//
// TODO: test.
func (position *Position) calcPieceRawMoveDests(piece Piece, origin Square) ([]Square, error) {
	color, err := piece.Color()
	if err != nil {
		return nil, fmt.Errorf("%s.Color(): %w", piece, err)
	}

	if color != position.activeColor {
		return nil, nil
	}

	role, err := piece.Role()
	if err != nil {
		return nil, fmt.Errorf("%s.Role(): %w", piece, err)
	}

	switch role {
	case RoleKing:
		rawDests, err := position.calcKingRawMoveDests(origin)
		if err != nil {
			return nil, fmt.Errorf("calcKingRawMoveDests(%s): %w", origin, err)
		}
		return rawDests, nil
	case RoleQueen:
		return []Square{}, nil
	case RoleRook:
		rawBitboard, err := position.calcHorVertRawMoveDestsBitboard(origin)
		if err != nil {
			return nil, fmt.Errorf("calcHorVertRawMoveDestsBitboard(%s): %w", origin, err)
		}
		return rawBitboard.GetSquares(), nil
	case RoleBishop:
		return []Square{}, nil
	case RoleKnight:
		rawDests, err := position.calcKnightRawMoveDests(origin)
		if err != nil {
			return nil, fmt.Errorf("calcKnightRawMoveDests(%s): %w", origin, err)
		}
		return rawDests, nil
	case RolePawn:
		rawDests, err := position.calcPawnRawMoveDests(origin)
		if err != nil {
			return nil, fmt.Errorf("calcPawnRawMoveDests(%s): %w", origin, err)
		}
		return rawDests, nil
	case RoleNil:
		return nil, errors.New("RoleNil always has no destinations")
	default:
		return nil, fmt.Errorf("unknown role %s", role)
	}
}
