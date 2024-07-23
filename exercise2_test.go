package main

import "testing"

func TestExercise2(t *testing.T) {
	got := printName("A", "B", "C")
	want := "A B C"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
