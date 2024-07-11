package user_module

import (
 "fmt"
 "time"
)

// use uppercase for struct and variables

type  User struct {  // group related data together using struct (or) a blueprint
	FirstName  string // feild with type , use uppercase for exporting
	LastName   string
	Birthdate  string
	createdAt  time.Time  
	// Note : createdAt only available within this module because createdAt is lower case
	// this is useful if we want it to to be available only for this module
}

// Note : attaching functions to struct is called method 
// example method : func (u User ) OutputUserDetails () {}  --- > now this func is part of struct 




func OutputUserDetails (u User){   // receives a struct 
	fmt.Println(u.FirstName, u.LastName, u.Birthdate) // accessing struct and printing the values 
}





func  (u *User )ClearUserName() {     // method which is part of User struct
	// i'm using the pointer to reference the address of User struct to change the below values 
	u.FirstName = "" // we are editing the first name and last name 
	u.LastName = ""
}


func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
