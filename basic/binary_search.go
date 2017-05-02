package basic

import "errors"

var ErrNotFound = errors.New("not found")

func binarySearch(haystack []int, needle int) (int, error) {

	l := 0
	r := len(haystack) - 1

	for l <= r {
		// Find the midpoint
		m := (r + l) / 2

		if haystack[m] == needle {
			return m, nil
		}

		if haystack[m] > needle {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1, ErrNotFound
}
