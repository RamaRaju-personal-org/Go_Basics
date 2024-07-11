package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string //exposed 
	Content string //exposed 
	CreatedAt time.Time  //exposed 
}


// method that interact with Note struct
func (n Note) Display () {
	fmt.Printf("Your note title %v has the content:\n%v\n", n.Title, n.Content)
}

// method that interact with Note struct and save it to a file
func (n Note) SaveNote() error{
	fileName := strings.ToLower(n.Title) + ".json" // file name from upper to lower case if any and to json 

	json, err := json.Marshal(n) // converting to json

	if err !=nil {
		return err
	}
	
	return os.WriteFile(fileName, json, 0644)  //json content is written to the file 
	
}

//constructor that returns struct
func New(title, content string) (Note, error) {  
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input") //returns empty Note struct in case of error 
	}	
	
	return Note{ //return the updated Note struct instance 
			Title: title,
			Content: content,
			CreatedAt: time.Now(),
		}, nil 
}
