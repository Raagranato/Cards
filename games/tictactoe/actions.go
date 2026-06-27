package tictactoe

import (
	"fmt"
	"net"
	"FeelGoodInc/network"
)

type tictactoe struct {
	matrix [3][3]string //" " ou X ou O
	player string       //X or O
}

func (t *tictactoe) PrintBoard() {
	fmt.Printf("%s | %s | %s\n", t.matrix[0][0], t.matrix[0][1], t.matrix[0][2])
	fmt.Println("──┼───┼──")
	fmt.Printf("%s | %s | %s\n", t.matrix[1][0], t.matrix[1][1], t.matrix[1][2])
	fmt.Println("──┼───┼──")
	fmt.Printf("%s | %s | %s\n", t.matrix[2][0], t.matrix[2][1], t.matrix[2][2])
}
func (t *tictactoe) GameLoop() {
	t.PrintBoard()
	for t.IsThereSpace() {
		x, y := getInput()
		for x < 0 || x > 2 || y < 0 || y > 2 {
			fmt.Println("Invalid input, numbers must be between 1 and 3")
			x, y = getInput()
		}
		for t.matrix[x][y] != " " {
			fmt.Println("Not possible to make this move! Try other:")
			x, y = getInput()
			for x < 0 || x > 2 || y < 0 || y > 2 {
				fmt.Println("Invalid input, numbers must be between 1 and 3")
				x, y = getInput()
			}
		}
		t.matrix[x][y] = t.player
		t.PrintBoard()
		if won, name := t.SomeoneWon(); won {
			fmt.Printf("%s ganhou!\n", name)
			return
		}

		t.switchPlayer()
	}
	fmt.Println("Draw!")
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

	for _, p := range []string{"X", "O"} {
        // linhas
        for i := 0; i < 3; i++ {
            if t.matrix[i][0] == p && t.matrix[i][1] == p && t.matrix[i][2] == p {
                return true, p
            }
            // colunas
            if t.matrix[0][i] == p && t.matrix[1][i] == p && t.matrix[2][i] == p {
                return true, p
            }
        }
        // diagonais
        if t.matrix[0][0] == p && t.matrix[1][1] == p && t.matrix[2][2] == p {
            return true, p
        }
        if t.matrix[0][2] == p && t.matrix[1][1] == p && t.matrix[2][0] == p {
            return true, p
        }
    }
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

	return y - 1, x - 1
}

func (t *tictactoe)GameLoopOnline(con net.Conn, host bool){

}

func PlayOnline(host bool, ip string) {
    game := tictactoe{
        matrix: [3][3]string{
            {" ", " ", " "},
            {" ", " ", " "},
            {" ", " ", " "},
        },
        player: "X",
    }

    var conn net.Conn
    if host {
        conn = network.StartHost(":8080")
        game.player = "X"
    } else {
        conn = network.Connect(ip)
        game.player = "O"
    }

    game.GameLoopOnline(conn, host)
}

