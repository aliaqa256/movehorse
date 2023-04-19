package main

import (
    "fmt"
    "sort"
)

func main() {
    n := 8
    board := make([][]int, n)
    for i := range board {
        board[i] = make([]int, n)
    }
    x, y := 5,5
    if solveKT(board, x, y, 1, n) {
        printBoard(board)
    } else {
        fmt.Println("Solution does not exist")
    }
}

func solveKT(board [][]int, x, y, moveNum, n int) bool {
    if moveNum == n*n+1 {
        return true
    }
    moves := getMoves(board, x, y, n)
    sort.Slice(moves, func(i, j int) bool {
        return len(getMoves(board, moves[i][0], moves[i][1], n)) < len(getMoves(board, moves[j][0], moves[j][1], n))
    })
    for _, move := range moves {
        i, j := move[0], move[1]
        board[i][j] = moveNum
        if solveKT(board, i, j, moveNum+1, n) {
            return true
        }
        board[i][j] = 0
    }
    return false
}

func getMoves(board [][]int, x, y, n int) [][2]int {
    moves := make([][2]int, 0, 8)
    for _, move := range [][2]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}} {
        i, j := x+move[0], y+move[1]
        if i >= 0 && i < n && j >= 0 && j < n && board[i][j] == 0 {
            moves = append(moves, [2]int{i, j})
        }
    }
    return moves
}

func printBoard(board [][]int) {
    for i := range board {
        for j := range board[i] {
            fmt.Printf("%3d ", board[i][j])
        }
        fmt.Println()
    }
}
