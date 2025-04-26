package main

import (
	"bufio"
	"os"

	"github.com/smcri/golang_graph_db/input_parser"
)

func user_interface_init() {
	reader := bufio.NewReader(os.Stdin)

	for {

		var input string
		input, _ = reader.ReadString('\n')
		input_parser.Parse(input)

	}

}

func main() {
	user_interface_init()
}
