package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}

func ExampleAdd() {
	sum := Add(5, 1)
	fmt.Println(sum)
	// Output: 6
}
