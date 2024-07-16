package main

import (
	"fmt"

	"examples.com/recurense"
	"examples.com/variadic_func"
)

// type transformCustomFn func(int)int  //custom function type

func main() {

	num := []int{1, 2, 3, 4}
	doubled := TransformNumbers(&num, double) // just use the function name without "()", so that the value is not returened
	fmt.Println("Doubled Slice", doubled)
	Tripled := TransformNumbers(&num, triple) //  passing func as a parameter for other functions
	fmt.Println("Doubled Slice", Tripled)

	fmt.Println(recurense.Recure(5))

	sum := variadic_func.Sumup(num...) // passing all the slice values to the variadic func
	fmt.Println("sum:", sum)
}

// Generic func which performs doubles and triple
func TransformNumbers(num *[]int, transform func(int) int) []int { // transform func(int)int is passed as a parameter

	TNumbers := []int{}
	for _, value := range *num {
		TNumbers = append(TNumbers, transform(value)) // passing func as a parameter for other functions
	}
	return TNumbers
}

// helper function
func double(number int) int {
	return number * 2
}

// helper function
func triple(number int) int {
	return number * 3
}
