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
	for _, tc := range []struct{
		description string
		wantBoard     string
		startingBoard string
	}{
	{
		description: "Generation with 3 neighbour cell",
		startingBoard: `
**
*-
`,
		wantBoard: `
**
**
`,
	},
	{
		description: "Generation with 2 neighbour cell",

		startingBoard:`
**
--
` ,
		wantBoard: `
--
--
`,
	},
		{
			description: "Cell with no neighbours dies",
			startingBoard:`
*-
--
` ,
			wantBoard: `
--
--
`,
		},//the dead cell with three neighbours isnt getting set alive. Im pretty sure its cause we need a copy. gonna test it out
		{
			description: "4x4 generation with 3 neighbours",
			startingBoard:`
**--
*---
----
----
` ,
			wantBoard: `
**--
**--
----
----
`,
		},
}{
	t.Run(tc.description, func(t *testing.T) {
		gotBoard := game.NewBoardFromString(tc.startingBoard)
		newBoard := gotBoard.NextGeneration()
		is.Equal(newBoard.String(), tc.wantBoard)
	})
}
}