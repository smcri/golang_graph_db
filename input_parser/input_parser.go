package input_parser

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/smcri/golang_graph_db/file_io"
)

func read_key(key string, node string) (interface{}, error) {

	node_json, err := file_io.ReadFile(node)
	if err != nil {
		return nil, err
	}
	if key == "*" {
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

func Parse(input string) error {
	parse_list := strings.Fields(input)
	switch parse_list[0] {
	case "READ":
		read_key(parse_list[1], parse_list[2])
	case "WRITE":
		write_key(parse_list[1], parse_list[2], parse_list[3])
	case "DELETE":
		delete_key(parse_list[1], parse_list[2])
	case "DELETE_NODE":
		delete_node(parse_list[1])
	case "CREATE_RELATION":
		create_relation(parse_list[1], parse_list[2], parse_list[3])
	default:
		return fmt.Errorf("unknown command: %s", parse_list[0])

	}

	return nil
}
