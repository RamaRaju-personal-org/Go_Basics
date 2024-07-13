package maps

import "fmt"

func MapWebsite() {

	// map with key and the value are of type string
	websites := map[string]string{
		"Google":"https://google.com",
		"aws": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["aws"])

	websites["LinkedIn"] = "https://linkedin.com"   //adding more key values to map
	fmt.Println(websites)

	delete(websites, "Google") // delete a key in maps
	fmt.Println(websites)


}
