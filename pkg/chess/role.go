package chess

import "fmt"

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
