package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s MINSTEP MAXSTEP MAXNUM\n", os.Args[0])
		os.Exit(1)
	}

	minStep := parseNum(args[0])
	maxStep := parseNum(args[1])
	maxNum := parseNum(args[2])

	var last int
	for last+maxStep < maxNum {
		rand.Seed(time.Now().UnixNano())
		step := rand.Intn(maxStep-minStep+1) + minStep
		last += step
		fmt.Println(last)
	}
}

func parseNum(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(num)
}
