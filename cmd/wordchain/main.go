package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/patrislav/go-katas/wordchain"
)

var (
	dictFile = flag.String("d", "", "dictionary file to use: either a file path, URL starting with \"http(s)://\",\n or \"-\" to read from standard input")
)

func main() {
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage: %s [OPTION]... ORIGIN TARGET\n", os.Args[0])
		fmt.Fprintln(out, "Find a chain of words starting with ORIGIN and ending with TARGET, where successive\nentries must all be real words and each can differ from the previous by just one letter.")
		fmt.Fprintln(out)
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 2 {
		handleError(errors.New("requires 2 arguments"))
	}
	if dictFile == nil || *dictFile == "" {
		handleError(errors.New("dictionary file is required"))
	}

	origin := strings.ToLower(flag.Arg(0))
	target := strings.ToLower(flag.Arg(1))
	if len(origin) != len(target) {
		handleError(errors.New("both words must have the same length"))
	}

	r := dictionaryReadCloser()
	defer r.Close()

	dict := wordchain.NewDictionary(len(origin))
	_, err := dict.ReadFrom(r)
	if err != nil {
		handleError(err)
	}

	chain, err := wordchain.Build(dict, origin, target)
	if err != nil {
		handleError(err)
	}

	fmt.Printf("The shortest chain from %s to %s has %d steps:\n", origin, target, len(chain))
	for i, w := range chain {
		fmt.Printf("\t%d. %s\n", i+1, w)
	}
}

func handleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
	flag.Usage()
	os.Exit(1)
}

func dictionaryReadCloser() io.ReadCloser {
	switch {
	case strings.HasPrefix(*dictFile, "http://") || strings.HasPrefix(*dictFile, "https://"):
		res, err := http.Get(*dictFile)
		if err != nil {
			handleError(err)
		}
		return res.Body
	case *dictFile == "-":
		return os.Stdin
	default:
		f, err := os.Open(*dictFile)
		if err != nil {
			handleError(err)
		}
		return f
	}
}
