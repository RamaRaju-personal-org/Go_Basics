package main

import "fmt"
func main(){
	age := 24 

	agePointer := &age

	fmt.Println("age:", age)
	adultyears := GetAdultYears(*agePointer)
	fmt.Println(adultyears)
}

func GetAdultYears(age int) int {
	return age -18
}
