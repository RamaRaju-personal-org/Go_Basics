package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Note : always have starting letter of function in caps to use the func in other file 

func WriteBalanceToFile(balance float64, fileName string ) { // writing to a file
	balanceText := fmt.Sprint(balance)                          // using sprint to convert balance to byte
	os.WriteFile(fileName, []byte(balanceText), 0644) // 0644 file permissions
}

func GetBalanceFromFile(fileName string) (float64, error) { // read data from the file
	data, err := os.ReadFile(fileName)

	if err != nil { // error handling
		return 0, errors.New("unable to fetch the balance from file, balance set to 0")
	}

	balanceFromFile := string(data)                               // only string works
	balanceInFile, err := strconv.ParseFloat(balanceFromFile, 64) //64 bit

	if err != nil { // error handling
		return 0, errors.New("unable to parse the balance from file")
	}

	return balanceInFile, nil //returning without error
}
