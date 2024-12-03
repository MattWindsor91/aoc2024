package aocutil

import (
	"bufio"
	"io"
	"iter"
	"strconv"
	"strings"
)

// Sum sums an integer sequence.
func Sum(seq iter.Seq[int]) int {
	total := 0
	for x := range seq {
		total += x
	}
	return total
}

// Transform maps f over seq.
func Transform[T, U any](seq iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for x := range seq {
			if ok := yield(f(x)); !ok {
				return
			}
		}
	}
}

// ReadIntMatrix reads a sequence of rows of integers from a reader.
func ReadIntMatrix(r io.Reader) iter.Seq2[[]int, error] {
	return func(yield func([]int, error) bool) {
		scanner := bufio.NewScanner(r)

		for scanner.Scan() {
			line := scanner.Text()

			words := strings.Fields(line)
			ints, err := parseIntRow(words)
			if err != nil {
				_ = yield(nil, err)
				return
			}

			if ok := yield(ints, nil); !ok {
				return
			}
		}

		if err := scanner.Err(); err != nil {
			yield(nil, err)
		}
	}
}

func parseIntRow(words []string) ([]int, error) {
	ints := make([]int, len(words))
	var err error

	for i, word := range words {
		if ints[i], err = strconv.Atoi(word); err != nil {
			return nil, err
		}
	}

	return ints, nil
}
