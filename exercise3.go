package main

func exercise3() {
	var input int
	pl("Enter an integer")
	sl(&input)
	if input > 1 && input < 10 {
		pf("your number %d is between 1 and 10\n", input)
	} else {
		pl("your number is wrong")
	}
}
