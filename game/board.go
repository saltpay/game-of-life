package game

import (
	"fmt"
	"strings"
)

const (
	Dead  string = "-"
	Alive string = "*"
)

type Board struct {
	cells [][]bool
	width int
	height int
}

func NewBoard(width, height int) *Board {
	cells := make([][]bool, width)
	for i := range cells {
		cells[i] = make([]bool, height)
	}
	return &Board{cells, width, height}
}

func NewBoardFromString(input string) *Board {
	input = strings.Trim(input, "\n")
	rows := strings.Split(input, "\n")
	board := NewBoard(len(rows[0]), len(rows))
	for y, row := range rows {
		cells := strings.Split(row, "")
		for x, cell := range cells {
			if cell == Alive {
				board.SetAlive(x, y)
			}
			if cell == Dead{
				board.SetDead(x, y)
			}
		}
	}
	return board
}

func (b Board) NextGeneration() Board {
	for y, rows := range b.cells {
		for x, _ := range rows {

			aliveNeighbours := GetAliveNeighbours(x, y, b)

			fmt.Printf("%t %t", b.cells[y][x], b.IsCellAlive(x, y))

			if b.IsCellAlive(x, y) && aliveNeighbours < 2 {
				b.SetDead(x, y)
			}

			if b.IsCellAlive(x, y) && aliveNeighbours > 3 {
				b.SetDead(x, y)
			}

			if b.IsCellAlive(x, y) && (aliveNeighbours == 2 || aliveNeighbours == 3) {

			}

			if b.IsCellDead(x,y) && aliveNeighbours == 3 {
				b.SetAlive(x, y)
			}
		}
	}

	return b
}

func GetAliveNeighbours(x,y int, b Board) int {
	offset := []struct {
		y int
		x int
	}{
		{
			-1,
			0,
		},
		{
			1,
			0,
		},
		{
			-1,
			1,
		},
		{
			1,
			-1,
		},
		{
			0,
			-1,
		},
		{
			0,
			1,
		},
	}

	aliveNeighbours := 0
	for _, o := range offset {
		row := y + o.y
		col := x + o.x
		if row > -1 && row < b.height && col > -1 && col < b.width {
			if b.cells[row][col] {
				aliveNeighbours++
			}
		}
	}
	return aliveNeighbours
}

func (b *Board) SetAlive(x int, y int) {
	b.cells[y][x] = true
}

func (b *Board) SetDead(x int, y int) {
	b.cells[y][x] = false
}

func (b Board) IsCellAlive(x int, y int) bool {
	return b.cells[y][x]
}

func (b Board) IsCellDead(x int, y int) bool {
	return !b.cells[y][x]
}

func (b Board) String() string {
	result := "\n"
	for _, col := range b.cells {
		for _, cell := range col{
			if cell{
				result += Alive
			} else {
				result += Dead
			}
		}
		result += "\n"
	}
	return result
}

func (b Board) CoolString() string {
	result := "\n"
	for _, col := range b.cells {
		for _, cell := range col{
			if cell{
				result += "⬛"
			} else {
				result += "⬜"
			}
		}
		result += "\n"
	}
	return result
}
