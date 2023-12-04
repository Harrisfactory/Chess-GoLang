package main

import "fmt"

type Board struct {
	gameBoard [8][8]Piece
}

func (b *Board) Init() {
	// Initialize the board with empty pieces
	for i := range b.gameBoard {
		for j := range b.gameBoard[i] {
			es := EmptySpace{}
			es.Init()
			b.gameBoard[i][j] = es.Piece
			if i == 1 || i == 6 {
				p := Pawn{} //declare struct
				//initialize type
				if i == 1 {
					p.Init("b")
				} else {
					p.Init("w")
				}
				b.gameBoard[i][j] = p.Piece //assign piece to board position
			}
		}
	}

	// Add a Pawn to the board
	//b.game_board[0][1] = Piece{Type: Pawn}

	// Print the board
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == 0 {
				fmt.Print(" " + b.gameBoard[i][j].GetType() + " ")
			} else {
				fmt.Print(" | " + b.gameBoard[i][j].GetType() + " ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) movePiece() {

}
