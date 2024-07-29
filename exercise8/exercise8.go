package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var cityList = [...]string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}

const fileName = "cities.txt"

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	f := createFile(fileName)
	defer closeFile(f)
	writeFile(f)
	cities := readFile(fileName)
	sorted := sortCities(cities)
	fmt.Println(sorted)
}
func createFile(p string) *os.File {
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Fprintln(f, strings.Join(cityList[:], ","))
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v \n", err)
		os.Exit(1)
	}
}

func readFile(fileName string) string {
	data, err := os.ReadFile(fileName)
	checkError(err)
	return string(data)
}

func sortCities(cities string) []string {
	newCityList := strings.Split(cities, ",")
	slices.Sort(newCityList)
	return newCityList
}
