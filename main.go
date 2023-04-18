package main

import "fmt"

const N int = 9

func isValid(i, j int, sol [][]int) bool {
	if i >= 1 && i <= N && j >= 1 && j <= N {
		if sol[i][j] == -1 {
			return true
		}
	}
	return false
}

func knightTour(sol [][]int, i, j, stepCount int, xMove, yMove []int) bool {
	if stepCount == N*N {
		return true
	}

	for k := 0; k < 8; k++ {
		nextI := i + xMove[k]
		nextJ := j + yMove[k]

		if isValid(nextI, nextJ, sol) {
			sol[nextI][nextJ] = stepCount
			if knightTour(sol, nextI, nextJ, stepCount+1, xMove, yMove) {
				return true
			}
			sol[nextI][nextJ] = -1 // backtracking
		}
	}

	return false
}

func startKnightTour() bool {
	sol := make([][]int, N+1)
	for i := range sol {
		sol[i] = make([]int, N+1)
		for j := range sol[i] {
			sol[i][j] = -1
		}
	}

	xMove := []int{2, 1, -1, -2, -2, -1, 1, 2}
	yMove := []int{1, 2, 2, 1, -1, -2, -2, -1}

	sol[1][1] = 0 // placing knight at cell(1, 1)

	if knightTour(sol, 1, 1, 1, xMove, yMove) {
		for i := 1; i <= N; i++ {
			fmt.Println(sol[i][1:])
		}
		return true
	}
	return false
}

func main() {
	fmt.Println(startKnightTour())
}
