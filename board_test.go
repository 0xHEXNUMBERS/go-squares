package squares

import "fmt"

func ExampleTrivialSquare() {
	g := Game{
		board: board{
			[SIZE]Piece{RED, 0, RED, 0, 0},
			[SIZE]Piece{0, 0, 0, 0, 0},
			[SIZE]Piece{RED, 0, RED, 0, 0},
			[SIZE]Piece{0, 0, 0, 0, 0},
		},
		current: RED,
		winner:  NILPIECE,
	}

	fmt.Println(g.lostFromPosition(Position{0, 0}))
	fmt.Println(g.lostFromPosition(Position{0, 2}))
	fmt.Println(g.lostFromPosition(Position{2, 0}))
	fmt.Println(g.lostFromPosition(Position{2, 2}))
	//Output:
	//true
	//true
	//true
	//true
}

func ExampleNonTrivialSquare() {
	g := Game{
		board: board{
			[SIZE]Piece{0, RED, 0, 0, 0},
			[SIZE]Piece{0, 0, 0, RED, 0},
			[SIZE]Piece{RED, 0, 0, 0, 0},
			[SIZE]Piece{0, 0, RED, 0, 0},
		},
		current: RED,
		winner:  NILPIECE,
	}

	fmt.Println(g.lostFromPosition(Position{0, 1}))
	fmt.Println(g.lostFromPosition(Position{1, 3}))
	fmt.Println(g.lostFromPosition(Position{2, 0}))
	fmt.Println(g.lostFromPosition(Position{3, 2}))
	//Output:
	//true
	//true
	//true
	//true
}

func ExampleNonSquare() {
	g := Game{
		board: board{
			[SIZE]Piece{0, 0, RED, 0, 0},
			[SIZE]Piece{0, 0, 0, RED, 0},
			[SIZE]Piece{RED, 0, 0, 0, 0},
			[SIZE]Piece{0, 0, RED, 0, 0},
		},
		current: RED,
		winner:  NILPIECE,
	}

	fmt.Println(g.lostFromPosition(Position{0, 2}))
	fmt.Println(g.lostFromPosition(Position{1, 3}))
	fmt.Println(g.lostFromPosition(Position{2, 0}))
	fmt.Println(g.lostFromPosition(Position{3, 2}))
	//Output:
	//false
	//false
	//false
	//false
}
