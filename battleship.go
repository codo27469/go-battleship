package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/table"
)

func main() {
	log.SetPrefix("battleship: ")
	log.SetFlags(0)

	width := flag.Int("w", 10, "width of game board")
	height := flag.Int("h", 10, "height of game board")
	flag.Parse()

	board := InitializeBoard(*width, *height)
	s1, _ := CreateShip("A1", "A5")
	s2, _ := CreateShip("B3", "B5")
	s3, _ := CreateShip("D9", "J9")
	PlaceShip(s1, board)
	PlaceShip(s2, board)
	PlaceShip(s3, board)
	PrintBoard(board)
}

type BoardPiece struct {
	Text    string
	HasShip bool
}

type Ship struct {
	Col1 int
	Col2 int
	Row1 int
	Row2 int
	Name string
}

func CreateShip(p1 string, p2 string) (Ship, error) {
	if p1[0] != p2[0] && p1[1] != p2[1] {
		return Ship{}, errors.New("invalid points for ship placement")
	}
	col1 := int(p1[0]) - 65
	col2 := int(p2[0]) - 65
	row1, _ := strconv.Atoi(string(p1[1]))
	row2, _ := strconv.Atoi(string(p2[1]))
	row1 -= 1
	row2 -= 1
	ship := Ship{col1, col2, row1, row2, fmt.Sprintf("%v-%v", p1, p2)}
	return ship, nil
}

func InitializeBoard(width int, height int) [][]BoardPiece {
	log.Printf("build board with dimensions %v x %v\n", width, height)
	board := make([][]BoardPiece, height)
	for i := range board {
		board[i] = make([]BoardPiece, width)
		for j := range board[i] {
			board[i][j] = BoardPiece{"?", false}
		}
	}
	return board
}

func PlaceShip(ship Ship, board [][]BoardPiece) ([][]BoardPiece, error) {
	if ship.Col1 == ship.Col2 {
		// ship is placed vertically
		for i := range board {
			for j := range board[i] {
				if j == ship.Col1 && (i >= ship.Row1 && i <= ship.Row2) {
					board[i][j].HasShip = true
				}
			}
		}
	} else if ship.Row2 == ship.Row2 {
		// ship is place horizontally
		for i := range board {
			for j := range board[i] {
				if i == ship.Row1 && (j >= ship.Col1 && j <= ship.Col2) {
					board[i][j].HasShip = true
				}
			}
		}
	}
	return board, nil
}

func PrintBoard(board [][]BoardPiece) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetAutoIndex(true)
	for i := range board {
		row := table.Row{}
		for j := range board[i] {
			row = append(row, board[i][j].Text)
		}
		t.AppendRow(row)
	}
	t.SetStyle(table.StyleColoredDark)
	t.Render()
}
