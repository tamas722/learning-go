package multidimchess

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

type Position struct {
	Coordinates []int
}

type Piece struct {
	Type     string // e.g. Queen, King, Bishop, etc.
	Color    string // White or Black
	Position Position
}

type Board struct {
	Pieces []Piece
}

// INSERT YOUR CODE HERE

func copyPosition(p Position) Position {
	if p.Coordinates == nil {
		return Position{}
	}
	coords := make([]int, len(p.Coordinates))
	copy(coords, p.Coordinates)
	return Position{
		Coordinates: coords,
	}
}

func copyPiece(p Piece) Piece {
	return Piece{
		Type:     p.Type,
		Color:    p.Color,
		Position: copyPosition(p.Position),
	}
}

func copyBoard(b Board) Board {
	if b.Pieces == nil {
		return Board{}
	}
	pieces := make([]Piece, len(b.Pieces))
	for i, piece := range b.Pieces {
		pieces[i] = copyPiece(piece)
	}
	return Board{
		Pieces: pieces,
	}
}
