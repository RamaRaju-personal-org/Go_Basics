package main

import (
	"fmt"
	"exmaples.com/struct/user_module"
)

func main() {
	userFirstName := user_module.GetUserData("Please enter your first name: ")
	userLastName := user_module.GetUserData("Please enter your last name: ")
	userBirthDate := user_module.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user_module.User // defining appUser variable of custom struct type User

	// appUser = User {} defining empty struct

	appUser = user_module.User{ // instance of struct User with values
		FirstName: userFirstName,
		LastName:  userLastName,
		Birthdate: userBirthDate,
	}

	
	fmt.Println("User struct details:")
	user_module.OutputUserDetails(appUser) // we just need to pass the variable name assigned to struct instead of passing each and every variable
    
	


	//calling the method : appuser.OutputUserDetails()
	appUser.ClearUserName() //method clears first & last name

	fmt.Println("output of User struct after clearing some data:")
	user_module.OutputUserDetails(appUser) // after clearing some data

}
