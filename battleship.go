package main

import (
	"flag"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func main() {
	log.SetPrefix("battleship: ")
	log.SetFlags(0)

	width := flag.Int("w", 10, "width of game board")
	height := flag.Int("h", 10, "height of game board")
	flag.Parse()

	board := InitializeBoard(*width, *height)
	PrintBoard(board)
}

func InitializeBoard(width int, height int) [][]string {
	log.Printf("build board with dimensions %v x %v\n", width, height)
	board := make([][]string, width)
	for i := range board {
		board[i] = make([]string, height)
		for j := range board[i] {
			board[i][j] = "?"
		}
	}
	return board
}

func PrintBoard(board [][]string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAutoIndex(true)
	for i := range board {
		row := table.Row{}
		for j := range board[i] {
			row = append(row, board[i][j])
		}
		t.AppendRow(row)
	}
	t.SetStyle(table.StyleColoredDark)
	t.Render()
}
