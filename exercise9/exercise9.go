package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println
var sl = fmt.Scanln

type Name struct {
	firstName  string
	middleName string
	lastName   string
}

func (n Name) joinName() string {
	jointName := []string{n.firstName, n.middleName, n.lastName}
	return strings.Join(jointName[:], " ")
}

func (n *Name) inputReceiver() {
	pl("What's your name?")
	fmt.Printf("Enter first name: ")
	sl(&n.firstName)
	fmt.Printf("Enter middle name: ")
	sl(&n.middleName)
	fmt.Printf("Enter last name: ")
	sl(&n.lastName)
}

func (n Name) printToScreen() {
	pl("full-name: ", n)
	pl("middle-name: ", n.middleName)
	pl("surname: ", n.lastName)
}

func main() {
	n := Name{}
	n.inputReceiver()
	pl(n.joinName())
	n.printToScreen()
}
