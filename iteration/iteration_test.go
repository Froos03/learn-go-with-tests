package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 5)
	expected := "aaaaa"

	assertRepeat(t, actual, expected)
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func assertRepeat(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func ExampleRepeat() {
	repeated := Repeat("b", 2)
	fmt.Println(repeated)
	// Output: bb
}
