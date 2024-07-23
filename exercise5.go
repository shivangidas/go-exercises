package main

import (
	"math/rand"
	"slices"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

type number struct {
	lenArr int
}

func (n number) singleDigit() []int {
	var numArr = make([]int, n.lenArr)
	for i := 0; i < n.lenArr; i++ {
		numArr[i] = rand.Intn(10)
	}
	return numArr
}

func (n number) doubleDigit() []int {
	var numArr = make([]int, n.lenArr)
	for i := 0; i < n.lenArr; i++ {
		numArr[i] = random(10, 99)
	}
	return numArr
}
func (n number) tripleDigit() []int {
	var numArr = make([]int, n.lenArr)
	for i := 0; i < n.lenArr; i++ {
		numArr[i] = random(100, 999)
	}
	return numArr
}

func exercise5() {
	sN := number{lenArr: 3}
	singleArray := sN.singleDigit()
	doubleArray := sN.doubleDigit()
	tripleArray := sN.tripleDigit()
	pl(singleArray, doubleArray, tripleArray)
	allNum := slices.Concat(singleArray, doubleArray, tripleArray)
	sum := 0
	for _, num := range allNum {
		sum = sum + num
	}
	pl("total sum = ", sum)
}
