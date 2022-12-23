package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	repeated := Iteration("d")
	expected := "dddddd"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a")
	}
}
func ExampleIteration() {
	repeated := Iteration("a")
	fmt.Println(repeated)
}
