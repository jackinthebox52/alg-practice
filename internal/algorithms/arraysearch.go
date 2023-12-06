// Package algorithms provides functions for searching and sorting algorithms, and equivalent, inferior implementations to test against.
// This package is not intended for use in production code. Many functions are intentionally inefficient to test against, or already implemented in the standard library.
package algorithms

// LinearSearch searches for an integer in a slice of integers, returning the index of the first instance of the integer, or -1 if not found.
// This is a simple implementation of linear search, which is O(n) time complexity. Useful for benchmarking against more effecient list search algorithms.
func LinearSearch(a []int, x int) int {
	r := -1 // not found
	for i, v := range a {
		if v == x {
			r = i // found
			break
		}
	}
	return r
}

// BinarySearch searches for an integer in a slice of integers, returning the index of the first instance of the integer, or -1 if not found.
// This is a simple implementation of binary search, which is O(log n) time complexity.
func BinarySearch(a []int, x int) int {
	r := -1 // not found
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == x {
			r = mid // found
			break
		} else if a[mid] < x {
			start = mid + 1
		} else if a[mid] > x {
			end = mid - 1
		}
	}
	return r
}

func InterpolationSearch(a []int, x int) int {
	min, max := a[0], a[len(a)-1]
	l, h := 0, len(a)-1
	for {
		if x < min {
			return l
		}

		if x > max {
			return h + 1
		}
		var guess int
		if h == l {
			guess = h
		} else {
			size := h - l
			offset := int(float64(size-1) * (float64(x-min) / float64(max-min)))
			guess = l + offset
		}
		if a[guess] == x {
			for guess > 0 && a[guess-1] == x {
				guess--
			}
			return guess
		}
		if a[guess] > x {
			h = guess - 1
			max = a[h]
		} else {
			l = guess + 1
			min = a[l]
		}
	}
}
