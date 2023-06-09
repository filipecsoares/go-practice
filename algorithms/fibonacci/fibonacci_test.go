package fibonacci

import "testing"

func TestNthFibonacci(t *testing.T) {
	expected := 1
	result := NthFibonacci(1)
	if expected != result {
		t.Errorf("Expect: %d, Got: %d", expected, result)
	}
	expected = 1
	result = NthFibonacci(2)
	if expected != result {
		t.Errorf("Expect: %d, Got: %d", expected, result)
	}
	expected = 2
	result = NthFibonacci(3)
	if expected != result {
		t.Errorf("Expect: %d, Got: %d", expected, result)
	}
	expected = 3
	result = NthFibonacci(4)
	if expected != result {
		t.Errorf("Expect: %d, Got: %d", expected, result)
	}
	expected = 5
	result = NthFibonacci(5)
	if expected != result {
		t.Errorf("Expect: %d, Got: %d", expected, result)
	}
	expected = 8
	result = NthFibonacci(6)
	if expected != result {
		t.Errorf("Expect: %d, Got: %d", expected, result)
	}
}

func TestIsFibonacci(t *testing.T) {
	expected := true
	result := IsFibonacci(1)
	if expected != result {
		t.Errorf("Expect: %t, Got: %t", expected, result)
	}
	expected = true
	result = IsFibonacci(2)
	if expected != result {
		t.Errorf("Expect: %t, Got: %t", expected, result)
	}
	expected = true
	result = IsFibonacci(3)
	if expected != result {
		t.Errorf("Expect: %t, Got: %t", expected, result)
	}
	expected = false
	result = IsFibonacci(4)
	if expected != result {
		t.Errorf("Expect: %t, Got: %t", expected, result)
	}
	expected = true
	result = IsFibonacci(5)
	if expected != result {
		t.Errorf("Expect: %t, Got: %t", expected, result)
	}
}
