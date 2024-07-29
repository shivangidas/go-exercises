package main

import (
	"fmt"
	"io"
	"os"
)

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
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}
func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("Shivangi", "Spanish"))
	Greet(os.Stdout, "Shivangi")
}
