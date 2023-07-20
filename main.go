package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	board         [3][3]string
	currentPlayer string
	gameOver      bool
	winner        string
}

// inizializzo il gioco
func NewGame() *Game {
	game := &Game{}
	game.board = [3][3]string{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	game.currentPlayer = "X"
	game.gameOver = false
	game.winner = ""
	return game
}

// movimento dei giocatori
func (game *Game) makeMove(row, col int) error {
	if game.gameOver {
		return errors.New("The game ended")
	}

	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return errors.New("movement not valid")
	}

	if game.board[row][col] != " " {
		return errors.New("Already occupied cell")
	}

	game.board[row][col] = game.currentPlayer
	game.checkGameOver()

	if game.currentPlayer == "X" {
		game.currentPlayer = "O"
	} else {
		game.currentPlayer = "X"
	}

	return nil
}

// controllo della fine del gioco
func (game *Game) checkGameOver() {
	//controllo vittorie orizzontali
	for i := 0; i < 3; i++ {
		if game.board[i][0] == game.board[i][1] && game.board[i][0] == game.board[i][2] && game.board[i][0] != " " {
			game.gameOver = true
			game.winner = game.board[i][0]
			return
		}
	}

	//controllo vittorie verticali
	for i := 0; i < 3; i++ {
		if game.board[0][i] == game.board[1][i] && game.board[0][i] == game.board[2][i] && game.board[0][i] != " " {
			game.gameOver = true
			game.winner = game.board[0][i]
			return
		}
	}

	//controllo vittorie diagonali (da sx a dx)
	if game.board[0][0] == game.board[1][1] && game.board[0][0] == game.board[2][2] && game.board[0][0] != " " {
		game.gameOver = true
		game.winner = game.board[0][0]
		return
	}

	//controllo vittorie diagonali (da dx a dx)
	if game.board[0][2] == game.board[1][1] && game.board[0][2] == game.board[2][0] && game.board[0][2] != " " {
		game.gameOver = true
		game.winner = game.board[0][2]
		return
	}

	// controllo pareggio
	if game.isBoardFull() {
		game.gameOver = true
		game.winner = "Draw"
	}

}

func (game *Game) isBoardFull() bool {
	for _, row := range game.board {
		for _, cell := range row {
			if cell == " " {
				return false
			}
		}
	}
	return true
}

func main() {
	game := NewGame()
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("    0, 1, 2\n" +
		"0  [       ]\n" +
		"1  [       ]\n" +
		"2  [       ]\n")

	for !game.gameOver {
		fmt.Println("Current Player:", game.currentPlayer)
		fmt.Print("Enter the row number (0-2): ")
		rowStr, _ := reader.ReadString('\n')
		row, _ := strconv.Atoi(strings.TrimSpace(rowStr))

		fmt.Print("Enter the col number (0-2): ")
		colStr, _ := reader.ReadString('\n')
		col, _ := strconv.Atoi(strings.TrimSpace(colStr))

		err := game.makeMove(row, col)
		if err != nil {
			fmt.Println("Movement not valid:", err)
		}

		fmt.Println()

		//stampa la griglia di gioco
		for _, row := range game.board {
			fmt.Println(row)
		}
		fmt.Println()
	}

	if game.winner == "Draw" {
		fmt.Println("Draw!")
	} else {
		fmt.Println("The winner is: ", game.winner)
	}
}
