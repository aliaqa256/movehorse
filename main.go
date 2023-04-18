package main

import "fmt"

const (
	VERTICAL   = 0
	HORIZONTAL = 1
	SIZE = 6
)

var board [SIZE][SIZE]int

var horseMoves = [8][2]int{
	//    vertical, horizontal
	{2, 1},
	{2, -1},
	{-2, -1},
	{-2, 1},
	{-1, -2},
	{-1, 2},
	{1, -2},
	{1, 2},
}

func main() {
	horseMoveRecursive(0, 0, 1)
}

// horse move recursive
func horseMoveRecursive(row, col, move int) bool {
	if move == SIZE*SIZE+1{
		
		ShowBoard()
		return true
	}

	for i := 0; i < 8; i++ {
		newRow := row + horseMoves[i][VERTICAL]
		newCol := col + horseMoves[i][HORIZONTAL]

		if Canmove(newRow, newCol) {
			board[newRow][newCol] = move
			if horseMoveRecursive(newRow, newCol, move+1) {
				return true
			} else {
				board[newRow][newCol] = 0
			}
		}
	}
	return false
}

func Canmove(row, col int) bool {
	if row >= 0 && row < SIZE && col >= 0 && col < SIZE && board[row][col] == 0 {
		return true
	}
	return false
}


func ShowBoard(){
	for _, row := range board {
		fmt.Println("---------------------------------")
		for _, col := range row {
			if col <10{
				fmt.Printf(" %d | ",col)
			}else{

				fmt.Print(col," | ")
			}
		}
		fmt.Println()
	}
}