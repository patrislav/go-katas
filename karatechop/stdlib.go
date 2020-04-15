package karatechop

import (
	"sort"
)

// StdlibChopper implements the Chopper interface using the standard library's sort.SearchInts function.
//
// It is used as a base case to test all other implementations against.
type StdlibChopper struct{}

// Stdlib is a default StdlibChopper, instantiated and exported for ease of use.
var Stdlib = StdlibChopper{}

// Chop returns the index of an item (needle) in a slice (haystack) or -1 if it was not found.
func (StdlibChopper) Chop(needle int, haystack []int) int {
	index := sort.SearchInts(haystack, needle)

	// When sort.SearchInts cannot find a given item in the slice, it returns the index at which the
	// item could be placed to maintain sort order. The Chop function, however, should return -1 in
	// such case
	if index == len(haystack) || haystack[index] != needle {
		return -1
	}
	return index
}
