package tictactoe

func Play() {
	game := tictactoe{
    matrix: [3][3]string{
        {" ", " ", " "},
        {" ", " ", " "},
        {" ", " ", " "},
    },
    player: "X",
}
	game.GameLoop()
}
