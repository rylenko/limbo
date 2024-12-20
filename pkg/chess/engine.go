package chess

import (
	"errors"
	"fmt"
)

// The engine is responsible for the logic of movement and interaction.
type Engine struct{}

// CalcMoves calculates all possible moves in passed position for active color pieces.
//
// TODO: test.
// TODO: generate default moves and castlings.
func (engine Engine) CalcMoves(position *Position) ([]Move, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	var moves []Move

	pieces, err := NewPiecesOfColor(position.activeColor)
	if err != nil {
		return nil, fmt.Errorf("NewPiecesOfColor(%s): %w", position.activeColor, err)
	}

	for _, piece := range pieces {
		pieceMoves, err := engine.CalcPieceMoves(position, piece)
		if err != nil {
			return nil, fmt.Errorf("CalcPieceMoves(%s): %w", piece, err)
		}

		moves = append(moves, pieceMoves...)
	}

	return moves, nil
}

// CalcPieceMoves calculates all possible piece moves in the passed position from passed origin.
//
// TODO: test.
// TODO: generate default moves and castlings.
func (engine Engine) CalcPieceMoves(position *Position, piece Piece) ([]Move, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	color, err := piece.Color()
	if err != nil {
		return nil, fmt.Errorf("%s.Color(): %w", piece, err)
	}

	if color != position.activeColor {
		return nil, nil
	}

	var moves []Move

	for _, origin := range position.board.bitboards[piece].GetSquares() {
		rawDests, err := engine.calcPieceRawMoveDests(position, piece, origin)
		if err != nil {
			return nil, fmt.Errorf("calcPieceRawMoveDests(%s, %s): %w", piece, origin, err)
		}

		for _, rawDest := range rawDests {
			rank, err := rawDest.Rank()
			if err != nil {
				return nil, fmt.Errorf("%s.Rank(): %w", rawDest, err)
			}

			rawMove := NewMove(origin, rawDest, piece.NeedPromoInRank(rank))

			putsActiveColorInCheck, err := engine.checkPutsColorInCheck(position, rawMove, position.activeColor)
			if err != nil {
				return nil, fmt.Errorf("checkPutsColorInCheck(%+v, %+v, %s): %w", position, rawMove, position.activeColor, err)
			}

			if !putsActiveColorInCheck {
				moves = append(moves, rawMove)
			}
		}
	}

	return moves, nil
}

// calcBishopRawMoveDests calculates bishop raw move destinations in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put him in checkmate.
//
// TODO: test.
func (engine Engine) calcBishopRawMoveDests(position *Position, origin Square) ([]Square, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	bitboard, err := engine.calcDiagonalsRawMoveDestsBitboard(position, origin)
	if err != nil {
		return nil, fmt.Errorf("calcDiagonalsRawMoveDestsBitboard(%s): %w", origin, err)
	}

	return bitboard.GetSquares(), nil
}

// calcDiagonalsRawMoveDestsBitboard calculates diagonal and antidiagonal raw move destinations Bitboard in passed
// position from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their in checkmate.
//
// TODO: test.
func (engine Engine) calcDiagonalsRawMoveDestsBitboard(position *Position, origin Square) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	diagonalBitboard, err := moveGetDiagonalRawDestsBitboard(origin)
	if err != nil {
		return BitboardNil, fmt.Errorf("moveGetDiagonalRawDestsBitboard(%s): %w", origin, err)
	}

	diagonalDestsBitboard, err := engine.calcLinearRawMoveDestsBitboard(position, origin, diagonalBitboard)
	if err != nil {
		return BitboardNil, fmt.Errorf("calcLinearRawMoveDestsBitboard(%s, 0x%X): %w", origin, diagonalBitboard, err)
	}

	antidiagonalBitboard, err := moveGetAntidiagonalRawDestsBitboard(origin)
	if err != nil {
		return BitboardNil, fmt.Errorf("moveGetAntidiagonalRawDestsBitboard(%s): %w", origin, err)
	}

	antidiagonalDestsBitboard, err := engine.calcLinearRawMoveDestsBitboard(position, origin, antidiagonalBitboard)
	if err != nil {
		return BitboardNil, fmt.Errorf(
			"calcLinearRawMoveDestsBitboard(%s, 0x%X): %w", origin, antidiagonalBitboard, err)
	}

	return diagonalDestsBitboard | antidiagonalDestsBitboard, nil
}

// calcHorVertRawMoveDestsBitboard calculates horizontal and vertical raw move destinations Bitboard in passed
// position from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their in checkmate.
//
// TODO: test.
func (engine Engine) calcHorVertRawMoveDestsBitboard(position *Position, origin Square) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	horizontalBitboard, err := moveGetHorizontalRawDestsBitboard(origin)
	if err != nil {
		return BitboardNil, fmt.Errorf("moveGetHorizontalRawDestsBitboard(%s): %w", origin, err)
	}

	horizontalDestsBitboard, err := engine.calcLinearRawMoveDestsBitboard(position, origin, horizontalBitboard)
	if err != nil {
		return BitboardNil, fmt.Errorf(
			"calcLinearRawMoveDestsBitboard(%s, 0x%X): %w", origin, horizontalBitboard, err)
	}

	verticalBitboard, err := moveGetVerticalRawDestsBitboard(origin)
	if err != nil {
		return BitboardNil, fmt.Errorf("moveGetVerticalRawDestsBitboard(%s): %w", origin, err)
	}

	verticalDestsBitboard, err := engine.calcLinearRawMoveDestsBitboard(position, origin, verticalBitboard)
	if err != nil {
		return BitboardNil, fmt.Errorf(
			"calcLinearRawMoveDestsBitboard(%s, 0x%X): %w", origin, verticalBitboard, err)
	}

	return horizontalDestsBitboard | verticalDestsBitboard, nil
}

// calcKingRawMoveDests calculates king raw move destinations in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put him in checkmate.
//
// TODO: test.
func (engine Engine) calcKingRawMoveDests(position *Position, origin Square) ([]Square, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	colorBitboard, err := position.board.GetColorBitboard(position.activeColor)
	if err != nil {
		return nil, fmt.Errorf("GetColorBitboard(%s): %w", position.activeColor, err)
	}

	bitboard, err := moveGetKingRawDestsBitboard(origin)
	if err != nil {
		return nil, fmt.Errorf("moveGetKingRawDestsBitboard(%s): %w", origin, err)
	}
	bitboard &= ^colorBitboard

	return bitboard.GetSquares(), nil
}

// calcKnightRawMoveDests calculates knight raw move destinations in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the knight moves can put their king in checkmate.
//
// TODO: test.
func (engine Engine) calcKnightRawMoveDests(position *Position, origin Square) ([]Square, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	colorBitboard, err := position.board.GetColorBitboard(position.activeColor)
	if err != nil {
		return nil, fmt.Errorf("GetColorBitboard(%s): %w", position.activeColor, err)
	}

	bitboard, err := moveGetKnightRawDestsBitboard(origin)
	if err != nil {
		return nil, fmt.Errorf("moveGetKnightRawDestsBitboard(%s): %w", origin, err)
	}
	bitboard &= ^colorBitboard

	return bitboard.GetSquares(), nil
}

// calcLinearRawMoveDestsBitboard calculates linear raw move Bitboard in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put them in checkmate.
//
// TODO: test.
func (engine Engine) calcLinearRawMoveDestsBitboard(
	position *Position,
	origin Square,
	line Bitboard,
) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

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

// calcPawnRawMoveDests calculates pawn raw move destinations in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the pawn moves can put their king in checkmate.
//
// TODO: test.
func (engine Engine) calcPawnRawMoveDests(position *Position, origin Square) ([]Square, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	rank, err := origin.Rank()
	if err != nil {
		return nil, fmt.Errorf("%s.Rank(): %w", origin, err)
	}

	if RolePawn.CanBeInRank(rank) {
		return nil, fmt.Errorf("invalid pawn rank %s: either a wrong move occurred or the promotion was not completed", rank)
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

		if PieceBlackPawn.IsPawnLongMovePossibleFromRank(rank) {
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

		if PieceWhitePawn.IsPawnLongMovePossibleFromRank(rank) {
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

// calcPieceRawMoveDests calculates piece raw move destinations in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their king in checkmate.
//
// TODO: test.
func (engine Engine) calcPieceRawMoveDests(position *Position, piece Piece, origin Square) ([]Square, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

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
		rawDests, err := engine.calcKingRawMoveDests(position, origin)
		if err != nil {
			return nil, fmt.Errorf("calcKingRawMoveDests(%s): %w", origin, err)
		}

		return rawDests, nil
	case RoleQueen:
		rawDests, err := engine.calcQueenRawMoveDests(position, origin)
		if err != nil {
			return nil, fmt.Errorf("calcQueenRawMoveDests(%s): %w", origin, err)
		}

		return rawDests, nil
	case RoleRook:
		rawDests, err := engine.calcRookRawMoveDests(position, origin)
		if err != nil {
			return nil, fmt.Errorf("calcRookRawMoveDests(%s): %w", origin, err)
		}

		return rawDests, nil
	case RoleBishop:
		rawDests, err := engine.calcBishopRawMoveDests(position, origin)
		if err != nil {
			return nil, fmt.Errorf("calcBishopRawMoveDests(%s): %w", origin, err)
		}

		return rawDests, nil
	case RoleKnight:
		rawDests, err := engine.calcKnightRawMoveDests(position, origin)
		if err != nil {
			return nil, fmt.Errorf("calcKnightRawMoveDests(%s): %w", origin, err)
		}

		return rawDests, nil
	case RolePawn:
		rawDests, err := engine.calcPawnRawMoveDests(position, origin)
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

// calcQueenRawMoveDests calculates queen raw move destinations in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the queen moves can put her king in checkmate.
//
// TODO: Test.
func (engine Engine) calcQueenRawMoveDests(position *Position, origin Square) ([]Square, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	horVertBitboard, err := engine.calcHorVertRawMoveDestsBitboard(position, origin)
	if err != nil {
		return nil, fmt.Errorf("calcHorVertRawMoveDestsBitboard(%s): %w", origin, err)
	}

	diagonalBitboard, err := engine.calcDiagonalsRawMoveDestsBitboard(position, origin)
	if err != nil {
		return nil, fmt.Errorf("calcDiagonalsRawMoveDestsBitboard(%s): %w", origin, err)
	}

	bitboard := horVertBitboard | diagonalBitboard

	return bitboard.GetSquares(), nil
}

// calcRookRawMoveDests calculates rook raw move destinations in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the rook moves can put his king in checkmate.
//
// TODO: Test.
func (engine Engine) calcRookRawMoveDests(position *Position, origin Square) ([]Square, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	bitboard, err := engine.calcHorVertRawMoveDestsBitboard(position, origin)
	if err != nil {
		return nil, fmt.Errorf("calcHorVertRawMoveDestsBitboard(%s): %w", origin, err)
	}

	return bitboard.GetSquares(), nil
}

// checkInCheck checks that the passed position is in check, that is, the king of the active color is attacked.
//
// TODO: test.
func (engine Engine) checkInCheck(position *Position) (bool, error) {
	if position == nil {
		return false, errors.New("position is nil")
	}

	kingPiece, err := NewPiece(position.activeColor, RoleKing)
	if err != nil {
		return false, fmt.Errorf("NewPiece(%s, %s): %w", position.activeColor, RoleKing, err)
	}

	// In the classic game there is only one square here.
	kingSquares := position.board.bitboards[kingPiece].GetSquares()

	inCheck, err := engine.checkSquaresAttacked(position, kingSquares...)
	if err != nil {
		return false, fmt.Errorf("checkSquareIsAttacked(%+v, %s): %w", position, kingPiece, err)
	}

	return inCheck, nil
}

// checkMovePutsInCheck checks that passed move will put the king of passed color in check.
//
// TODO: test.
func (engine Engine) checkPutsColorInCheck(position *Position, move Move, color Color) (bool, error) {
	if position == nil {
		return false, errors.New("position is nil")
	}

	newPosition, err := position.Move(move)
	if err != nil {
		return false, fmt.Errorf("Move(%+v): %w", move, err)
	}

	newPosition.activeColor = color

	inCheck, err := engine.checkInCheck(newPosition)
	if err != nil {
		return false, fmt.Errorf("checkInCheck(%+v): %w", newPosition, err)
	}

	return inCheck, nil
}

// checkSquaresAttacked checks that passed squares are attacked by pieces of opposite of active color.
//
// TODO: test.
func (engine Engine) checkSquaresAttacked(position *Position, _ ...Square) (bool, error) {
	if position == nil {
		return false, errors.New("position is nil")
	}

	// TODO

	return true, nil
}
