package karatechop

import (
	"fmt"
	"testing"
)

var tests = []struct {
	expected int
	needle   int
	haystack []int
}{
	{-1, 3, []int{}},
	{-1, 3, []int{1}},
	{0, 1, []int{1}},
	{0, 1, []int{1, 3, 5}},
	{1, 3, []int{1, 3, 5}},
	{2, 5, []int{1, 3, 5}},
	{-1, 0, []int{1, 3, 5}},
	{-1, 2, []int{1, 3, 5}},
	{-1, 4, []int{1, 3, 5}},
	{-1, 6, []int{1, 3, 5}},
	{0, 1, []int{1, 3, 5, 7}},
	{1, 3, []int{1, 3, 5, 7}},
	{2, 5, []int{1, 3, 5, 7}},
	{3, 7, []int{1, 3, 5, 7}},
	{-1, 0, []int{1, 3, 5, 7}},
	{-1, 2, []int{1, 3, 5, 7}},
	{-1, 4, []int{1, 3, 5, 7}},
	{-1, 6, []int{1, 3, 5, 7}},
	{-1, 8, []int{1, 3, 5, 7}},
}

func testChopper(chopper Chopper, t *testing.T) {
	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := chopper.Chop(tc.needle, tc.haystack)
			if got != tc.expected {
				t.Errorf("got = %d, want = %d", got, tc.expected)
			}
		})
	}
}

func benchmarkChopper(chopper Chopper, b *testing.B) {
	haystack := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
	for n := 0; n < b.N; n++ {
		needle := haystack[n%(len(haystack)-1)]
		_ = chopper.Chop(needle, haystack)
	}
}
