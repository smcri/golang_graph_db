package input_parser

import (
	"errors"
	"fmt"
	"os"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/smcri/golang_graph_db/file_io"
)

func read_key(key string, node string) (interface{}, error) {

	node_json, err := file_io.ReadFile(node)
	if err != nil {
		return nil, err
	}
	if key == "*" {
		fmt.Println("JSON : ", node_json)
		return node_json, err
	}
	required_value, exists := node_json[key]
	if !exists {
		return nil, fmt.Errorf("key '%s' not found", key)
	}
	fmt.Println("Required value : ", required_value)
	return required_value, nil

}

func write_key(key string, value interface{}, node string) error {

	err := file_io.WriteFile(key, value, node)
	if err == nil {
		fmt.Println("Write success!")
	}
	return err

}

func delete_key(key string, node string) error {
	node_json, err := file_io.ReadFile(node)
	delete(node_json, key)
	if err == nil {
		fmt.Println("Key delete success!")
	}
	return err
}

func delete_node(node string) error {
	err := file_io.DeleteFile(node)
	if err == nil {
		fmt.Println("Node delete success!")
	}
	return err
}

func create_relation(from string, to string, relation string) error {
	node_json, err := file_io.ReadFile("relations.json")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		fmt.Println(err)
		return err
	} else {
		node_json = make(map[string]interface{})
	}

	rel, ok := node_json[relation]
	if !ok {

		node_json[relation] = make(map[string][]string)
		rel = node_json[relation]
	}

	relMap, ok := rel.(map[string][]string)
	if !ok {
		return fmt.Errorf("relation '%s' has invalid type", relation)
	}

	if _, exists := relMap[from]; !exists {

		relMap[from] = []string{}
	}

	relMap[from] = append(relMap[from], to)

	err = file_io.WriteFile(relation, relMap, "relations")
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Println("Relation creation success!")

	return nil

}

func Parse(input string) {
	parse_list := strings.Fields(input)

	if len(parse_list) == 0 {
		fmt.Println("No command entered.")
		return
	}

	if parse_list[0] == "EXIT" {
		fmt.Println("Exiting the shell...")
		os.Exit(0)
		return
	}

	switch parse_list[0] {
	case "READ":
		if len(parse_list) < 3 {
			fmt.Println("Usage: READ <key> <node>")
			return
		}
		read_key(parse_list[1], parse_list[2])

	case "WRITE":
		if len(parse_list) < 4 {
			fmt.Println("Usage: WRITE <key> <value> <node>")
			return
		}
		write_key(parse_list[1], parse_list[2], parse_list[3])

	case "DELETE":
		if len(parse_list) < 3 {
			fmt.Println("Usage: DELETE <key> <node>")
			return
		}
		delete_key(parse_list[1], parse_list[2])

	case "DELETE_NODE":
		if len(parse_list) < 2 {
			fmt.Println("Usage: DELETE_NODE <node>")
			return
		}
		delete_node(parse_list[1])

	case "CREATE_RELATION":
		if len(parse_list) < 4 {
			fmt.Println("Usage: CREATE_RELATION <source_node> <target_node> <relation>")
			return
		}
		create_relation(parse_list[1], parse_list[2], parse_list[3])

	default:
		fmt.Printf("Unknown command: %s\n", parse_list[0])
	}
}

func Completer(d prompt.Document) []prompt.Suggest {
	text := strings.TrimSpace(d.TextBeforeCursor())
	words := strings.Fields(text)

	// Base commands
	baseCommands := []prompt.Suggest{
		{Text: "READ", Description: "Read a key from a node"},
		{Text: "WRITE", Description: "Write a value to a key in a node"},
		{Text: "DELETE", Description: "Delete a key from a node"},
		{Text: "DELETE_NODE", Description: "Delete a node"},
		{Text: "CREATE_RELATION", Description: "Create a relation between two nodes"},
		{Text: "EXIT", Description: "Exit the shell"},
	}

	if len(words) == 0 {
		return baseCommands
	}

	switch words[0] {
	case "READ":
		switch len(words) {
		case 1:
			return []prompt.Suggest{{Text: "<key>", Description: "Key name"}}
		case 2:
			return []prompt.Suggest{{Text: "<node>", Description: "Node ID"}}
		}
	case "WRITE":
		switch len(words) {
		case 1:
			return []prompt.Suggest{{Text: "<key>", Description: "Key name"}}
		case 2:
			return []prompt.Suggest{{Text: "<value>", Description: "Value to store"}}
		case 3:
			return []prompt.Suggest{{Text: "<node>", Description: "Node ID"}}
		}
	case "DELETE":
		switch len(words) {
		case 1:
			return []prompt.Suggest{{Text: "<key>", Description: "Key name"}}
		case 2:
			return []prompt.Suggest{{Text: "<node>", Description: "Node ID"}}
		}
	case "DELETE_NODE":
		if len(words) == 1 {
			return []prompt.Suggest{{Text: "<node>", Description: "Node ID to delete"}}
		}
	case "CREATE_RELATION":
		switch len(words) {
		case 1:
			return []prompt.Suggest{{Text: "<source_node>", Description: "Source node"}}
		case 2:
			return []prompt.Suggest{{Text: "<target_node>", Description: "Target node"}}
		case 3:
			return []prompt.Suggest{{Text: "<relation>", Description: "Relationship label"}}
		}
	default:
		return prompt.FilterHasPrefix(baseCommands, words[0], true)
	}

	return []prompt.Suggest{}
}
