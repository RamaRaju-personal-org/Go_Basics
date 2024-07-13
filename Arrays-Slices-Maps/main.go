package main

import (
	"fmt"

	"examples.com/list"
	"examples.com/maps"
	"examples.com/maps_vs_structs"
)

func main() {
	list.ListOperations()
	maps.MapWebsite()
	maps_vs_structs.Maps_structs()

	// userNames := []string{} // empty slice

	// using make for slice
	userNames := make([]string, 2, 5) //empty slice with length 2 and with max capacity 5
	userNames[0] = "Raju"
	userNames[1] = "vj"
	userNames = append(userNames, "Ram") // updating a slice
	fmt.Println(userNames)

	// using make for map
	courseRatings := make(map[string]float64, 3) // current cap is 3 for map
	courseRatings["go"] = 4.7
	courseRatings["DevOps"] = 5.0

	fmt.Println(courseRatings)

	//for loop for Slices
	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	//for loop for maps
	for key, value := range courseRatings {
		fmt.Println("key:", key)
		fmt.Println("Value:", value)
	}
}
