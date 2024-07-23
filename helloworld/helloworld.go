package main

import "fmt"

const hello = "Hello"
const hola = "Hola"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	switch language {
	case "Spanish":
		return hola + ", " + name
	default:
		return hello + ", " + name
	}
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("Shivangi", "Spanish"))
}
