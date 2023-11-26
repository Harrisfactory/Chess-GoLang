package main

type PieceType string

const (
	Empty PieceType = "Empty"
	Pawn  PieceType = "Pawn"
	// Add other piece types as needed
)

type Piece struct {
	Type PieceType
}
