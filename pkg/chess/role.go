package chess

type Role uint8

const (
	RoleKing Role = iota
	RoleQueen
	RoleRook
	RoleBishop
	RoleKnight
	RolePawn
)
