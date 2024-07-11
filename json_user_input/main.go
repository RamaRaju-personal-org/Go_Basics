package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/user_input/note"
)

func main() {
	title, content := GetNoteData()

	UserNote, err := note.New(title, content) // calling the constructor func  in note/note.go which has struct values

	if err != nil {
		fmt.Println(err)
		return
	}

	UserNote.Display()
	
	err = UserNote.SaveNote()

	if err !=nil {
		fmt.Println(" Saving the note failed")
		return
	}

	fmt.Println("saving the note succeeded!")

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
