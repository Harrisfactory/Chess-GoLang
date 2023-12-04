package main

import "fmt"

func main() {
	fmt.Println("This is a test")
	play()
	//move piece

}

// game loop
func play() {

	fmt.Println("Do you want to play chess? y/n")
	var userInput string
	fmt.Scanln(&userInput)

	if userInput != "y" {
		return
	}

	//initialize a game
	var currentGame = true
	//initialize a board
	board := Board{}
	board.Init()
	for currentGame {
		//make your move
		//board.movePiece()
		//print board
	}
}
