package karatechop

// RecursiveChopper implements the Chopper interface using recursive approach.
type RecursiveChopper struct{}

// Recursive is a default RecursiveChopper, instantiated and exported for ease of use.
var Recursive = RecursiveChopper{}

// Chop returns the index of an item (needle) in a slice (haystack) or -1 if it was not found.
//
// This implementation calls itself recursively, slicing the haystack in half each time.
func (c RecursiveChopper) Chop(needle int, haystack []int) int {
	if len(haystack) == 0 {
		return -1
	}

	midIndex := len(haystack) / 2
	midValue := haystack[midIndex]

	switch {
	case needle < midValue:
		if index := c.Chop(needle, haystack[:midIndex]); index >= 0 {
			return index
		}
	case needle > midValue:
		if index := c.Chop(needle, haystack[midIndex+1:]); index >= 0 {
			return index + midIndex + 1
		}
	default:
		return midIndex
	}

	return -1
}
