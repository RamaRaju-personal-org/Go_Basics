// To create a moudule : go mod init investment-program.com/investment-cal

package main // first instruction

import (
	"errors"
	"fmt"
	"math"
	"os"
)

const inflationRate = 2.5 //doesn't change and global declaration

func main() {

	var investAmount, years, expectedReturnRate float64 // local declaration

	// variable values can be changed
	// a:= 1.1 can also delclare like this and the type is automatically fetched

	fmt.Print("Amount You want to Investment: ")
	fmt.Scan(&investAmount) // pointer to get the input from user dynamically

	years, err1 := getUserInput("Total number of years to invest: ") // another way of giving inupt using fuction

	if err1 != nil { // error handling for years
		fmt.Println(err1)
		return // exits the code
		//(or) panic(err1)
	}

	expectedReturnRate, err2 := getUserInput("expected return rate: ")

	if err2 != nil { // error handling for expected return rate
		fmt.Println(err2)
		return // exits the code
		//(or) panic(err2)

	}

	futureValue, futureRealValue := calculateFutureValues(investAmount, expectedReturnRate, years)

	fmt.Println("Future Amount :", futureValue)
	fmt.Println("Future Amount with inflation of 2.5 :", futureRealValue)
	writeValuesToFile(futureValue, futureRealValue)

	// to run the program : go run firstCode.go
	// go run . will execute the module directly
}

func calculateFutureValues(investAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investAmount * math.Pow(1+expectedReturnRate/100, years)
	frv = fv / math.Pow(1+inflationRate/100, years)

	return fv, frv // returning multiple values
}

// calculateFutureValues(investAmount, expectedReturnRate, years float64) (fv float64,  frv float64) - to return two values fv ,frv

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("user input should be greater than 0") // error handling
	}
	return userInput, nil // return without errors if we dont have any
}

func writeValuesToFile(futureValue, futureRealValue float64) { // writing to a file
	results := fmt.Sprintf("futureValue: %.1f\nfutureRealValue: %.1f\n", futureValue, futureRealValue) // using Sprintf to convert balance to string and format
	os.WriteFile("result.txt", []byte(results), 0644)                                                  // 0644 file permissions
}
