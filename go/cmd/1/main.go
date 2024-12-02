package main

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"os"
	"slices"
	"strconv"
	"strings"
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
	scanner := bufio.NewScanner(r)
	fst = LocationList{}
	snd = LocationList{}

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)
		if len(words) != 2 {
			return nil, nil, fmt.Errorf("invalid line: %q (got %d words)", line, len(words))
		}

		x, err := parseLocationID(words[0])
		if err != nil {
			return nil, nil, err
		}
		fst = append(fst, x)

		y, err := parseLocationID(words[1])
		if err != nil {
			return nil, nil, err
		}
		snd = append(snd, y)
	}

	return fst, snd, scanner.Err()
}

func parseLocationID(word string) (LocationID, error) {
	x, err := strconv.Atoi(word)
	if err != nil {
		return 0, fmt.Errorf("invalid location ID: %q (%w)", word, err)
	}

	return LocationID(x), nil
}

func PartOne(fst, snd LocationList) int {
	return sum(sortedDistances(fst, snd))
}

func PartTwo(fst, snd LocationList) int {
	return sum(similarityScores(fst, snd))
}

func sum(seq iter.Seq[int]) int {
	total := 0
	for x := range seq {
		total += x
	}
	return total
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
