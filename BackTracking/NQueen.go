//BACKTRACKING IS A ALGORITHMIC TECHNIQUE
//SOLVING PROBLEM RECURSIVELY
//TRY TO BUILD SOLUTION INCREMENTALLY
//REMOVE THOSE SOLUTION THAT FAIL TO SATISFY THE CONSTRAINT OF THE PROBLEM

package main

import (
	"fmt"
	"strconv"
)

func main() {
	board := [4][4]bool{}
	//(matrix,row , column, total_queen,currently_placed_queen, ans)
	Nqueen(board, 0, 0, 4, 0, "")
}
func Nqueen(board [4][4]bool, row int, col int, tq int, cpq int, ans string) {

	if cpq == tq {
		fmt.Println(ans)
		return
	}
	if col == len(board[0]) {
		row += 1
		col = 0
	}
	if row == len(board[0]) {
		return
	}
	if safe_for_place(board, row, col) {
		//do
		board[row][col] = true
		//recur
		Nqueen(board, row, col+1, tq, cpq+1, ans+"["+strconv.Itoa(row)+"]"+"["+strconv.Itoa(col)+"]"+" ")
		//undo
		board[row][col] = false
	}

	Nqueen(board, row, col+1, tq, cpq, ans)

}

func safe_for_place(board [4][4]bool, row int, col int) bool {
	//vertical up
	r := row - 1
	c := col
	for r >= 0 {
		if board[r][c] == true {
			return false
		}
		r -= 1
	}
	// horizontal left
	r = row
	c = col - 1
	for c >= 0 {
		if board[r][c] == true {
			return false
		}
		c -= 1
	}
	// digonal left
	r = row - 1
	c = col - 1
	for c >= 0 && r >= 0 {
		if board[r][c] == true {
			return false
		}
		c -= 1
		r -= 1
	}

	// digonal right
	r = row - 1
	c = col + 1
	for c < len(board[0]) && r >= 0 {
		if board[r][c] == true {
			return false
		}
		c += 1
		r -= 1
	}
	return true

}
