package karatechop

// Chopper is an interface representing an implementation of the binary search algorithm.
//
// Chop finds and returns the index of needle in the haystack. A return value of -1 means that the
// needle value could not be found within the passed haystack slice.
type Chopper interface {
	Chop(needle int, haystack []int) int
}
