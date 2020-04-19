# go-katas

Solutions to katas from http://codekata.com/ implemented in Go.

## Building

```bash
make all
```

This command will run the tests and then build all the binaries.

## [Karate Chop](http://codekata.com/kata/kata02-karate-chop/)

The `karatechop` program implements binary search to find a position of an element in a sorted array
of values. It requires a file containing a sorted list of integers. This repository includes an
example list in the file `cmd/karatechop/numlist.txt`.

To execute the search run the program with:

```bash
bin/karatechop -f PATH-TO-LIST 47
```

This will output the position of the number 47 in the provided list.

The path can also be a URL or `-` in which case the list will be read from the standard input, e.g.:

```bash
cat PATH-TO-LIST | bin/karatechop -f - 47
```

More information about the implementation can be found in the Go documentation of the
`github.com/patrislav/go-katas/karatechop` package.

### List generation

A list matching the requirements can also be generated using a helper program:

```bash
bin/numgen 1 10 5000 > numlist.txt
```

This command will generate a sorted list of numbers from 0 to 5000, where each successive number is
bigger than the previous by a random value between 1 and 10.

## [Word Chains](http://codekata.com/kata/kata19-word-chains/)

Given a dictionary and two words of the same length, this program will find a chain of words from one
to the other. Successive entries must all be real words from the dictionary and can differ from the
previous word by just one letter.

To execute and print a chain from origin to target:

```bash
bin/wordchain -d wordchain/testdata/wordlist.txt origin target
```

The path can also be a URL or `-` in which case the dictionary will be read from the standard input.

For example, to use the word list from codekata.com:

```bash
bin/wordchain -d http://codekata.com/data/wordlist.txt cat dog
```

More information about the implementation can be found in the Go documentation of the
`github.com/patrislav/go-katas/wordchain` package.
