// To create a moudule : go mod init profit-calculator.com/practice

package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	outputText("profit calculator by ", 1)

	fmt.Print("Revenue:")
	fmt.Scan(&revenue)
	fmt.Print("Expenses:")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate:")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	fmt.Println("ebt:", ebt)
	fmt.Println("profit:",profit)
	fmt.Println("ratio", ratio)
}

func outputText(mssg string, userID int) { //custom user defined function
	fmt.Println(mssg, "user id:", userID)
}

func calculateFinancials(revenue, expenses, taxRate float64 ) ( float64, float64, float64){
	ebt := revenue - expenses // defining a variable and it's going to pick the type automatically
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}
