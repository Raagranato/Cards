package tictactoe

import (
	"fmt"
)

type tictactoe struct {
	matrix [3][3]string //" " ou X ou O
	player string       //X or O
}

func (t *tictactoe) PrintBoard() {
	fmt.Printf("%s|%s|%s\n", t.matrix[0][0], t.matrix[0][1], t.matrix[0][2])
	fmt.Printf("%s|%s|%s\n", t.matrix[1][0], t.matrix[1][1], t.matrix[1][2])
	fmt.Printf("%s|%s|%s\n", t.matrix[2][0], t.matrix[2][1], t.matrix[2][2])
}
func (t *tictactoe) GameLoop() {
	for t.IsThereSpace() {
		x, y := getInput()
		for t.matrix[x][y] != " " {
			println("Not possible to make this move! Try other:\n")
			x, y = getInput()
		}
		if won, name := t.SomeoneWon(); won {
			fmt.Printf("%s ganhou!\n", name)
			return
        }
        t.switchPlayer()
	}

}
func (t *tictactoe) IsThereSpace() bool {
	for _, row := range t.matrix {
		for _, value := range row {
			if value == " " {
				return true
			}
		}
	}
	return false

}

func (t *tictactoe) SomeoneWon() (bool, string) {

	// Linhas Horizontais (O)
	if t.matrix[0][0] == "O" && t.matrix[0][1] == "O" && t.matrix[0][2] == "O" {
		return true, "O"
	} else if t.matrix[1][0] == "O" && t.matrix[1][1] == "O" && t.matrix[1][2] == "O" {
		return true, "O"
	} else if t.matrix[2][0] == "O" && t.matrix[2][1] == "O" && t.matrix[2][2] == "O" {
		return true, "O"
	}

	// Colunas Verticais (O)
	if t.matrix[0][0] == "O" && t.matrix[1][0] == "O" && t.matrix[2][0] == "O" {
		return true, "O"
	} else if t.matrix[0][1] == "O" && t.matrix[1][1] == "O" && t.matrix[2][1] == "O" {
		return true, "O"
	} else if t.matrix[0][2] == "O" && t.matrix[1][2] == "O" && t.matrix[2][2] == "O" {
		return true, "O"
	}

	// Diagonais (O)
	if t.matrix[0][0] == "O" && t.matrix[1][1] == "O" && t.matrix[2][2] == "O" {
		return true, "O"
	} else if t.matrix[0][2] == "O" && t.matrix[1][1] == "O" && t.matrix[2][0] == "O" {
		return true, "O"
	}

	// Linhas Horizontais (X)
	if t.matrix[0][0] == "X" && t.matrix[0][1] == "X" && t.matrix[0][2] == "X" {
		return true, "X"
	} else if t.matrix[1][0] == "X" && t.matrix[1][1] == "X" && t.matrix[1][2] == "X" {
		return true, "X"
	} else if t.matrix[2][0] == "X" && t.matrix[2][1] == "X" && t.matrix[2][2] == "X" {
		return true, "X"
	}

	// Colunas Verticais (X)
	if t.matrix[0][0] == "X" && t.matrix[1][0] == "X" && t.matrix[2][0] == "X" {
		return true, "X"
	} else if t.matrix[0][1] == "X" && t.matrix[1][1] == "X" && t.matrix[2][1] == "X" {
		return true, "X"
	} else if t.matrix[0][2] == "X" && t.matrix[1][2] == "X" && t.matrix[2][2] == "X" {
		return true, "X"
	}

	// Diagonais (X)
	if t.matrix[0][0] == "X" && t.matrix[1][1] == "X" && t.matrix[2][2] == "X" {
		return true, "X"
	} else if t.matrix[0][2] == "X" && t.matrix[1][1] == "X" && t.matrix[2][0] == "X" {
		return true, "X"
	}

	// no winners
	return false, " "
}

func (t *tictactoe) switchPlayer() {
	if t.player == "X" {
		t.player = "O"
	} else {
		t.player = "X"
	}
}
func getInput() (int, int) {
	var x, y int
	fmt.Print("What is your move?[x] [y] (ex:1 2): ")
	_, err := fmt.Scanln(&x, &y)
	if err != nil {
		fmt.Println("Not an input, enter 2 numbers separeted by space")
		return getInput()
	}

	return x + 1, y + 1
}
