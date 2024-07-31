package main

import (
	"fmt"
	"time"

	"github.com/bearbin/go-age"
)

var pl = fmt.Println

type Student struct {
	fullName string
	dob      time.Time
	age      int
}

type Register map[int]Student

func (r Register) ListStudents() {
	for i := 0; i < len(r); i++ {
		pl(r[i].fullName, r[i].dob.Format("2006-01-02"), r[i].age)
	}
}
func (r Register) AddStudents(name string, dob time.Time) {
	age := age.AgeAt(dob, time.Now())
	student := Student{name, dob, age}
	r[len(r)] = student
}
func main() {
	r := Register{}
	r.AddStudents("Shivangi", time.Date(1992, 2, 16, 0, 0, 0, 0, time.UTC))
	r.AddStudents("ABC", time.Date(1995, 2, 16, 0, 0, 0, 0, time.UTC))
	r.AddStudents("GHJ", time.Date(2000, 4, 20, 0, 0, 0, 0, time.UTC))
	r.ListStudents()
}
