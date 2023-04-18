package main

import "fmt"

const (
	SIZE       = 4
	VERTICAL   = 0
	HORIZONTAL = 1
)

var chessBoard [SIZE][SIZE]int

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
	MoveTheHorse(0, 0, 0)
	PrintBoard()
}

func MoveTheHorse(row, col, moveNumber int)  {
	if moveNumber == 0 {
		chessBoard[row][col] = 99
		moveNumber++
		
	}



	if moveNumber == SIZE*SIZE -1 {
		return 
	}
	for _, mm := range horseMoves {
		if CanMove(row, col, mm) {
			MoveHores(row, col, mm, moveNumber)
			MoveTheHorse(row+mm[VERTICAL], col+mm[HORIZONTAL], moveNumber+1) 
			chessBoard[row+mm[VERTICAL]][col+mm[HORIZONTAL]] = moveNumber 
		}
		

	}

	
}

func CanMove(row, col int, movement [2]int) bool {
	if row+movement[VERTICAL] < 0 || row+movement[VERTICAL] >= SIZE {
		return false
	}
	if col+movement[HORIZONTAL] < 0 || col+movement[HORIZONTAL] >= SIZE {
		return false
	}
	if chessBoard[row+movement[VERTICAL]][col+movement[HORIZONTAL]] != 0 {
		return false
	}
	return true
}

func MoveHores(row, col int, movement [2]int, sign int) {
	chessBoard[row+movement[VERTICAL]][col+movement[HORIZONTAL]] = sign
}

func PrintBoard() {
	for _, row := range chessBoard {
		fmt.Println("-------------------")
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


func IsOutOfIndex(row, col int, movement [2]int) bool{
	if row+movement[VERTICAL] < 0 || row+movement[VERTICAL] >= SIZE {
		return false
	}
	if col+movement[HORIZONTAL] < 0 || col+movement[HORIZONTAL] >= SIZE {
		return false
	}

	return true
}