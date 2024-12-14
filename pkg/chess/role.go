package chess

import (
	"errors"
	"fmt"
)

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

var (
	// Note that the moves are raw, that is, for example, the piece moves can put him in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	roleAntidiagonalRawMoveDestBitboards = map[Square]Bitboard{
		SquareA1: 0x8000000000000000,
		SquareB1: 0x4080000000000000,
		SquareC1: 0x2040800000000000,
		SquareD1: 0x1020408000000000,
		SquareE1: 0x0810204080000000,
		SquareF1: 0x0408102040800000,
		SquareG1: 0x0204081020408000,
		SquareH1: 0x0102040810204080,
		SquareA2: 0x4080000000000000,
		SquareB2: 0x2040800000000000,
		SquareC2: 0x1020408000000000,
		SquareD2: 0x0810204080000000,
		SquareE2: 0x0408102040800000,
		SquareF2: 0x0204081020408000,
		SquareG2: 0x0102040810204080,
		SquareH2: 0x0001020408102040,
		SquareA3: 0x2040800000000000,
		SquareB3: 0x1020408000000000,
		SquareC3: 0x0810204080000000,
		SquareD3: 0x0408102040800000,
		SquareE3: 0x0204081020408000,
		SquareF3: 0x0102040810204080,
		SquareG3: 0x0001020408102040,
		SquareH3: 0x0000010204081020,
		SquareA4: 0x1020408000000000,
		SquareB4: 0x0810204080000000,
		SquareC4: 0x0408102040800000,
		SquareD4: 0x0204081020408000,
		SquareE4: 0x0102040810204080,
		SquareF4: 0x0001020408102040,
		SquareG4: 0x0000010204081020,
		SquareH4: 0x0000000102040810,
		SquareA5: 0x0810204080000000,
		SquareB5: 0x0408102040800000,
		SquareC5: 0x0204081020408000,
		SquareD5: 0x0102040810204080,
		SquareE5: 0x0001020408102040,
		SquareF5: 0x0000010204081020,
		SquareG5: 0x0000000102040810,
		SquareH5: 0x0000000001020408,
		SquareA6: 0x0408102040800000,
		SquareB6: 0x0204081020408000,
		SquareC6: 0x0102040810204080,
		SquareD6: 0x0001020408102040,
		SquareE6: 0x0000010204081020,
		SquareF6: 0x0000000102040810,
		SquareG6: 0x0000000001020408,
		SquareH6: 0x0000000000010204,
		SquareA7: 0x0204081020408000,
		SquareB7: 0x0102040810204080,
		SquareC7: 0x0001020408102040,
		SquareD7: 0x0000010204081020,
		SquareE7: 0x0000000102040810,
		SquareF7: 0x0000000001020408,
		SquareG7: 0x0000000000010204,
		SquareH7: 0x0000000000000102,
		SquareA8: 0x0102040810204080,
		SquareB8: 0x0001020408102040,
		SquareC8: 0x0000010204081020,
		SquareD8: 0x0000000102040810,
		SquareE8: 0x0000000001020408,
		SquareF8: 0x0000000000010204,
		SquareG8: 0x0000000000000102,
		SquareH8: 0x0000000000000001,
	}

	// Note that the moves are raw, that is, for example, the piece moves can put him in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	roleDiagonalRawMoveDestBitboards = map[Square]Bitboard{
		SquareA1: 0x8040201008040201,
		SquareB1: 0x4020100804020100,
		SquareC1: 0x2010080402010000,
		SquareD1: 0x1008040201000000,
		SquareE1: 0x0804020100000000,
		SquareF1: 0x0402010000000000,
		SquareG1: 0x0201000000000000,
		SquareH1: 0x0100000000000000,
		SquareA2: 0x0080402010080402,
		SquareB2: 0x8040201008040201,
		SquareC2: 0x4020100804020100,
		SquareD2: 0x2010080402010000,
		SquareE2: 0x1008040201000000,
		SquareF2: 0x0804020100000000,
		SquareG2: 0x0402010000000000,
		SquareH2: 0x0201000000000000,
		SquareA3: 0x0000804020100804,
		SquareB3: 0x0080402010080402,
		SquareC3: 0x8040201008040201,
		SquareD3: 0x4020100804020100,
		SquareE3: 0x2010080402010000,
		SquareF3: 0x1008040201000000,
		SquareG3: 0x0804020100000000,
		SquareH3: 0x0402010000000000,
		SquareA4: 0x0000008040201008,
		SquareB4: 0x0000804020100804,
		SquareC4: 0x0080402010080402,
		SquareD4: 0x8040201008040201,
		SquareE4: 0x4020100804020100,
		SquareF4: 0x2010080402010000,
		SquareG4: 0x1008040201000000,
		SquareH4: 0x0804020100000000,
		SquareA5: 0x0000000080402010,
		SquareB5: 0x0000008040201008,
		SquareC5: 0x0000804020100804,
		SquareD5: 0x0080402010080402,
		SquareE5: 0x8040201008040201,
		SquareF5: 0x4020100804020100,
		SquareG5: 0x2010080402010000,
		SquareH5: 0x1008040201000000,
		SquareA6: 0x0000000000804020,
		SquareB6: 0x0000000080402010,
		SquareC6: 0x0000008040201008,
		SquareD6: 0x0000804020100804,
		SquareE6: 0x0080402010080402,
		SquareF6: 0x8040201008040201,
		SquareG6: 0x4020100804020100,
		SquareH6: 0x2010080402010000,
		SquareA7: 0x0000000000008040,
		SquareB7: 0x0000000000804020,
		SquareC7: 0x0000000080402010,
		SquareD7: 0x0000008040201008,
		SquareE7: 0x0000804020100804,
		SquareF7: 0x0080402010080402,
		SquareG7: 0x8040201008040201,
		SquareH7: 0x4020100804020100,
		SquareA8: 0x0000000000000080,
		SquareB8: 0x0000000000008040,
		SquareC8: 0x0000000000804020,
		SquareD8: 0x0000000080402010,
		SquareE8: 0x0000008040201008,
		SquareF8: 0x0000804020100804,
		SquareG8: 0x0080402010080402,
		SquareH8: 0x8040201008040201,
	}

	// Note that the moves are raw, that is, for example, the piece moves can put him in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	
	roleHorizontalRawMoveDestBitboards = map[Square]Bitboard{}

	// Note that the moves are raw, that is, for example, the king moves can put him in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	roleKingRawMoveDestBitboards = map[Square]Bitboard{
		SquareA1: 0x40c0000000000000,
		SquareB1: 0xa0e0000000000000,
		SquareC1: 0x5070000000000000,
		SquareD1: 0x2838000000000000,
		SquareE1: 0x141c000000000000,
		SquareF1: 0x0a0e000000000000,
		SquareG1: 0x0507000000000000,
		SquareH1: 0x0203000000000000,
		SquareA2: 0xc040c00000000000,
		SquareB2: 0xe0a0e00000000000,
		SquareC2: 0x7050700000000000,
		SquareD2: 0x3828380000000000,
		SquareE2: 0x1c141c0000000000,
		SquareF2: 0x0e0a0e0000000000,
		SquareG2: 0x0705070000000000,
		SquareH2: 0x0302030000000000,
		SquareA3: 0x00c040c000000000,
		SquareB3: 0x00e0a0e000000000,
		SquareC3: 0x0070507000000000,
		SquareD3: 0x0038283800000000,
		SquareE3: 0x001c141c00000000,
		SquareF3: 0x000e0a0e00000000,
		SquareG3: 0x0007050700000000,
		SquareH3: 0x0003020300000000,
		SquareA4: 0x0000c040c0000000,
		SquareB4: 0x0000e0a0e0000000,
		SquareC4: 0x0000705070000000,
		SquareD4: 0x0000382838000000,
		SquareE4: 0x00001c141c000000,
		SquareF4: 0x00000e0a0e000000,
		SquareG4: 0x0000070507000000,
		SquareH4: 0x0000030203000000,
		SquareA5: 0x000000c040c00000,
		SquareB5: 0x000000e0a0e00000,
		SquareC5: 0x0000007050700000,
		SquareD5: 0x0000003828380000,
		SquareE5: 0x0000001c141c0000,
		SquareF5: 0x0000000e0a0e0000,
		SquareG5: 0x0000000705070000,
		SquareH5: 0x0000000302030000,
		SquareA6: 0x00000000c040c000,
		SquareB6: 0x00000000e0a0e000,
		SquareC6: 0x0000000070507000,
		SquareD6: 0x0000000038283800,
		SquareE6: 0x000000001c141c00,
		SquareF6: 0x000000000e0a0e00,
		SquareG6: 0x0000000007050700,
		SquareH6: 0x0000000003020300,
		SquareA7: 0x0000000000c040c0,
		SquareB7: 0x0000000000e0a0e0,
		SquareC7: 0x0000000000705070,
		SquareD7: 0x0000000000382838,
		SquareE7: 0x00000000001c141c,
		SquareF7: 0x00000000000e0a0e,
		SquareG7: 0x0000000000070507,
		SquareH7: 0x0000000000030203,
		SquareA8: 0x000000000000c040,
		SquareB8: 0x000000000000e0a0,
		SquareC8: 0x0000000000007050,
		SquareD8: 0x0000000000003828,
		SquareE8: 0x0000000000001c14,
		SquareF8: 0x0000000000000e0a,
		SquareG8: 0x0000000000000705,
		SquareH8: 0x0000000000000302,
	}

	// Note that the moves are raw, that is, for example, the knight moves can put his king in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	//nolint:mnd // Magic numbers represents move Bitboards from specific square.
	roleKnightRawMoveDestBitboards = map[Square]Bitboard{
		SquareA1: 0x0020400000000000,
		SquareB1: 0x0010a00000000000,
		SquareC1: 0x0088500000000000,
		SquareD1: 0x0044280000000000,
		SquareE1: 0x0022140000000000,
		SquareF1: 0x00110a0000000000,
		SquareG1: 0x0008050000000000,
		SquareH1: 0x0004020000000000,
		SquareA2: 0x2000204000000000,
		SquareB2: 0x100010a000000000,
		SquareC2: 0x8800885000000000,
		SquareD2: 0x4400442800000000,
		SquareE2: 0x2200221400000000,
		SquareF2: 0x1100110a00000000,
		SquareG2: 0x0800080500000000,
		SquareH2: 0x0400040200000000,
		SquareA3: 0x4020002040000000,
		SquareB3: 0xa0100010a0000000,
		SquareC3: 0x5088008850000000,
		SquareD3: 0x2844004428000000,
		SquareE3: 0x1422002214000000,
		SquareF3: 0x0a1100110a000000,
		SquareG3: 0x0508000805000000,
		SquareH3: 0x0204000402000000,
		SquareA4: 0x0040200020400000,
		SquareB4: 0x00a0100010a00000,
		SquareC4: 0x0050880088500000,
		SquareD4: 0x0028440044280000,
		SquareE4: 0x0014220022140000,
		SquareF4: 0x000a1100110a0000,
		SquareG4: 0x0005080008050000,
		SquareH4: 0x0002040004020000,
		SquareA5: 0x0000402000204000,
		SquareB5: 0x0000a0100010a000,
		SquareC5: 0x0000508800885000,
		SquareD5: 0x0000284400442800,
		SquareE5: 0x0000142200221400,
		SquareF5: 0x00000a1100110a00,
		SquareG5: 0x0000050800080500,
		SquareH5: 0x0000020400040200,
		SquareA6: 0x0000004020002040,
		SquareB6: 0x000000a0100010a0,
		SquareC6: 0x0000005088008850,
		SquareD6: 0x0000002844004428,
		SquareE6: 0x0000001422002214,
		SquareF6: 0x0000000a1100110a,
		SquareG6: 0x0000000508000805,
		SquareH6: 0x0000000204000402,
		SquareA7: 0x0000000040200020,
		SquareB7: 0x00000000a0100010,
		SquareC7: 0x0000000050880088,
		SquareD7: 0x0000000028440044,
		SquareE7: 0x0000000014220022,
		SquareF7: 0x000000000a110011,
		SquareG7: 0x0000000005080008,
		SquareH7: 0x0000000002040004,
		SquareA8: 0x0000000000402000,
		SquareB8: 0x0000000000a01000,
		SquareC8: 0x0000000000508800,
		SquareD8: 0x0000000000284400,
		SquareE8: 0x0000000000142200,
		SquareF8: 0x00000000000a1100,
		SquareG8: 0x0000000000050800,
		SquareH8: 0x0000000000020400,
	}

	// Note that the moves are raw, that is, for example, the piece moves can put him in checkmate. Moreover, these
	// moves do not exclude a collision with a piece of their own color.
	//
	
	roleVerticalRawMoveDestBitboards = map[Square]Bitboard{}
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

// roleGetAntidiagonalRawMoveDestsBitboard returns a bitboard of all possible antidiagonal destinations from passed
// origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their in checkmate. Moreover, these
// moves do not exclude a collision with a piece of their own color.
func roleGetAntidiagonalRawMoveDestsBitboard(origin Square) (Bitboard, error) {
	bitboard, ok := roleAntidiagonalRawMoveDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("unknown origin")
	}
	return bitboard, nil
}

// roleGetDiagonalRawMoveDestsBitboard returns a bitboard of all possible diagonal destinations from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their in checkmate. Moreover, these
// moves do not exclude a collision with a piece of their own color.
func roleGetDiagonalRawMoveDestsBitboard(origin Square) (Bitboard, error) {
	bitboard, ok := roleDiagonalRawMoveDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("unknown origin")
	}
	return bitboard, nil
}

// roleGetHorizontalRawMoveDestsBitboard returns a bitboard of all possible horizontal destinations from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their in checkmate. Moreover, these
// moves do not exclude a collision with a piece of their own color.
func roleGetHorizontalRawMoveDestsBitboard(origin Square) (Bitboard, error) {
	bitboard, ok := roleHorizontalRawMoveDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("unknown origin")
	}
	return bitboard, nil
}

// roleGetKingRawMoveDestsBitboard returns a bitboard of all possible king destinations from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put him in checkmate. Moreover, these
// moves do not exclude a collision with a piece of their own color.
func roleGetKingRawMoveDestsBitboard(origin Square) (Bitboard, error) {
	bitboard, ok := roleKingRawMoveDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("unknown origin")
	}
	return bitboard, nil
}

// roleGetKnightRawMoveDestsBitboard returns a bitboard of all possible knight destinations from passed origin.
//
// Note that the moves are raw, that is, for example, the king moves can put him in checkmate. Moreover, these
// moves do not exclude a collision with a piece of their own color.
func roleGetKnightRawMoveDestsBitboard(origin Square) (Bitboard, error) {
	bitboard, ok := roleKnightRawMoveDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("unknown origin")
	}
	return bitboard, nil
}

// roleGetVerticalRawMoveDestsBitboard returns a bitboard of all possible vertical destinations from passed origin.
//
// Note that the moves are raw, that is, for example, the piece moves can put their in checkmate. Moreover, these
// moves do not exclude a collision with a piece of their own color.
func roleGetVerticalRawMoveDestsBitboard(origin Square) (Bitboard, error) {
	bitboard, ok := roleVerticalRawMoveDestBitboards[origin]
	if !ok {
		return BitboardNil, errors.New("unknown origin")
	}
	return bitboard, nil
}
