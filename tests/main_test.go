package main

import "testing"

func TestAdder(t *testing.T) {
	// Arrange
	l, r := 20, 5
	expected := 25
	// Act
	got := Adder(l, r)
	// Assert
	if expected != got {
		t.Errorf("Failed to add %v with %v. Got %v and expected %v\n",
			l, r, got, expected)
	}

}
