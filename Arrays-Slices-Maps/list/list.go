package list

import (
	"fmt"
	"examples.com/DA"
)

// Note : In Go, the append function can only be used with slices, not arrays

func ListOperations() {
	var productNames [4]string // array of strings
	productNames = [4]string{"A", "B", "C", "D"}

	productNames [2] = "Z"

	prices := [4]float64{10.99, 9.9, 6.77, 10.15} // array of 4 values
	fmt.Println(prices)
	fmt.Println(productNames)
	
	fmt.Println(prices[0]) // access specific item 

	featuredPrices := prices[1:3] // slice
	featuredPrices1 := prices[:3] // slice
	featuredPrices2 := prices[1:] // slice

	fmt.Println(len(featuredPrices), cap(featuredPrices)) // len - no of items , 

	fmt.Println(featuredPrices)
	fmt.Println(featuredPrices1)
	fmt.Println(featuredPrices2)

   
	discountPrices := []float64{2.99, 1.9, 1.77}

	// convert prices array to a slice before appending
	pricesSlice := prices[:]

	// exactly 3 dots indicating 3 elemnts of discountPrices i.e unpacking discountPrices list values
	pricesSlice = append(pricesSlice, discountPrices...) 
	fmt.Println(pricesSlice)

	DA.DynamicArray()
}
