package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println
var sl = fmt.Scanln

func joinName(firstName, middleName, lastName string) string {
	name := []string{}
	if firstName != "" {
		name = append(name, firstName)
	}
	if middleName != "" {
		name = append(name, middleName)
	}
	if lastName != "" {
		name = append(name, lastName)
	}
	return strings.Join(name, " ")
}
func main() {
	var firstName, middleName, lastName string
	pl("What's your name?")
	fmt.Printf("Enter first name: ")
	sl(&firstName)
	fmt.Printf("Enter middle name: ")
	sl(&middleName)
	fmt.Printf("Enter last name: ")
	sl(&lastName)
	pl(joinName(firstName, middleName, lastName))
}
