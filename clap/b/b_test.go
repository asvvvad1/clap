package b

import "testing"

func TestB(t *testing.T) {
	result := B("DAMN BOI")
	expected := "DAMN 🅱️OI"

	if result != expected {
		t.Errorf("FAILED: Expected "+expected+"\" got %v", result)
	} else {
		t.Log("PASSED")
	}
}
