package main

import (
	"math/rand"
	"slices"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func roll() int {
	return random(1, 6)
}
func exercise7() {
	for i := 0; i < 50; i++ {
		dice := []int{roll(), roll()}
		natural := []int{7, 12}
		loss := []int{3, 12}
		pf("Roll %d : %d %d\n", i+1, dice[0], dice[1])
		switch {
		case slices.Equal(dice, natural):
			pl("NATURAL")
		case slices.Equal(dice, loss):
			pl("LOSS-CRAPS")
		case dice[0] == 2 || dice[1] == 2:
			pl("SNAKE-EYES-CRAPS")
		default:
			pl("NEUTRAL")
		}
	}
}
