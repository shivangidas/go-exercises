package main

import (
	"time"

	"github.com/bearbin/go-age"
)

func exercise6() {
	birthDate := time.Date(1992, 2, 16, 0, 0, 0, 0, time.UTC)
	ageNow := age.AgeAt(birthDate, time.Now())
	pl(ageNow)
}
