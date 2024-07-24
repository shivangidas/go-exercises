package main

import (
	"fmt"
	"math/rand"
	"slices"
)

var pl = fmt.Println

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

type number struct {
	lenArr int
}

func (n number) SingleDigit() []int {
	var numArr = make([]int, n.lenArr)
	for i := 0; i < n.lenArr; i++ {
		numArr[i] = rand.Intn(10)
	}
	return numArr
}

func (n number) DoubleDigit() []int {
	var numArr = make([]int, n.lenArr)
	for i := 0; i < n.lenArr; i++ {
		numArr[i] = random(10, 99)
	}
	return numArr
}
func (n number) TripleDigit() []int {
	var numArr = make([]int, n.lenArr)
	for i := 0; i < n.lenArr; i++ {
		numArr[i] = random(100, 999)
	}
	return numArr
}

func main() {
	sN := number{lenArr: 3}
	singleArray := sN.SingleDigit()
	doubleArray := sN.DoubleDigit()
	tripleArray := sN.TripleDigit()
	pl(singleArray, doubleArray, tripleArray)
	allNum := slices.Concat(singleArray, doubleArray, tripleArray)
	sum := 0
	for _, num := range allNum {
		sum = sum + num
	}
	pl("total sum = ", sum)
}
