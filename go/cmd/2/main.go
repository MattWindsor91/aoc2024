package main

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"

	aocutil "github.com/MattWindsor91/aoc2024"
)

type Level int
type Report []Level
type Input []Report

func main() {
	in, err := readInput(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("part 1:", in.NumSafe())
}

func readInput(r io.Reader) (Input, error) {
	in := Input{}

	for ints, err := range aocutil.ReadIntMatrix(r) {
		if err != nil {
			return nil, err
		}

		report := slices.Collect(aocutil.Transform(slices.Values(ints), func(x int) Level {
			return Level(x)
		}))

		in = append(in, report)
	}

	return in, nil
}

// IsSafe determines if a report is safe (is monotonic and in the 1-3 step bound)
func (r Report) IsSafe() bool {
	minstep, maxstep := 0, 0

	// which direction are we expecting the report to go in?
	switch cmp.Compare(r[0], r[1]) {
	case -1: // upwards
		minstep, maxstep = 1, 3
	case 1: // downwards
		minstep, maxstep = -3, -1
	case 0: // this report can't be safe
		return false
	}

	// pairwise iteration
	for j := 1; j < len(r); j++ {
		i := j - 1
		diff := int(r[j] - r[i])

		if unsafe := diff < minstep || maxstep < diff; unsafe {
			return false
		}
	}

	return true
}

// NumSafe counts the number of safe reports in an input.
func (i Input) NumSafe() int {
	safeties := aocutil.Transform(slices.Values(i), func(x Report) int {
		if x.IsSafe() {
			return 1
		} else {
			return 0
		}
	})
	return aocutil.Sum(safeties)
}
