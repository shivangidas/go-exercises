package main

import (
	"fmt"
)

func printName(firstName, middleName, lastName string) string {
	return firstName + " " + middleName + " " + lastName
}
func exercise2() {
	var firstName, middleName, lastName string
	pl("What's your name?")
	fmt.Printf("Enter first name: ")
	sl(&firstName)
	fmt.Printf("Enter middle name: ")
	sl(&middleName)
	fmt.Printf("Enter last name: ")
	sl(&lastName)
	pl(printName(firstName, middleName, lastName))
}
