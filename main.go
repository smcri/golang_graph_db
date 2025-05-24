package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
	"github.com/smcri/golang_graph_db/input_parser"
	"github.com/smcri/golang_graph_db/struct_template"
)

func main() {
	fmt.Println("Welcome to golang GraphDB!")
	ctx := &struct_template.CurrentContext{CurrentDatabase: ""}

	p := prompt.New(
		func(in string) { input_parser.Parse(in, ctx) },
		input_parser.Completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("golang GraphDB shell"),
	)
	p.Run()

}
