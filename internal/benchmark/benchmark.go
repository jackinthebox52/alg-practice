package benchmark

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"
)

// getFuncName returns the name of a function as a string.
func getFuncName(i interface{}) string {
	qualified := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return strings.Split(qualified, ".")[len(strings.Split(qualified, "."))-1]
}

type SearchFn_T func([]int, int) int

type BenchmarkSearchResult struct {
	name string
	time time.Duration
}

// BenchmarkSearchFunction runs a search function {iter} (1000) times and returns the average time taken as a time.Duration.
// The search function must take a slice of integers and an integer to search for, and return the index of the first instance of the integer, or -1 if not found.
func BenchmarkSearchFunction(target SearchFn_T, iter int, maxRangeLength int, maxRangeValue int) BenchmarkSearchResult {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	var total_time time.Duration
	for i := 0; i < iter; i++ {
		length := random.Intn(maxRangeLength-100+1) + 100 // Random length between 100 and maxRangeLength
		list := make([]int, length)
		for i := 0; i < length; i++ {
			list[i] = random.Intn(maxRangeValue-100+1) + 100 // Random number between 100 and maxRangeValue
		}
		ind := random.Intn(length) // Random index between 0 and length of list

		start := time.Now() //TIMER START
		target(list, list[ind])
		total_time += time.Since(start)
	}
	avg_time := total_time / time.Duration(iter)
	func_name := getFuncName(target)
	result := BenchmarkSearchResult{func_name, avg_time}
	return result
}

// CompareSearchFunctions runs multiple search functions {iter} times each and prints the results.
func CompareSearchFunctions(functions []SearchFn_T, iter int, maxRangeLength int, maxRangeValue int) {
	log.Printf("BENCHMARK.go - Benchmarking & comparing %d search functions for %d iterations...\n", len(functions), iter)
	results := make([]BenchmarkSearchResult, len(functions))
	for i, SearchFn_T := range functions {
		results[i] = BenchmarkSearchFunction(SearchFn_T, iter, maxRangeLength, maxRangeValue)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].time < results[j].time
	})

	fmt.Println("Fastest to slowest:")
	for _, result := range results {
		fmt.Printf("%v: %v\n", result.name, result.time)
	}
}

type SortFn_T func([]int) []int

type BenchmarkSortResult struct {
	name string
	time time.Duration
}

// BenchmarkSortFunction runs a sort function {iter} (1000) times and returns the average time taken as a time.Duration.
// The sort function must take a slice of integers and return a sorted slice of integers, or nil if the sort failed (BogoSort).
func BenchmarkSortFunction(target SortFn_T, iter int, maxRangeLength int, maxRangeValue int) BenchmarkSortResult {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	var total_time time.Duration
	for i := 0; i < iter; i++ {
		length := random.Intn(maxRangeLength-100+1) + 100 // Random length between 100 and maxRangeLength
		list := make([]int, length)
		for i := 0; i < length; i++ {
			list[i] = random.Intn(maxRangeValue-100+1) + 100 // Random number between 100 and maxRangeValue
		}

		start := time.Now() //TIMER START
		target(list)
		total_time += time.Since(start)
	}
	avg_time := total_time / time.Duration(iter)
	func_name := getFuncName(target)
	result := BenchmarkSortResult{func_name, avg_time}
	return result
}

// CompareSortFunctions runs multiple sort functions {iter} times each and prints the results.
func CompareSortFunctions(functions []SortFn_T, iter int, maxRangeLength int, maxRangeValue int) {
	start := time.Now()
	log.Printf("BENCHMARK.go - Benchmarking & comparing %d sort functions for %d iterations...\n", len(functions), iter)
	fmt.Printf("The inputs are randomly generated slices of integers with lengths between 100 and %d, and values between 100 and %d.\n", maxRangeLength, maxRangeValue)
	fmt.Println("-------------------------------------------------------------------------------------------------------------------------")
	fmt.Print("Results")
	results := make([]BenchmarkSortResult, len(functions))
	for i, sft := range functions {
		results[i] = BenchmarkSortFunction(sft, iter, maxRangeLength, maxRangeValue)
		fmt.Printf(" -- %v: %v", results[i].name, results[i].time)
	}
	fmt.Println()
	sort.Slice(results, func(i, j int) bool {
		return results[i].time < results[j].time
	})

	fmt.Print("Fastest to slowest")
	for _, result := range results {
		fmt.Printf(" -- %v: %v", result.name, result.time)
	}
	fmt.Println()
	log.Printf("Benchmark took %v\n", time.Since(start))
}
