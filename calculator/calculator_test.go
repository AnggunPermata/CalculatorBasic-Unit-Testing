package calculator

import (
	"testing"
)

func TestAddition(t *testing.T) {
	if Addition(4, 5) != 9 {
		t.Error("Expected 4 (+) 5 equal to 9")
	}
	if Addition(-4, -5) != -9 {
		t.Error("Expected -4 (+) -5 equal to -9")
	}
}

func TestSubstraction(t *testing.T) {
	if Substraction(5, 4) != 1 {
		t.Error("Expected 5 (-) 4 equal to 1")
	}
	if Substraction(-5, -4) != -1 {
		t.Error("Expected -5 (-) -4 equal to -1")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(5, 4) != 20 {
		t.Error("Expected 4 (*) 5 equal to 20")
	}
	if Multiplication(-5, -4) != 20 {
		t.Error("Expected -5 (*) -4 equal to 20")
	}
	if Multiplication(-5, 4) != -20 {
		t.Error("Expected -5 (*) 4 equal to -20")
	}
}

func TestDivision(t *testing.T) {
	if Division(10, 5) != 2 {
		t.Error("Expected 10 (/) 5 equal to 2")
	}
	if Division(10, -5) != -2 {
		t.Error("Expected 10 (/) -5 equal to -2")
	}
	if Division(-10, -5) != 2 {
		t.Error("Expected -10 (/) -5 equal to 2")
	}
}
