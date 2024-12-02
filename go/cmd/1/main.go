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

func distances(fst, snd iter.Seq[LocationID]) iter.Seq[int] {
	return func(yield func(int) bool) {
		nextFst, stopFst := iter.Pull(fst)
		nextSnd, stopSnd := iter.Pull(snd)

		defer stopFst()
		defer stopSnd()

		looping := true
		for looping {
			x, ok := nextFst()
			if !ok {
				return
			}

			y, ok := nextSnd()
			if !ok {
				return
			}

			if x < y {
				looping = yield(int(y - x))
			} else {
				looping = yield(int(x - y))
			}
		}
	}
}

func sortedDistances(fst LocationList, snd LocationList) iter.Seq[int] {
	slices.Sort(fst)
	slices.Sort(snd)

	fstSeq := slices.Values(fst)
	sndSeq := slices.Values(snd)

	return distances(fstSeq, sndSeq)
}

func PartOne(fst, snd LocationList) int {
	total := 0
	for dist := range sortedDistances(fst, snd) {
		total += dist
	}

	return total

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

func main() {
	fst, snd, err := ReadLists(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	one := PartOne(fst, snd)
	fmt.Println("Part 1:", one)
}
