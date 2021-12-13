package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Boarder interface {
	won()
	mark_board(drew_number int)
	sum_unmarked()
	final_score(last_number_drew int)
}

type Board struct {
	board        [][]int
	marked_board [5][5]bool
}

func NewBoard() *Board {
	var marked_board [5][5]bool

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			marked_board[i][j] = false
		}
	}
	return &Board{marked_board: marked_board}
}

func (x *Board) mark_board(drew_number int) {
	len_line := len(x.marked_board)
	for i := 0; i < len_line; i++ {
		for j := 0; j < len_line; j++ {
			if x.board[i][j] == drew_number {
				x.marked_board[i][j] = true
			}
		}
	}
}

func (x Board) sum_unmarked() int {
	sum_of_unmarked_cells := 0
	len_line := len(x.marked_board)

	for i := 0; i < len_line; i++ {
		for j := 0; j < len_line; j++ {
			if !x.marked_board[i][j] {
				sum_of_unmarked_cells += x.board[i][j]
			}
		}
	}
	return sum_of_unmarked_cells
}

func (x Board) final_score(last_number_drew int) int {
	return last_number_drew * x.sum_unmarked()
}

func (x Board) won() bool {
	for _, line := range x.marked_board {
		if all_true(line) {
			return true
		}
	}

	len_line := len(x.marked_board)
	for i := 0; i < len_line; i++ {
		won := true
		for j := 0; j < len_line; j++ {
			won = won && x.marked_board[j][i]
		}
		if won {
			return true
		}
	}
	return false
}

func all_true(line [5]bool) bool {
	for _, value := range line {
		if !value {
			return false
		}
	}
	return true
}

func create_drew_numbers_array(scanner *bufio.Scanner) []int {
	scanner.Scan()
	drew_numbers_str := strings.Split(scanner.Text(), ",")
	return array_str_to_int(drew_numbers_str)
}

func array_str_to_int(array []string) []int {
	array_int := make([]int, 0)
	for _, value := range array {
		value_int, _ := strconv.Atoi(value)
		array_int = append(array_int, value_int)
	}
	return array_int
}

func add_board(new_board [][]int, boards []Board) []Board {
	var board Board = *NewBoard()
	board.board = new_board
	return append(boards, board)
}

func play(boards []Board, drew_numbers []int) int {
	for _, drew := range drew_numbers {
		for i := 0; i < len(boards); i++ {
			boards[i].mark_board(drew)
			if boards[i].won() {
				return boards[i].final_score(drew)
			}
		}
	}
	return -1
}

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var boards []Board

	drew_numbers := create_drew_numbers_array(scanner)
	scanner.Scan()

	var new_board [][]int

	for scanner.Scan() {
		current_read := scanner.Text()

		if current_read != "" {
			current_read_ints := array_str_to_int(strings.Split(strings.ReplaceAll(strings.Trim(current_read, " "), "  ", " "), " "))
			new_board = append(new_board, current_read_ints)
		} else {
			boards = add_board(new_board, boards)

			new_board = nil
		}
	}
	boards = add_board(new_board, boards)

	score := play(boards, drew_numbers)

	fmt.Println(score)
}
