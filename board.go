package squares

import "fmt"

const (
	//SIZE is the size of the nxn board
	SIZE = 5
)

//Position represents the position of a move
type Position struct {
	y, x int
}

func (p Position) String() string {
	return fmt.Sprintf("y=%d,x=%d", p.y, p.x)
}

//Add takes 2 positions and performs a vector addition
func (p Position) Add(p2 Position) Position {
	return Position{p.y + p2.y, p.x + p2.x}
}

//Sub takes 2 positions and performs a vector subtraction
func (p Position) Sub(p2 Position) Position {
	return Position{p.y - p2.y, p.x - p2.x}
}

//InBounds checks if the position is on the board
func (p Position) InBounds() bool {
	return p.y >= 0 && p.y < SIZE && p.x >= 0 && p.x < SIZE
}

func (p Position) rotate() Position {
	return Position{p.x, -p.y}
}

//Piece represents the colored piece on the board
type Piece byte

const (
	NILPIECE Piece = iota
	RED
	GREEN
)

func (p Piece) String() string {
	switch p {
	case RED:
		return "r"
	case GREEN:
		return "g"
	default:
		return " "
	}
}

//Opponent returns the opposite piece.
//An empty piece's opponent is the empty piece.
func (p Piece) Opponent() Piece {
	switch p {
	case RED:
		return GREEN
	case GREEN:
		return RED
	default:
		return NILPIECE
	}
}

type board [SIZE][SIZE]Piece

func (b board) String() string {
	ret := ""
	for i := 0; i < SIZE; i++ {
		ret += "|"
		for j := 0; j < SIZE; j++ {
			ret += b[i][j].String() + "|"
		}
		ret += "\n"
	}
	return ret
}

func (b board) PieceAt(p Position) Piece {
	if !p.InBounds() {
		return NILPIECE
	}
	return b[p.y][p.x]
}

func (b board) lostFromPosition(p Position) bool {
	pieceType := b.PieceAt(p)
	pieces := make([]Position, 0)
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			pos := Position{i, j}
			if p != pos && b.PieceAt(pos) == pieceType {
				pieces = append(pieces, pos)
			}
		}
	}

pieceLoop:
	for _, pos := range pieces {
		delta := pos.Sub(p)
		check := pos

		for i := 0; i < 2; i++ {
			delta = delta.rotate()
			check = check.Add(delta)
			if b.PieceAt(check) != pieceType {
				continue pieceLoop
			}
		}

		//Only way to get here is if a square was made
		return true
	}

	return false
}

func (b board) draw() bool {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if b[i][j] == NILPIECE {
				return false
			}
		}
	}
	return true
}
