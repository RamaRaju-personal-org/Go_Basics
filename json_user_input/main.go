package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/user_input/note"
	"example.com/user_input/todo"
)

type SaveNoter interface {
	SaveNote() error
}

func main() {
	title, content := GetNoteData()
	todoText := UserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)
    if err !=nil {
		return 
	}




	UserNote, err := note.New(title, content) // calling the constructor func  in note/note.go which has struct values

	if err != nil {
		fmt.Println(err)
		return
	}

	UserNote.Display()
	
	err = saveData(UserNote)
    if err !=nil {
		return 
	}
}

func getTodoData() string {
 return UserInput("Enter Todo:")
}


// func somedata (value interface{}){  ------> in this case any value is allowed 
// 	fmt.Println(value)
// }



//excepts either a todo or note  struct , SaveNoter is an interface
// data variable will be of type interface

func saveData(data SaveNoter ) error { 
	err := data.SaveNote()

	if err !=nil {
		fmt.Println(" Saving the note failed")
		return err 
	}

	fmt.Println("saving the note succeeded!")
	return nil 
}

func GetNoteData() (string, string) {

	title := UserInput("Note title: ")
	content := UserInput("Note content: ")

	return title, content

}

func UserInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin) // constructor that reads user input 
	text, err := reader.ReadString('\n') // stop reading when '\n' comes

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")  // trim 
	text = strings.TrimSuffix(text, "\r")

	return text // returning valid input
}
