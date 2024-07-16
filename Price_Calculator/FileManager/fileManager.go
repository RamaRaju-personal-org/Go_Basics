package FileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string)([]string, error) {
	file, err := os.Open(path)
	if err!=nil {
		fmt.Println(err)
		return nil, errors.New("failed to open file")
	}

	scanner :=bufio.NewScanner(file)
	var lines []string
	//reads oneline at a time and returns bool 0 if no lines to be read
	for scanner.Scan()  {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err() //call if error occurs 
	if err !=nil {
		file.Close()
		fmt.Println(err)
		return nil, errors.New("failed to read line in file")

	}

	file.Close() // close after reading the file 
	return lines, nil
    
}



func WriteJSON (path string, data any) error { // any values allowed
  file, err := os.Create(path) // create or overwrite the file 

  if err!=nil {
	return errors.New("Failed to create file.")
  }

  encoder := json.NewEncoder(file)
  err = encoder.Encode(data)

  if err != nil {
	file.Close()
	return errors.New("failed to convert data to json")
  }

  file.Close()
  return nil

}
