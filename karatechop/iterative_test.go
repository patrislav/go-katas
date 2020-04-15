package karatechop

import (
	"fmt"
	"testing"
)

func ExampleIterativeChopper() {
	chopper := IterativeChopper{}
	haystack := []int{1, 3, 5, 7, 10, 14, 15, 17, 20, 21, 27, 28, 29, 31, 36}
	fmt.Println(chopper.Chop(27, haystack))
	// Output: 10
}

func TestIterativeChopper_Chop(t *testing.T) {
	testChopper(IterativeChopper{}, t)
}

func BenchmarkIterativeChopper_Chop(b *testing.B) {
	benchmarkChopper(RecursiveChopper{}, b)
}
