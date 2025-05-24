package database_manager

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/smcri/golang_graph_db/struct_template"
)

func CreateDatabase(db_path string, ctx *struct_template.CurrentContext) error {

	if ctx.CurrentDatabase != "" {
		return fmt.Errorf("cannot create database while in a open database")
	}

	if _, err := os.Stat(db_path + "/relations.json"); err == nil {
		return fmt.Errorf("database already exists")
	}

	err := os.MkdirAll(db_path, 0777)
	if err != nil {
		fmt.Println("Error creating Database directory : ", err)
		return err
	} else {
		fmt.Println("Database directory " + db_path + " created successfully")
	}

	empty_json := make(map[string]interface{})

	jsonString, err := json.MarshalIndent(empty_json, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal JSON string: %w", err)
	}

	// Write back to file
	if err := os.WriteFile(db_path+"/relations.json", jsonString, 0644); err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}

	err = OpenDatabase(db_path, ctx)
	if err != nil {
		return fmt.Errorf("could not open database: %w", err)
	}
	return nil

}

func OpenDatabase(db_path string, ctx *struct_template.CurrentContext) error {

	if ctx.CurrentDatabase != "" {
		return fmt.Errorf("exit database before opening another one")
	} else {
		if _, err := os.Stat(db_path + "/relations.json"); os.IsNotExist(err) {
			return fmt.Errorf("no database exists at location")
		} else if err != nil {
			return fmt.Errorf("error opening database : %f", err)
		}
		ctx.CurrentDatabase = db_path
		fmt.Println("Database at " + ctx.CurrentDatabase + " opened successfully!")
	}

	return nil

}

func ExitDatabase(ctx *struct_template.CurrentContext) error {

	if ctx.CurrentDatabase != "" {
		ctx.CurrentDatabase = ""
		fmt.Println("Database exited successfully!")
	} else {
		return fmt.Errorf("no database to close")
	}

	return nil

}

func DeleteDatabase(db_path string, ctx *struct_template.CurrentContext) error {

	if ctx.CurrentDatabase != "" {
		return fmt.Errorf("please exit database before proceeding to delete database")
	}

	if _, err := os.Stat(db_path + "/relations.json"); err == nil {
		fmt.Println("Deleting database at : ", db_path)
		err := os.RemoveAll(db_path)
		if err != nil {
			return fmt.Errorf("error deleting database : %f", err)
		}
		fmt.Println("Database deleted sucessfully!")

	} else if os.IsNotExist(err) {
		return fmt.Errorf("database does not exist")
	} else {
		return fmt.Errorf("error checking file: %f", err)
	}

	return nil

}
