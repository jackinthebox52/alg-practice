package algorithms

import (
	"math/rand"
	"time"
)

// MergeSort sorts a slice of integers using the Merge Sort algorithm. It returns a slice of integers.
// It has an average time complexity of O(n log n) and a worst case time complexity of O(n log n).
func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	mid := len(a) / 2
	return merge(MergeSort(a[:mid]), MergeSort(a[mid:]))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result[l+r] = left[l]
			l++
		} else {
			result[l+r] = right[r]
			r++
		}
	}
	for l < len(left) {
		result[l+r] = left[l]
		l++
	}
	for r < len(right) {
		result[l+r] = right[r]
		r++
	}
	return result
}

// TimSort sorts a slice of integers using the Tim Sort algorithm. It returns a slice of integers.
/*
The main idea behind Tim Sort is to exploit the existing order in the data to minimize the number of comparisons and swaps.
It achieves this by dividing the array into small subarrays called runs, which are already sorted,
and then merging these runs using a modified merge sort algorithm.
*/
func TimSort(a []int) []int {
	const minRun = 32

	n := len(a)
	if n < 2 {
		return a
	}

	// Create runs
	var runs [][]int
	for i := 0; i < n; i += minRun {
		end := i + minRun
		if end > n {
			end = n
		}
		runs = append(runs, InsertionSort(a[i:end]))
	}

	// Merge runs
	for len(runs) > 1 {
		nextRuns := [][]int{}
		for i := 0; i < len(runs); i += 2 {
			if i+1 < len(runs) {
				nextRuns = append(nextRuns, merge(runs[i], runs[i+1]))
			} else {
				nextRuns = append(nextRuns, runs[i])
			}
		}
		runs = nextRuns
	}

	return runs[0]
}

// QuickSort sorts a slice of integers using the Quick Sort algorithm. It returns a slice of integers.
// It has an average time complexity of O(n log n) and a worst case time complexity of O(n^2).
func QuickSort(a []int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := random.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]
	for i := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	QuickSort(a[:left])
	QuickSort(a[left+1:])
	return a
}

// InsertionSort sorts a slice of integers using the Insertion Sort algorithm. It returns a slice of integers.
// It has an average time complexity of O(n^2) and a worst case time complexity of O(n^2).
func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 && a[j-1] > a[j] {
			a[j-1], a[j] = a[j], a[j-1] // swap
			j--
		}
	}
	return a
}

// BubbleSort sorts a slice of integers using the Bubble Sort algorithm. It returns a slice of integers.
// It has an average time complexity of O(n^2) and a worst case time complexity of O(n^2).
func BubbleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1] // swap
			}
		}
	}
	return a
}

//
