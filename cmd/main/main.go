package main

import (
	"github.com/jackinthebox52/alg-practice/internal/algorithms"
	"github.com/jackinthebox52/alg-practice/internal/benchmark"
)

func main() {
	//benchmark.CompareSearchFunctions([]benchmark.SearchFn_T{algorithms.LinearSearch, algorithms.BinarySearch, algorithms.InterpolationSearch}, 10000)
	benchmark.CompareSortFunctions([]benchmark.SortFn_T{algorithms.BubbleSort, algorithms.InsertionSort, algorithms.MergeSort, algorithms.TimSort, algorithms.QuickSort}, 5000, 1000, 5000)
}
