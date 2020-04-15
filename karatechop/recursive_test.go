package karatechop

import (
	"fmt"
	"testing"
)

func ExampleRecursiveChopper() {
	chopper := RecursiveChopper{}
	haystack := []int{1, 3, 5, 7, 10, 14, 15, 17, 20, 21, 27, 28, 29, 31, 36}
	fmt.Println(chopper.Chop(5, haystack))
	// Output: 2
}

func TestRecursiveChopper_Chop(t *testing.T) {
	testChopper(RecursiveChopper{}, t)
}

func BenchmarkRecursiveChopper_Chop(b *testing.B) {
	benchmarkChopper(RecursiveChopper{}, b)
}
