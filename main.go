package main

import "fmt"

const (
	SIZE = 8
)

var sol [SIZE][SIZE]int

func main() {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			sol[i][j] = -1
		}
	}

	x_move := [8]int{2, 1, -1, -2, -2, -1, 1, 2}
	y_move := [8]int{1, 2, 2, 1, -1, -2, -2, -1}

	sol[1][1] = 0

	if knight_tour(1, 1, 1, x_move, y_move) {
		for _, row := range sol {
			fmt.Println("---------------------------------")
			for _, col := range row {
				if col < 10 {
					fmt.Printf(" %d | ", col)
				} else {
					fmt.Print(col, " | ")
				}
			}
			fmt.Println()
		}
	}
}

func isValid(i, j int) bool {
	return (i >= 0 && i < SIZE && j >= 0 && j < SIZE && sol[i][j] == -1)
}

func knight_tour(x, y, move int, x_move, y_move [8]int) bool {
	if move == SIZE*SIZE {
		return true
	}

	for k := 0; k < 8; k++ {
		next_x := x + x_move[k]
		next_y := y + y_move[k]

		if isValid(next_x, next_y) {
			sol[next_x][next_y] = move
			if knight_tour(next_x, next_y, move+1, x_move, y_move) {
				return true
			}
			sol[next_x][next_y] = -1

		}
	}

	return false
}
