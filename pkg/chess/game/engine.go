package game

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

// CalcPieceMoves calculates all possible piece moves in the passed position.
//
// TODO: test.
// TODO: generate default moves and castlings.
func (engine Engine) CalcPieceMoves(position *Position, piece Piece) ([]Move, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	TODO CASTLING

	color, err := piece.Color()
	if err != nil {
		return nil, fmt.Errorf("%s.Color(): %w", piece, err)
	}

	if color != position.activeColor {
		return nil, nil
	}

	bitboard, ok := position.board.bitboards[piece]
	if !ok {
		return nil, errors.New("piece bitboard not found")
	}

	var moves []Move

	for _, origin := range bitboard.GetSquares() {
		movesFromOrigin, err := engine.calcPieceMovesFromOrigin(position, piece, origin)
		if err != nil {
			return nil, fmt.Errorf("calcPieceMovesFromOrigin(%s, %s): %w", piece, origin, err)
		}

		moves = append(moves, movesFromOrigin...)
	}

	return moves, nil
}

// addRawMoveAttackTags adds, for example, capture or check tags to the passed move if needed.
//
// Note that the moves are raw, that is, for example, the king moves can put him in checkmate.
//
// TODO: test.
func (engine Engine) addRawMoveAttackTags(position *Position, move *Move, attackColor Color) error {
	if position == nil {
		return errors.New("position is nil")
	}

	if move == nil {
		return errors.New("move is nil")
	}

	defendColor, err := attackColor.Opposite()
	if err != nil {
		return fmt.Errorf("%s.Opposite(): %w", attackColor, err)
	}

	destOccupiedByDefendColor, err := position.board.OccupiedByColor(move.dest, defendColor)
	if err != nil {
		return fmt.Errorf("Occupied(%s): %w", move.dest, err)
	}

	if destOccupiedByDefendColor {
		move.tags.Set(MoveTagCapture)
	}

	originPiece, err := position.board.GetPieceFromSquare(move.origin)
	if err != nil {
		return fmt.Errorf("GetPieceFromSquare(%s): %w", move.origin, err)
	}

	if originPiece == PieceNil {
		return fmt.Errorf("no piece on the origin %s", move.origin)
	}

	originPieceRole, err := originPiece.Role()
	if err != nil {
		return fmt.Errorf("%s.Role(): %w", originPiece, err)
	}

	if move.dest == position.enPassantSquare && originPieceRole == RolePawn {
		move.tags.Set(MoveTagEnPassantCapture)
	}

	putsDefendColorInCheck, err := engine.checkPutsColorInCheck(position, *move, defendColor)
	if err != nil {
		return fmt.Errorf("checkPutsColorInCheck(%+v, %s): %w", *move, defendColor, err)
	}

	if putsDefendColorInCheck {
		move.tags.Set(MoveTagCheck)
	}

	return nil
}

// calcPieceMovesFromOrigin calculates all possible piece moves in the passed position from passed origin.
//
// Before calling this function make sure the piece is actually on the passed origin.
//
// TODO: test
// TODO: generate default moves and castlings.
func (engine Engine) calcPieceMovesFromOrigin(position *Position, piece Piece, origin Square) ([]Move, error) {
	if position == nil {
		return nil, errors.New("position is nil")
	}

	color, err := piece.Color()
	if err != nil {
		return nil, fmt.Errorf("%s.Color(): %w", piece, err)
	}

	rawDestsBitboard, err := engine.calcPieceRawMoveDestsBitboard(position, piece, origin)
	if err != nil {
		return nil, fmt.Errorf("calcPieceRawMoveDestsBitboard(%s, %s): %w", piece, origin, err)
	}

	rawDests := rawDestsBitboard.GetSquares()

	moves := make([]Move, 0, len(rawDests))

	for _, rawDest := range rawDests {
		rank, err := rawDest.Rank()
		if err != nil {
			return nil, fmt.Errorf("%s.Rank(): %w", rawDest, err)
		}

		rawMoves := make([]Move, 0, 1)

		if piece.NeedPromoInRank(rank) {
			rawMoves = NewMovesPromo(origin, rawDest, MoveTagsNil)
		} else {
			rawMoves = append(rawMoves, NewMove(origin, rawDest, MoveTagsNil, RoleNil))
		}

		for _, rawMove := range rawMoves {
			putsHisColorInCheck, err := engine.checkPutsColorInCheck(position, rawMove, color)
			if err != nil {
				return nil, fmt.Errorf("checkPutsColorInCheck(%+v, %s): %w", rawMove, color, err)
			}

			if putsHisColorInCheck {
				continue
			}

			if err := engine.addRawMoveAttackTags(position, &rawMove, color); err != nil {
				return nil, fmt.Errorf("addRawMoveAttackTags(%+v, %s): %w", rawMove, color, err)
			}

			moves = append(moves, rawMove)
		}
	}

	return moves, nil
}

// calcBishopRawMoveDestsBitboard calculates passed color bishop raw move destinations bitboard in passed position from
// passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put him in checkmate.
//
// TODO: test.
func (engine Engine) calcBishopRawMoveDestsBitboard(position *Position, origin Square, color Color) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	bitboard, err := engine.calcDiagonalsRawMoveDestsBitboard(position, origin, color)
	if err != nil {
		return BitboardNil, fmt.Errorf("calcDiagonalsRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
	}

	return bitboard, nil
}

// calcDiagonalsRawMoveDestsBitboard calculates passed color diagonal and antidiagonal raw move destinations Bitboard
// in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their in checkmate.
//
// TODO: test.
func (engine Engine) calcDiagonalsRawMoveDestsBitboard(
	position *Position,
	origin Square,
	color Color,
) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	diagonalBitboard, ok := moveDiagonalRawDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("diagonal bitboard not found")
	}

	diagonalDestsBitboard, err := engine.calcLinearRawMoveDestsBitboard(position, origin, diagonalBitboard, color)
	if err != nil {
		return BitboardNil, fmt.Errorf(
			"calcLinearRawMoveDestsBitboard(%s, 0x%X, %s): %w", origin, diagonalBitboard, color, err)
	}

	antidiagonalBitboard, ok := moveAntidiagonalRawDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("antidiagonal bitboard not found")
	}

	antidiagonalDestsBitboard, err := engine.calcLinearRawMoveDestsBitboard(position, origin, antidiagonalBitboard, color)
	if err != nil {
		return BitboardNil, fmt.Errorf(
			"calcLinearRawMoveDestsBitboard(%s, 0x%X, %s): %w", origin, antidiagonalBitboard, color, err)
	}

	return diagonalDestsBitboard | antidiagonalDestsBitboard, nil
}

// calcHorVertRawMoveDestsBitboard calculates passed color horizontal and vertical raw move destinations Bitboard in
// passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their in checkmate.
//
// TODO: test.
func (engine Engine) calcHorVertRawMoveDestsBitboard(position *Position, origin Square, color Color) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	horizontalBitboard, ok := moveHorizontalRawDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("horizontal bitboard not found")
	}

	horizontalDestsBitboard, err := engine.calcLinearRawMoveDestsBitboard(position, origin, horizontalBitboard, color)
	if err != nil {
		return BitboardNil, fmt.Errorf(
			"calcLinearRawMoveDestsBitboard(%s, 0x%X, %s): %w", origin, horizontalBitboard, color, err)
	}

	verticalBitboard, ok := moveVerticalRawDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("vertical bitboard not found")
	}

	verticalDestsBitboard, err := engine.calcLinearRawMoveDestsBitboard(position, origin, verticalBitboard, color)
	if err != nil {
		return BitboardNil, fmt.Errorf(
			"calcLinearRawMoveDestsBitboard(%s, 0x%X, %s): %w", origin, verticalBitboard, color, err)
	}

	return horizontalDestsBitboard | verticalDestsBitboard, nil
}

// calcKingRawMoveDestsBitboard calculates passed color king raw move destinations bitboard in passed position from
// passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put him in checkmate.
//
// TODO: test.
func (engine Engine) calcKingRawMoveDestsBitboard(position *Position, origin Square, color Color) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	colorBitboard, err := position.board.GetColorBitboard(color)
	if err != nil {
		return BitboardNil, fmt.Errorf("GetColorBitboard(%s): %w", color, err)
	}

	bitboard, ok := moveKingRawDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("bitboard not found")
	}

	return bitboard & ^colorBitboard, nil
}

// calcKnightRawMoveDestsBitboard calculates passed color knight raw move destinations bitboard in passed position from
// passed origin.
//
// Note that the moves are raw, that is, for example, the knight moves can put their king in checkmate.
//
// TODO: test.
func (engine Engine) calcKnightRawMoveDestsBitboard(position *Position, origin Square, color Color) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	colorBitboard, err := position.board.GetColorBitboard(color)
	if err != nil {
		return BitboardNil, fmt.Errorf("GetColorBitboard(%s): %w", color, err)
	}

	bitboard, ok := moveKnightRawDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("bitboard not found")
	}

	return bitboard & ^colorBitboard, nil
}

// calcLinearRawMoveDestsBitboard calculates passed color linear raw move Bitboard in passed position from passed
// origin.
//
// Note that the moves are raw, that is, for example, the king moves can put them in checkmate.
//
// TODO: test.
func (engine Engine) calcLinearRawMoveDestsBitboard(
	position *Position,
	origin Square,
	line Bitboard,
	color Color,
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

	colorBitboard, err := position.board.GetColorBitboard(color)
	if err != nil {
		return BitboardNil, fmt.Errorf("GetColorBitboard(%s): %w", color, err)
	}

	return movesToBlockerBitboard & ^colorBitboard, nil
}

// calcPawnRawMoveDestsBitboard calculates passed color pawn raw move destinations bitboard in passed position from
// passed origin.
//
// Note that the moves are raw, that is, for example, the pawn moves can put their king in checkmate.
//
// TODO: test.
func (engine Engine) calcPawnRawMoveDestsBitboard(
	position *Position,
	origin Square,
	attackColor Color,
) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	rank, err := origin.Rank()
	if err != nil {
		return BitboardNil, fmt.Errorf("%s.Rank(): %w", origin, err)
	}

	if RolePawn.CanBeInRank(rank) {
		return BitboardNil, fmt.Errorf(
			"invalid pawn rank %s: either a wrong move occurred or the promotion was not completed", rank)
	}

	defendColor, err := attackColor.Opposite()
	if err != nil {
		return BitboardNil, fmt.Errorf("%s.Opposite(): %w", attackColor, err)
	}

	allCapturesBitboard, err := position.board.GetColorBitboard(defendColor)
	if err != nil {
		return BitboardNil, fmt.Errorf("GetColorBitboard(%s): %w", defendColor, err)
	}

	if position.enPassantSquare != SquareNil {
		allCapturesBitboard, err = allCapturesBitboard.SetSquares(position.enPassantSquare)
		if err != nil {
			return BitboardNil, fmt.Errorf("SetSquares(%s): %w", position.enPassantSquare, err)
		}
	}

	originBitboard, err := BitboardNil.SetSquares(origin)
	if err != nil {
		return BitboardNil, fmt.Errorf("SetSquares(%s): %w", origin, err)
	}

	occupiedBitboard, err := position.board.GetOccupiedBitboard()
	if err != nil {
		return BitboardNil, fmt.Errorf("GetOccupiedBitboard(): %w", err)
	}
	unoccupiedBitboard := ^occupiedBitboard

	file, err := origin.File()
	if err != nil {
		return BitboardNil, fmt.Errorf("%s.File(): %w", origin, err)
	}

	var bitboard Bitboard

	switch attackColor {
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
		return BitboardNil, errors.New("no moves for ColorNil")
	default:
		return BitboardNil, fmt.Errorf("unknown color %s", attackColor)
	}

	return bitboard, nil
}

// calcPieceRawMoveDestsBitboard calculates piece raw move destinations bitboard in passed position from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their king in checkmate.
//
// TODO: test.
func (engine Engine) calcPieceRawMoveDestsBitboard(position *Position, piece Piece, origin Square) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	color, err := piece.Color()
	if err != nil {
		return BitboardNil, fmt.Errorf("%s.Color(): %w", piece, err)
	}

	role, err := piece.Role()
	if err != nil {
		return BitboardNil, fmt.Errorf("%s.Role(): %w", piece, err)
	}

	switch role {
	case RoleKing:
		rawBitboard, err := engine.calcKingRawMoveDestsBitboard(position, origin, color)
		if err != nil {
			return BitboardNil, fmt.Errorf("calcKingRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
		}

		return rawBitboard, nil
	case RoleQueen:
		rawBitboard, err := engine.calcQueenRawMoveDestsBitboard(position, origin, color)
		if err != nil {
			return BitboardNil, fmt.Errorf("calcQueenRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
		}

		return rawBitboard, nil
	case RoleRook:
		rawBitboard, err := engine.calcRookRawMoveDestsBitboard(position, origin, color)
		if err != nil {
			return BitboardNil, fmt.Errorf("calcRookRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
		}

		return rawBitboard, nil
	case RoleBishop:
		rawBitboard, err := engine.calcBishopRawMoveDestsBitboard(position, origin, color)
		if err != nil {
			return BitboardNil, fmt.Errorf("calcBishopRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
		}

		return rawBitboard, nil
	case RoleKnight:
		rawBitboard, err := engine.calcKnightRawMoveDestsBitboard(position, origin, color)
		if err != nil {
			return BitboardNil, fmt.Errorf("calcKnightRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
		}

		return rawBitboard, nil
	case RolePawn:
		rawBitboard, err := engine.calcPawnRawMoveDestsBitboard(position, origin, color)
		if err != nil {
			return BitboardNil, fmt.Errorf("calcPawnRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
		}

		return rawBitboard, nil
	case RoleNil:
		return BitboardNil, errors.New("RoleNil always has no destinations")
	default:
		return BitboardNil, fmt.Errorf("unknown role %s", role)
	}
}

// calcQueenRawMoveDestsBitboard calculates passed color queen raw move destinations bitboard in passed position from
// passed origin.
//
// Note that the moves are raw, that is, for example, the queen moves can put her king in checkmate.
//
// TODO: Test.
func (engine Engine) calcQueenRawMoveDestsBitboard(position *Position, origin Square, color Color) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	horVertBitboard, err := engine.calcHorVertRawMoveDestsBitboard(position, origin, color)
	if err != nil {
		return BitboardNil, fmt.Errorf("calcHorVertRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
	}

	diagonalBitboard, err := engine.calcDiagonalsRawMoveDestsBitboard(position, origin, color)
	if err != nil {
		return BitboardNil, fmt.Errorf("calcDiagonalsRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
	}

	return horVertBitboard | diagonalBitboard, nil
}

// calcRookRawMoveDestsBitboard calculates passed color rook raw move destinations bitboard in passed position from
// passed origin.
//
// Note that the moves are raw, that is, for example, the rook moves can put his king in checkmate.
//
// TODO: Test.
func (engine Engine) calcRookRawMoveDestsBitboard(position *Position, origin Square, color Color) (Bitboard, error) {
	if position == nil {
		return BitboardNil, errors.New("position is nil")
	}

	bitboard, err := engine.calcHorVertRawMoveDestsBitboard(position, origin, color)
	if err != nil {
		return BitboardNil, fmt.Errorf("calcHorVertRawMoveDestsBitboard(%s, %s): %w", origin, color, err)
	}

	return bitboard, nil
}

// checkAnySquaresOpenToAttack checks that passed square open to attack by pieces of passed color.
//
// Note that the function determines whether a square is open for attack, but does not check whether an attack is
// possible. For example, a square may be open for attack, but after an attack, the attacker may put his king in
// checkmate.
//
// TODO: test.
func (engine Engine) checkAnySquaresOpenToAttack(
	position *Position,
	defendColor Color,
	squares ...Square,
) (bool, error) {
	if position == nil {
		return false, errors.New("position is nil")
	}

	for _, square := range squares {
		attacked, err := engine.checkSquareOpenToAttack(position, defendColor, square)
		if err != nil {
			return false, fmt.Errorf("checkSquareOpenToAttack(%s, %s): %w", defendColor, square, err)
		}

		if attacked {
			return true, nil
		}
	}

	return false, nil
}

// checkChecked checks that the passed color king in check in passed position.
//
// TODO: test.
func (engine Engine) checkChecked(position *Position, color Color) (bool, error) {
	if position == nil {
		return false, errors.New("position is nil")
	}

	king, err := NewPiece(color, RoleKing)
	if err != nil {
		return false, fmt.Errorf("NewPiece(%s, %s): %w", color, RoleKing, err)
	}

	kingSquares := position.board.bitboards[king].GetSquares()
	if len(kingSquares) != 1 {
		return false, fmt.Errorf("expected 1 king square but got %d", len(kingSquares))
	}
	kingSquare := kingSquares[0]

	inCheck, err := engine.checkSquareOpenToAttack(position, color, kingSquare)
	if err != nil {
		return false, fmt.Errorf("checkSquareOpenToAttack(%s, %s): %w", color, kingSquare, err)
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

	newPosition, err := position.DeepCopy()
	if err != nil {
		return false, fmt.Errorf("DeepCopy(): %w", err)
	}

	if err := newPosition.MoveRaw(move); err != nil {
		return false, fmt.Errorf("Move(%+v): %w", move, err)
	}

	checked, err := engine.checkChecked(newPosition, color)
	if err != nil {
		return false, fmt.Errorf("checkChecked(%s): %w", color, err)
	}

	return checked, nil
}

// checkSquareOpenToAttack checks that passed square open to attack by pieces of attack color.
//
// Note that the function determines whether a square is open for attack, but does not check whether an attack is
// possible. For example, a square may be open for attack, but after an attack, the attacker may put his king in
// checkmate.
//
// TODO: test.
func (engine Engine) checkSquareOpenToAttack(position *Position, defendColor Color, square Square) (bool, error) {
	if position == nil {
		return false, errors.New("position is nil")
	}

	attackColor, err := defendColor.Opposite()
	if err != nil {
		return false, fmt.Errorf("%s.Opposite(): %w", defendColor, err)
	}

	queen, err := NewPiece(attackColor, RoleQueen)
	if err != nil {
		return false, fmt.Errorf("NewPiece(%s, %s): %w", attackColor, RoleQueen, err)
	}

	squareQueenBitboard, err := engine.calcQueenRawMoveDestsBitboard(position, square, defendColor)
	if err != nil {
		return false, fmt.Errorf("calcQueenRawMoveDestsBitboard(%s, %s): %w", square, defendColor, err)
	}

	if position.board.bitboards[queen]&squareQueenBitboard != BitboardNil {
		return true, nil
	}

	rook, err := NewPiece(attackColor, RoleRook)
	if err != nil {
		return false, fmt.Errorf("NewPiece(%s, %s): %w", attackColor, RoleRook, err)
	}

	squareRookBitboard, err := engine.calcRookRawMoveDestsBitboard(position, square, defendColor)
	if err != nil {
		return false, fmt.Errorf("calcRookRawMoveDestsBitboard(%s, %s): %w", square, defendColor, err)
	}

	if position.board.bitboards[rook]&squareRookBitboard != BitboardNil {
		return true, nil
	}

	bishop, err := NewPiece(attackColor, RoleBishop)
	if err != nil {
		return false, fmt.Errorf("NewPiece(%s, %s): %w", attackColor, RoleBishop, err)
	}

	squareBishopBitboard, err := engine.calcBishopRawMoveDestsBitboard(position, square, defendColor)
	if err != nil {
		return false, fmt.Errorf("calcBishopRawMoveDestsBitboard(%s, %s): %w", square, defendColor, err)
	}

	if position.board.bitboards[bishop]&squareBishopBitboard != BitboardNil {
		return true, nil
	}

	knight, err := NewPiece(attackColor, RoleKnight)
	if err != nil {
		return false, fmt.Errorf("NewPiece(%s, %s): %w", attackColor, RoleKnight, err)
	}

	squareKnightBitboard, err := engine.calcKnightRawMoveDestsBitboard(position, square, defendColor)
	if err != nil {
		return false, fmt.Errorf("calcKnightRawMoveDestsBitboard(%s, %s): %w", square, defendColor, err)
	}

	if position.board.bitboards[knight]&squareKnightBitboard != BitboardNil {
		return true, nil
	}

	pawn, err := NewPiece(attackColor, RolePawn)
	if err != nil {
		return false, fmt.Errorf("NewPiece(%s, %s): %w", attackColor, RolePawn, err)
	}

	squareBitboard, err := BitboardNil.SetSquares(square)
	if err != nil {
		return false, fmt.Errorf("0x%X.SetSquares(%s): %w", BitboardNil, square, err)
	}

	for _, pawnSquare := range position.board.bitboards[pawn].GetSquares() {
		pawnDestsBitboard, err := engine.calcPawnRawMoveDestsBitboard(position, pawnSquare, attackColor)
		if err != nil {
			return false, fmt.Errorf("calcPawnRawMoveDestsBitboard(%s, %s): %w", pawnSquare, attackColor, err)
		}

		if pawnDestsBitboard&squareBitboard != BitboardNil {
			return true, nil
		}
	}

	return false, nil
}
