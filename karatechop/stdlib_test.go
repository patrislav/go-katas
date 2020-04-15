package karatechop

import (
	"fmt"
	"testing"
)

func ExampleStdlibChopper() {
	chopper := StdlibChopper{}
	haystack := []int{1, 3, 5, 7, 10, 14, 15, 17, 20, 21, 27, 28, 29, 31, 36}
	fmt.Println(chopper.Chop(29, haystack))
	// Output: 12
}

func TestStdlibChopper_Chop(t *testing.T) {
	testChopper(StdlibChopper{}, t)
}

func BenchmarkStdlibChopper_Chop(b *testing.B) {
	benchmarkChopper(StdlibChopper{}, b)
}
