package builder

import "testing"

func TestPersonBuilder(t *testing.T) {
    expected := Person{"Thrs", 52, 170}
    result := NewPersonBuilder().name("Thrs").age(52).height(170).build()
    if result != expected {
        t.Errorf("Expected: %v, Got: %v", expected, result)
    }
}