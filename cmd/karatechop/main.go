package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/patrislav/go-katas/karatechop"
)

type result struct {
	name     string
	index    int
	duration time.Duration
}

var (
	haystackFile = flag.String("f", "", "haystack file to use; it must be a list of numbers separated by newlines")
)

func main() {
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage: %s [OPTION]... NEEDLE\n", os.Args[0])
		fmt.Fprintln(out, "Find the index of a number (NEEDLE) in a given list of numbers (haystack)")
		fmt.Fprintln(out)
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		handleError(errors.New("requires exactly 1 argument"))
	}
	if haystackFile == nil || *haystackFile == "" {
		handleError(errors.New("haystack file is required"))
	}
	needle, err := strconv.ParseInt(flag.Arg(0), 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Argument (NEEDLE) must be a number: %v\n", err)
		os.Exit(1)
	}

	haystack := readHaystack()

	choppers := map[string]karatechop.Chopper{
		"stdlib":    karatechop.StdlibChopper{},
		"recursive": karatechop.RecursiveChopper{},
		"iterative": karatechop.IterativeChopper{},
	}

	executeChoppers(choppers, int(needle), haystack)
}

func handleError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n\n", err)
	flag.Usage()
	os.Exit(1)
}

func executeChoppers(choppers map[string]karatechop.Chopper, needle int, haystack []int) {
	start := time.Now()
	results := make(chan result, len(choppers))
	for name, chopper := range choppers {
		name, chopper := name, chopper
		go func() {
			index := chopper.Chop(needle, haystack)
			res := result{name: name, index: index, duration: time.Now().Sub(start)}
			results <- res
		}()
	}

	w := tabwriter.NewWriter(os.Stdout, 10, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintln(w, "CHOPPER\tLINE NO.\tTIME")
	for range choppers {
		res := <-results
		d := float64(res.duration) / float64(time.Millisecond)
		if res.index >= 1 {
			fmt.Fprintf(w, "%s\t%d\t%f ms\n", res.name, res.index+1, d)
		} else {
			fmt.Fprintf(w, fmt.Sprintf("%s\t(n/a)\t%f ms\n", res.name, d))
		}
	}
	w.Flush()
}

func readHaystack() []int {
	r := haystackReadCloser()
	defer r.Close()

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	haystack := make([]int, 0)

	var line int
	for sc.Scan() {
		line++
		n, err := strconv.ParseInt(sc.Text(), 10, 64)
		if err != nil {
			handleError(fmt.Errorf("line %d: %w", line, err))
		}
		haystack = append(haystack, int(n))
	}
	if err := sc.Err(); err != nil {
		handleError(err)
	}
	return haystack
}

func haystackReadCloser() io.ReadCloser {
	switch {
	case strings.HasPrefix(*haystackFile, "http://") || strings.HasPrefix(*haystackFile, "https://"):
		res, err := http.Get(*haystackFile)
		if err != nil {
			handleError(err)
		}
		return res.Body
	case *haystackFile == "-":
		return os.Stdin
	default:
		f, err := os.Open(*haystackFile)
		if err != nil {
			handleError(err)
		}
		return f
	}
}
