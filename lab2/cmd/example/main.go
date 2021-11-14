package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	repo "github.com/software-engineering-components/go-architecture/lab2"
)

//
// var (
// 	inputExpression = flag.String("e", "", "Expression to compute")
// 	// TODO: Add other flags support for input and output configuration.
// )
//
// func main() {
// 	flag.Parse()
//
// 	// TODO: Change this to accept input from the command line arguments as described in the task and
// 	//       output the results using the ComputeHandler instance.
// 	//       handler := &lab2.ComputeHandler{
// 	//           Input: {construct io.Reader according the command line parameters},
// 	//           Output: {construct io.Writer according the command line parameters},
// 	//       }
// 	//       err := handler.Compute()
//
// 	res, _ := lab2.PrefixToPostfix("+ 2 2")
// 	fmt.Println(res)
// }



// package main
//
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

func ReadFromInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}

func main() {

	fmt.Print("Enter infix expression: ")
	infixString, err := ReadFromInput()

	if err != nil {
		fmt.Println("Error when scanning input:", err.Error())
		return
	}

	fmt.Println("Ya you postfix notation:", repo.lab2.ToPostfix(infixString))
	return
}
