package DA

import "fmt"

type Product struct {
	id string
	title string
	price float64
}


func DynamicArray() {
	// values := []float64{}

	// values = append(values, 12.57)
	// fmt.Println(values)

	hobbies := [3]string{"Learn", "implement", "play"}
	fmt.Println(hobbies)


	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	slice1 := hobbies[:2]
	fmt.Println(slice1)

	slice1 = slice1[1:3]
	fmt.Println(slice1)


	Goals := []string{}  //dynamic array 
	Goals = append(Goals, "Devops")
	fmt.Println(Goals)

// storing Product struct in an products array
	products := []Product{
		{
			"Cloud", 
			"AWS", 
			10.99,
	       },
		{
			"DevOps", 
			"Engineer", 
			20.24,
	       },
		{
			"SRE", 
			"Engineer", 
			20.10,
	       },
		{
			"Backend", 
			"Engineer", 
			20.22,
	       },
		}  

		fmt.Println(products)


	Updateproducts := Product{
		 "software",
		 "engineer",
		 20.58,
	}

	products = append(products, Updateproducts)
	fmt.Println(products)
}
