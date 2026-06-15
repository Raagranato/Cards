package tictactoe


import("fmt" 
"os"
"bufio"
"strings"
)
type tictactoe struct {
	matrix [3][3]string //" " ou X ou O
}

func (t *tictactoe)PrintBoard(){
	fmt.Printf("%s|%s|%s\n",t.matrix[0][0],t.matrix[0][1],t.matrix[0][2])
	fmt.Printf("%s|%s|%s\n",t.matrix[1][0],t.matrix[1][1],t.matrix[1][2])
	fmt.Printf("%s|%s|%s\n",t.matrix[2][0],t.matrix[2][1],t.matrix[2][2])
}
func (t *tictactoe)GameLoop(){

}
func GetInput() string {
	fmt.Print("What is your move? ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}