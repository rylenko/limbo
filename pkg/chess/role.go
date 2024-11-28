package chess

type Role uint8

const (
	RoleBishop Role = iota
	RoleKing
	RoleKnight
	RolePawn
	RoleQueen
	RoleRook
)
