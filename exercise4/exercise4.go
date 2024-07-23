package main

import (
	"fmt"
	"sort"
)

var pl = fmt.Println

func reverseFun(numArr []int) []int {
	reverse := make([]int, len(numArr))
	copy(reverse, numArr)
	sort.Slice(reverse, func(i, j int) bool {
		return reverse[i] > reverse[j]
	})
	return reverse
}
func count(arr []int) (int, int) {
	evenCount, oddCount := 0, 0
	for _, num := range arr {
		if num%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return evenCount, oddCount
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pl(arr)

	reverse := reverseFun(arr)
	pl(reverse)

	evenCount, oddCount := count(arr)
	pl("Number of even numbers: ", evenCount)
	pl("Number of odd numbers: ", oddCount)
}
