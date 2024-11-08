package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.Error("The result must be 20")
	}
}
