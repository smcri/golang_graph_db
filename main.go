package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
	"github.com/smcri/golang_graph_db/input_parser"
)

func main() {
	fmt.Println("Welcome to golang GraphDB!")

	p := prompt.New(
		input_parser.Parse,
		input_parser.Completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("golang GraphDB shell"),
	)
	p.Run()

}
