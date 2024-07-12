package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	
)

type Todo struct {
	Text string `json:"content"`          //exposed and Struct Tags ``
}


// method that interact with Note struct
func (t Todo) Display () {
	fmt.Printf(t.Text)
}

// method that interact with Note struct and save it to a file
func (t Todo) SaveNote() error{
	fileName :=  "todo.json" 

	json, err := json.Marshal(t) // converting to json

	if err !=nil {
		return err
	}
	
	return os.WriteFile(fileName, json, 0644)  //json content is written to the file 
	
}

//constructor that returns struct
func New(content string) (Todo, error) {  
	if  content == "" {
		return Todo{}, errors.New("invalid input") //returns empty Note struct in case of error 
	}	
	
	return Todo{ //return the updated Note struct instance 
			Text: content,
		}, nil 
}
