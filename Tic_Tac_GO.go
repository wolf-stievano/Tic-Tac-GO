package main

import (
	"fmt"
	"strings"
)

const (
	boardSize = 3
	empty     = " "
	reset     = "\033[0m"
	red	  = "\033[31m"
)

type Board [boardSize][boardSize]string
type MainBoard [3][3]*Board

func (b Board) Print() {
	for i, row := range b {
		fmt.Printf(" %s \n", strings.Join(row[:], " | "))
		if i < boardSize-1 {
		    fmt.Println(strings.Repeat("-", boardSize*3+1))
		} 
	}
}

func convertMove(move int) (int, int) {
	move--
	return move / 3, move % 3
}

func (b *Board) makeMove(move int, player string) bool {
	row, col := convertMove(move)
	if row < 0 || row >= boardSize || col < 0 || col >= boardSize || b[row][col] != empty{
		return false
	}
	b[row][col] = player
	return true
}

func (b Board) checkWin(player string) bool {
	for i := 0; i < boardSize; i++ {
		if b[i][0] == player && b[i][1] == player && b[i][2] == player {
			return true
		} 
	}

	for j := 0; j < boardSize; j++ {
		if b[0][j] == player && b[1][j] == player && b[2][j] == player {
			return true
		}
	}

	if b[0][0] == player && b[1][1] == player && b[2][2] == player {
		return true
	}

	if b[0][2] == player && b[1][1] == player && b[2][0] == player {
		return true
	}
	return false
}

func (b Board) isFull() bool {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b[i][j] == empty {
				return false
			}
		}
	}
	return true
}

func NewBoard() *Board {
	return &Board{
		{empty, empty, empty},
		{empty, empty, empty},
		{empty, empty, empty},	
	}
}

func NewMainBoard() MainBoard {
	var mb MainBoard
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			mb[i][j] = NewBoard()
		}
	}
	return mb
}

func (mb MainBoard) print(nextBoardRow, nextBoardCol int) {
	for i := 0; i < 3; i++ {
		for r := 0; r < 3; r++ {
			for j := 0; j < 3; j++ {
				if i == nextBoardRow && j == nextBoardCol {
					fmt.Printf(" %s%s%s |", red, strings.Join(mb[i][j][r][:], " | "), reset)
				} else {
					fmt.Printf(" %s |", strings.Join(mb[i][j][r][:], " | "))
				}
			}
			fmt.Println()
		}
		if i < 2 {
			fmt.Println(strings.Repeat("-", 35))
		}
	}
}
func (mb *MainBoard) MakeMove(mainMove, subMove int, player string) bool {
	mainRow, mainCol := convertMove(mainMove)
	return mb[mainRow][mainCol].makeMove(subMove, player)
}

func (mb MainBoard) CheckWin(player string) bool {
	for i := 0; i < 3; i++ {
		if mb[i][0].checkWin(player) && mb[i][1].checkWin(player) && mb[i][2].checkWin(player) {
			return true
		}
		if mb[0][i].checkWin(player) && mb[1][i].checkWin(player) && mb[2][i].checkWin(player) {
			return true
		}
	}
	if mb[0][0].checkWin(player) && mb[1][1].checkWin(player) && mb[2][2].checkWin(player) {
		return true
	}
	if mb[0][2].checkWin(player) && mb[1][1].checkWin(player) && mb[2][0].checkWin(player) {
		return true
	}
	return false
}

func (mb MainBoard) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !mb[i][j].isFull() {
				return false
			}
		}
	}
	return true
}

func main() {
	mb := NewMainBoard()
	player := "X"
	lastMove := -1

	for {
		var mainMove, subMove int
		nextBoardRow, nextBoardCol := -1, -1

		if lastMove != -1 && !mb[lastMove/3][lastMove%3].isFull() {
			mainMove = lastMove + 1
			nextBoardRow, nextBoardCol = lastMove/3, lastMove%3
			fmt.Printf("Jogador %s, você deve jogar no mini-tabuleiro %d\n", player, mainMove)
		} else {
			fmt.Printf("Jogador %s, escolha um mini-tabuleiro (1-9): ", player)
			fmt.Scan(&mainMove)
			if mainMove < 1 || mainMove > 9 {
				fmt.Println("Mini-tabuleiro inválido. Por favor, escolha um número entre 1 e 9.")
				continue
			}
			nextBoardRow, nextBoardCol = convertMove(mainMove)
		}

		mb.print(nextBoardRow, nextBoardCol)

		fmt.Printf("Jogador %s, faça sua jogada no mini-tabuleiro %d (1-9): ", player, mainMove)
		fmt.Scan(&subMove)
		if subMove < 1 || subMove > 9 {
			fmt.Println("Jogada inválida. Por favor, escolha um número entre 1 e 9.")
			continue
		}

		if !mb.MakeMove(mainMove, subMove, player) {
			fmt.Println("Posição já ocupada ou inválida. Tente novamente.")
			continue
		}

		lastMove = subMove - 1

		if mb.CheckWin(player) {
			mb.print(nextBoardRow, nextBoardCol)
			fmt.Printf("Jogador %s venceu o jogo principal!\n", player)
			break
		}

		if mb.IsFull() {
			mb.print(nextBoardRow, nextBoardCol)
			fmt.Println("O jogo principal empatou!")
			break
		}

		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}
}

