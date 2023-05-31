package factorial

import "testing"

func TestFactorial(t *testing.T) {
	expected := 1
	got := Factorial(1)
	if expected != got {
		t.Errorf("Expected: %d, Got: %d", expected, got)
	}
	expected = 0
	got = Factorial(0)
	if expected != got {
		t.Errorf("Expected: %d, Got: %d", expected, got)
	}
	expected = 120
	got = Factorial(5)
	if expected != got {
		t.Errorf("Expected: %d, Got: %d", expected, got)
	}
	expected = 720
	got = Factorial(6)
	if expected != got {
		t.Errorf("Expected: %d, Got: %d", expected, got)
	}
}
