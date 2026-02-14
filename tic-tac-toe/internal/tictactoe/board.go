package tictactoe

import (
	"errors"
	"fmt"
)

type Board struct {
	Grid       [3][3]rune
	MovesCount int
	WinPattern [][]int
}

var winpat [][]int= [][]int{
	{0, 1, 2}, {0, 3, 6}, {0, 4, 8},
	{1, 4, 7}, {2, 5, 8}, {2, 4, 6},
	{3, 4, 5}, {6, 7, 8},
}

func NewBoard() *Board {
	board := &Board{}
	board.InitializeBoard()
	return board
}

func (b *Board) InitializeBoard() {
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			b.Grid[row][col] = '-'
		}
	}
	b.MovesCount = 0
	b.WinPattern = winpat
}

func (b *Board) IsFull() bool {
	return b.MovesCount == 9
}

func (b *Board) MakeMove(row, col int, symbol rune) error {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return errors.New("sorry out of area")
	}
	if b.Grid[row][col] != '-' {
		return errors.New("sorry space already taken")
	}
	b.Grid[row][col] = symbol
	b.MovesCount++
	return nil
}

func (b *Board) HasWinner() bool {
	return false
}

func (b *Board)PrintBoard() {
	for i:=0;i<3;i++ {
		for j:=0;j<3;j++ {
			fmt.Print(string(b.Grid[i][j])+" ")
		}
		fmt.Println()
	}
}