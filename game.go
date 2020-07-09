package squares

import (
	"errors"
	"fmt"
)

var (
	//ErrMoveInvalid is the base error for invalid moves.
	ErrMoveInvalid = errors.New("squares: move is invalid")
)

//Game is the game state
type Game struct {
	board
	current Piece

	//Keep a var holding the winning piece
	winner Piece
}

//NewGame returns a new game of Squares.
func NewGame() Game {
	return Game{current: RED}
}

//GetActions returns the list of y,x positions that can be made
//from this game state.
func (g Game) GetActions() []Position {
	moves := make([]Position, 0)
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if g.board[i][j] == NILPIECE {
				moves = append(moves, Position{i, j})
			}
		}
	}
	return moves
}

//ApplyAction generates a new game state by placing a piece p on the board
//If p is an invalid action, an error is returned with an unchanged board.
func (g Game) ApplyAction(p Position) (Game, error) {
	if g.PieceAt(p) != NILPIECE {
		return g, fmt.Errorf("%w: %s already contains a piece", ErrMoveInvalid, p)
	}

	g.board[p.y][p.x] = g.current
	g.current = g.current.Opponent()

	if g.lostFromPosition(p) {
		g.winner = g.current
	}

	return g, nil
}

//Winner returns the piece type of the winner
//Returns NILPIECE on draw and non-terminal states.
func (g Game) Winner() Piece {
	return g.winner
}

//IsTerminal returns true on a terminal state.
func (g Game) IsTerminal() bool {
	winner := g.Winner()
	if winner == NILPIECE {
		return g.draw()
	}
	return true
}

//Player returns the player that is currently considering
//a move.
func (g Game) Player() Piece {
	return g.current
}
