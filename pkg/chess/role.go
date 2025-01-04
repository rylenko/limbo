package chess

import "fmt"

// Role represents Piece role.
type Role uint8

const (
	RoleNil Role = iota
	RoleKing
	RoleQueen
	RoleRook
	RoleBishop
	RoleKnight
	RolePawn
)

// All roles for which promotion is possible.
var rolePromos = [...]Role{RoleQueen, RoleRook, RoleBishop, RoleKnight}

// CanBeInRank returns true if current role can be located in passed rank.
func (role Role) CanBeInRank(rank Rank) bool {
	// Pawns cannot move backwards, if a distant rank is reached an immediate promotion must occur.
	return role != RolePawn || (rank != Rank1 && rank != Rank8)
}

// IsEnPassantPossibleInRank returns true if current role is pawn and En Passant possible in passed rank.
func (role Role) IsEnPassantPossibleInRank(rank Rank) bool {
	return role == RolePawn && (rank == Rank3 || rank == Rank6)
}

// String returns string representation of current role.
func (role Role) String() string {
	switch role {
	case RoleNil:
		return "RoleNil"
	case RoleKing:
		return "RoleKing"
	case RoleQueen:
		return "RoleQueen"
	case RoleRook:
		return "RoleRook"
	case RoleBishop:
		return "RoleBishop"
	case RoleKnight:
		return "RoleKnight"
	case RolePawn:
		return "RolePawn"
	default:
		return fmt.Sprintf("<unknown Role=%d>", role)
	}
}
