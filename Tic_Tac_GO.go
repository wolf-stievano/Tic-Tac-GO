package main

import (
	"fmt"
	"strings"
)

const (
	boardSize = 3
	empty     = " "
)

type Board [boardSize][boardSize]string

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

func main() {
	board := Board{
		{empty, empty, empty},
		{empty, empty, empty},
		{empty, empty, empty},
	}

	player := "X"
	for {
		board.Print()
		fmt.Printf("Jogador %s, faça sua jogada (1-9): ", player)
		var move int
		fmt.Scan(&move)

		if move < 1 || move > 9 {
			fmt.Println("Jogada inválida. Por favor, escolha um número entre 1 e 9.")
			continue
		}

		if !board.makeMove(move, player) {
			fmt.Println("Posição já ocupada ou inválida. tente novamente.")
			continue
		}
		
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}
	}

}

// nao quero fazer isto daqui
