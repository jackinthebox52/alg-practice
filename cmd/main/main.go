package main

import (
	"github.com/jackinthebox52/alg-practice/internal/benchmark"
	"github.com/jackinthebox52/alg-practice/internal/algorithms"
)

func main() {
	benchmark.CompareSearchFunctions(algorithms.LinearSearch, algorithms.BinarySearch, 10000)
	//benchmark.BenchmarkSearchFunction(algorithms.LinearSearch)
	//benchmark.BenchmarkSearchFunction(algorithms.BinarySearch)
}