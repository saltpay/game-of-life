package game_test

import (
	"github.com/matryer/is"
	"github.com/saltpay/game-of-life/game"
	"testing"
)

func TestMakingABoardFromAString(t *testing.T) {
	is := is.New(t)
	input := `
**
*-
`

	board := game.NewBoardFromString(input)

	t.Log(board)
	is.True(board.IsCellAlive(0, 0))
	is.True(board.IsCellAlive(1, 0))
	is.True(board.IsCellAlive(0, 1))
	is.True(board.IsCellDead(1, 1))
}

func TestHappyGeneration(t *testing.T) {
	is := is.New(t)
	input := `
**
*-
`

	board := game.NewBoardFromString(input)
	board.NextGeneration()
	wantBoard := game.NewBoardFromString(`
**
*-
`)
	is.Equal(board, wantBoard)
}
