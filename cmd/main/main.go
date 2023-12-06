package main

import (
	"github.com/jackinthebox52/alg-practice/internal/algorithms"
	"github.com/jackinthebox52/alg-practice/internal/benchmark"
)

func main() {
	benchmark.CompareSearchFunctions([]benchmark.Fn{algorithms.LinearSearch, algorithms.BinarySearch, algorithms.InterpolationSearch}, 10000)
}
