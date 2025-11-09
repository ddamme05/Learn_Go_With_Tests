package iteration

import "testing"

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	result := Repeat("a", 5)

	if result != expected {
		t.Errorf("got %q want %q", result, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
