package game

import (
	"strings"
)

const (
	Dead  string = "-️"
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
	board := NewBoard(2, 2)
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

func (b *Board) NextGeneration()  {

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
