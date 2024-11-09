package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.Error("The result must be 20")
	}

}

func TestSub(t *testing.T) {
	result := Sub(30, 10)

	if result != 20 {
		t.Error("The result must be 20")
	}
}

func TestTimes(t *testing.T) {
	result := Times(2, 10)

	if result != 20 {
		t.Error("The result must be 20")
	}
}

func TestSumX(t *testing.T) {
	result := SumX(10, 10)

	if result != 30 {
		t.Error("The result must be 20")
	}
}
