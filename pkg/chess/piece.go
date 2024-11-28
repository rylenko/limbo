package chess

// Piece is the structure of pieces on the board.
type Piece struct {
	color Color
	role  Role
}

func NewPiece(color Color, role Role) *Piece {
	return &Piece{
		color: color,
		role:  role,
	}
}
