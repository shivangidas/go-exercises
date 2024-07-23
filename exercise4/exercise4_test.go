package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("sort a slice of numbers in descending order", func(t *testing.T) {
		got := reverseFun([]int{3, 4, 2, 6, 8})
		want := []int{8, 6, 4, 3, 2}
		assertMessage(t, got, want)
	})
	t.Run("return empty is nothing is sent", func(t *testing.T) {
		got := reverseFun([]int{})
		want := []int{}
		assertMessage(t, got, want)
	})
}
func TestCount(t *testing.T) {
	t.Run("count even and odd", func(t *testing.T) {
		got1, got2 := count([]int{3, 4, 2, 6, 8})
		want1, want2 := 4, 1
		if got1 != want1 || got2 != want2 {
			t.Errorf("doesn't work")
		}
	})
	t.Run("count with empty slice", func(t *testing.T) {
		got1, got2 := count([]int{})
		want1, want2 := 0, 0
		if got1 != want1 || got2 != want2 {
			t.Errorf("doesn't work")
		}
	})
}
func assertMessage(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
}
