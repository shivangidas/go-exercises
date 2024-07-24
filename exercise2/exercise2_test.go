package main

import "testing"

func TestJoinName(t *testing.T) {
	tests := []struct {
		firstName, middleName, lastName, expected string
	}{
		{"A", "", "C", "A C"},
		{"John", "Doe", "", "John Doe"},
		{"A", "B", "C", "A B C"},
	}

	for _, test := range tests {
		result := joinName(test.firstName, test.middleName, test.lastName)
		if result != test.expected {
			t.Errorf("printName(%q, %q, %q) = %q; want %q",
				test.firstName, test.middleName, test.lastName, result, test.expected)
		}
	}

}
