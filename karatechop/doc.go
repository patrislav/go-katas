// Package karatechop provides various implementations of the binary search algorithm on integers.
//
// This package provides a few different implementations of the algorithm. All of them implement the
// Chop method to find an index of an int in a slice of ints, e.g.:
//
//     var index int
//     index = karatechop.Iterative.Chop(5, []int{1, 2, 5, 8}) // 2
//     index = karatechop.Recursive.Chop(2, []int{1, 2, 5, 8}) // 1
//     index = karatechop.Stdlib.Chop(8, []int{1, 2, 5, 8}) // 3
package karatechop
