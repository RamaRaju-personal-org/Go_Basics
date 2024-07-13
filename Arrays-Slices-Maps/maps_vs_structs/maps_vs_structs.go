package maps_vs_structs

import (
	"fmt"
)

// Define a struct
type Product struct {
    Name  string
    Price float64
    Stock int
}


func Maps_structs(){
	 
	
	// Declare a map
	 prices := map[string]float64{
        "Apple":  1.99,
        "Banana": 0.99,
        "Cherry": 2.99,
    }

    // Accessing elements
    fmt.Println(prices["Apple"]) // 1.99

    // Adding a new key-value pair
    prices["Date"] = 3.49

    // Deleting a key-value pair
    delete(prices, "Banana")

    // Iterating over a map
    for fruit, price := range prices {
        fmt.Println(fruit, price)
    }

    

	// Initialize a struct
    product := Product{
        Name:  "Apple",
        Price: 1.99,
        Stock: 100,
    }

    // Accessing struct fields
    fmt.Println(product.Name)  // Apple
    fmt.Println(product.Price) // 1.99
    fmt.Println(product.Stock) // 100

    // Updating struct fields
    product.Stock = 120
    fmt.Println(product.Stock) // 120

    // Another way to declare and initialize a struct
    anotherProduct := Product{"Banana", 0.99, 200}
    fmt.Println(anotherProduct)


}


// 
// Summary
// Maps: Best for dynamic collections with key-based access. Use when the number of elements can change, and quick lookup by key is needed.
// Structs: Best for fixed data structures with named fields. Use when you have a set of related fields that will be accessed and modified together.
