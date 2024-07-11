// Note should have the file BalanceReceipt.txt with data before starting the program
// to import third party packages : go get github_url 
// to install the third party pacakages mentioned in go.mod file : go get
package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/pallinder/go-randomdata" // third-party library 
)

const accountBalanceFile = "BalanceReceipt.txt" //global declaration


func main() {

	var balance, err = fileops.GetBalanceFromFile(accountBalanceFile) // gets the balance from the file BalanceReceipt.txt
	if err != nil {                                           // error handling
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("can't continue, sorry.")  (this will also exit the program with the mentioned message)
		// return // this will exit the program
	}

	fmt.Println("welcome to Go Bank")
	fmt.Println("what do you want to do?")
	fmt.Println("Reach Us at: ", randomdata.PhoneNumber())
	
	for { // loop executes until the user enters 4 to exit

		options()

		var choice int
		fmt.Print("Enter your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Current account balance:", balance)

		case 2:
			var depositAmount float64
			fmt.Print("enter amount to deposit")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit amount should be greater than 0")
				continue
			}
			balance += depositAmount
			fmt.Println("Updated Balanace:", balance)
			fileops.WriteBalanceToFile(balance, accountBalanceFile)
			fmt.Println("Balance Receipt generated succesfully")

		case 3:
			var withDrawAmount float64
			fmt.Print("Enter amount to withdraw:")
			fmt.Scan(&withDrawAmount)
			if withDrawAmount > balance {
				fmt.Println("cannot withdraw greater than your balance")
				continue
			} else if withDrawAmount <= 0 {
				fmt.Println(" withdraw amount should be greather than 0")
			} else {
				balance -= withDrawAmount
				fmt.Println("Updated balance after Withdrwal:", balance)
				fileops.WriteBalanceToFile(balance, accountBalanceFile)
				fmt.Println("Balance Receipt generated succesfully")
			}
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing Go Bank")
			return // ends the switch and exits the main
		}
	}

}
