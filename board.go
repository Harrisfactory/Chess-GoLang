package main

import "fmt"

type Board struct {
	game_board [7][7]Piece
}

func (b *Board) Init() {
	// Initialize the board with empty pieces
	for i := range b.game_board {
		for j := range b.game_board[i] {
			b.game_board[i][j] = Piece{Type: Empty}
		}
	}

	// Add a Pawn to the board
	b.game_board[0][1] = Piece{Type: Pawn}

	// Print the board
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print("| " + string(b.game_board[i][j].Type) + " |")
		}
		fmt.Println()
	}
}
