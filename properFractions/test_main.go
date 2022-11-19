package main

import "testing"

func TestProperFractions(t *testing.T) {

	if ProperFractions(15) != 8 {
		t.Error("Expected 15 to equal 8")
	}
}
