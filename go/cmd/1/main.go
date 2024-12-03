package main

import (
	"fmt"
	"io"
	"iter"
	"os"
	"slices"

	aocutil "github.com/MattWindsor91/aoc2024"
)

type LocationID int
type LocationList []LocationID

func main() {
	fst, snd, err := ReadLists(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println("Part 1:", PartOne(fst, snd))
	fmt.Println("Part 2:", PartTwo(fst, snd))
}

func ReadLists(r io.Reader) (fst LocationList, snd LocationList, err error) {
	fst = LocationList{}
	snd = LocationList{}

	for ints, err := range aocutil.ReadIntMatrix(r) {
		if err != nil {
			return nil, nil, fmt.Errorf("invalid line: %w", err)
		}
		if len(ints) != 2 {
			return nil, nil, fmt.Errorf("invalid line: %v (got %d words)", ints, len(ints))
		}

		fst = append(fst, LocationID(ints[0]))
		snd = append(snd, LocationID(ints[1]))
	}

	return fst, snd, nil
}

func PartOne(fst, snd LocationList) int {
	return aocutil.Sum(sortedDistances(fst, snd))
}

func PartTwo(fst, snd LocationList) int {
	return aocutil.Sum(similarityScores(fst, snd))
}

func sortedDistances(fstUnsorted, sndUnsorted LocationList) iter.Seq[int] {
	// Deep copy the slices when sorting them, lest we accidentally mutate values
	fst := slices.Sorted(slices.Values(fstUnsorted))
	snd := slices.Sorted(slices.Values(sndUnsorted))

	return func(yield func(int) bool) {
		for i := 0; i < len(fst) && i < len(snd); i++ {
			x := fst[i]
			y := snd[i]

			dist := 0
			if x < y {
				dist = int(y - x)
			} else {
				dist = int(x - y)
			}

			if ok := yield(dist); !ok {
				return
			}
		}
	}
}

func similarityScores(fst, snd LocationList) iter.Seq[int] {
	occs := occurrences(fst, snd)

	return func(yield func(int) bool) {
		for _, x := range fst {
			if ok := yield(int(x) * occs[x]); !ok {
				return
			}
		}
	}
}

func occurrences(fst, snd LocationList) map[LocationID]int {
	result := make(map[LocationID]int, len(fst))

	for _, x := range fst {
		result[x] = 0
	}

	for _, x := range snd {
		if occ, ok := result[x]; ok {
			result[x] = occ + 1
		}
	}

	return result
}
