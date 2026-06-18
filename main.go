package main

//go run main.go
import (
	"FeelGoodInc/games/blackjack"
	"FeelGoodInc/games/minesweeper"
	"FeelGoodInc/games/tictactoe"
	//"fmt"
	"FeelGoodInc/internal/ui"
	//"FeelGoodInc/styles"
	"FeelGoodInc/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

/*TODO:
-Shake tela quando selecionar uma bomba no minesweeper
-Play localy
-Play online(host e client)
-Upgrades?
-Bet in blackjack(atributo(global?) q é chamado da main e salvo no json, double)
-Jogos:
xadrez
craps
batalha naval
Jogo da velha - multplayer local
roleta -> Ok
truco - local deve ser legal
liars dice
horse race
-
-
-
-
-
-*/

func main() {
	//fmt.Println(styles.Welcome.Render("Feel Good Inc"))
	ui.Welcome()

	m := ui.FirstState{
		Choices: []string{"Blackjack", "Minesweeper", "Tictactoe"},
	}
	result, _ := tea.NewProgram(m).Run()
	finalState := result.(ui.FirstState)
	utils.ClearTerminal()
	utils.SkipLine()
	switch finalState.Opc {
	case 0:
		blackjack.Play()
	case 1:
		minesweeper.Play()
	case 2:
		tictactoe.Play()
	}
}
