package clap

import "testing"

func TestBetweenWords(t *testing.T) {
	result := BetweenWords("mitochondria is the powerhouse of the cell", "🔪️")
	expected := "mitochondria🔪️is🔪️the🔪️powerhouse🔪️of🔪️the🔪️cell"

	if result != expected {
		t.Errorf("FAILED: Expected "+expected+"\" got %v", result)
	} else {
		t.Log("PASSED")
	}
}

func TestBetweenLetter(t *testing.T) {
	result := BetweenLetters("mitochondria", "🔪️")
	expected := "M🔪️I🔪️T🔪️O🔪️C🔪️H🔪️O🔪️N🔪️D🔪️R🔪️I🔪️A"

	if result != expected {
		t.Errorf("FAILED: Expected "+expected+"\" got %v", result)
	} else {
		t.Log("PASSED")
	}
}

func TestClap(t *testing.T) {
	result := Clap("mitochondria is the powerhouse of the cell")
	expected := "mitochondria👏️is👏️the👏️powerhouse👏️of👏️the👏️cell"

	if result != expected {
		t.Errorf("FAILED: Expected "+expected+"\" got %v", result)
	} else {
		t.Log("PASSED")
	}

	result = Clap("mitochondria")
	expected = "M👏️I👏️T👏️O👏️C👏️H👏️O👏️N👏️D👏️R👏️I👏️A"

	if result != expected {
		t.Errorf("FAILED: Expected "+expected+"\" got %v", result)
	} else {
		t.Log("PASSED")
	}
}

func TesEmoji(t *testing.T) {
	result := Emoji("mitochondria is the powerhouse of the cell", "🔪️")
	expected := "mitochondria🔪️is🔪️the🔪️powerhouse🔪️of🔪️the🔪️cell"

	if result != expected {
		t.Errorf("FAILED: Expected "+expected+"\" got %v", result)
	} else {
		t.Log("PASSED")
	}

	result = Emoji("mitochondria", "🔪️")
	expected = "M🔪️I🔪️T🔪️O🔪️C🔪️H🔪️O🔪️N🔪️D🔪️R🔪️I🔪️A"

	if result != expected {
		t.Errorf("FAILED: Expected "+expected+"\" got %v", result)
	} else {
		t.Log("PASSED")
	}
}
