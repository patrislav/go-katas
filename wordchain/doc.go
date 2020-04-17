// Package wordchain provides an implementation of an algorithm for solving word-chain puzzles.
//
// The puzzle
//
// The word chain puzzle is a challenge to build a chain of words, starting with one particular
// word (called the origin here) and ending with another (called the target.) Successive entries
// in the chain must all be real words coming from the provided dictionary, and each can differ
// from the previous by just one letter.
//
// For example, you can get from "cat" to dog using the following chain:
//
//      cat -> cot -> cog -> dog
//
// Dictionary
//
// A Dictionary holds a list of words, as well as an alphabet consisting of letters from all
// the words. It is used in the Build function to find the shortest chain, since as it's
// mentioned above, all words must be real.
//
// It can be created from any io.Reader whose content must be a whitespace-separated list of
// words. For example to read dictionary from a file:
//
//      f, _ := os.Open("wordlist.txt")
//      defer f.Close()
//
//      dict := wordchain.NewDictionary(3) // 3 is the length of words in the Dictionary
//      _, _ = dict.ReadFrom(f) // dict implements io.ReaderFrom
//
// A Dictionary will only read words that match the length provided to it in the NewDictionary
// call. This limitation is due to the fact that only words with the same length as the origin
// (and target) are considered.
//
// Building a chain
//
// After obtaining a dictionary it can be used to build a chain from the origin to target using
// the Build function:
//
//      chain, err := wordchain.Build(dict, "origin", "target")
//
package wordchain
