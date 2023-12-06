package benchmark

import (
	"fmt"
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

type Fn func([]int, int) int

// BenchmarkSearchFunction runs a search function {iter} (1000) times and returns the average time taken as a time.Duration.
// The search function must take a slice of integers and an integer to search for, and return the index of the first instance of the integer, or -1 if not found.
func BenchmarkSearchFunction(target Fn, iter int) BenchmarkResult {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	var total_time time.Duration
	for i := 0; i < iter; i++ {
		length := random.Intn(9901) + 100 // Random length between 100 and 10,000
		list := make([]int, length)
		for i := 0; i < length; i++ {
			list[i] = random.Intn(999_901) + 100 // Random number between 100 and 1,000,000
		}
		ind := random.Intn(length) // Random index between 0 and length of list

		start := time.Now() //TIMER START
		target(list, list[ind])
		total_time += time.Since(start)
	}
	avg_time := total_time / time.Duration(iter)
	func_name := getFuncName(target)
	result := BenchmarkResult{func_name, avg_time}
	return result
}

type BenchmarkResult struct {
	name string
	time time.Duration
}

// CompareSearchFunctions runs multiple search functions {iter} times each and prints the results.
func CompareSearchFunctions(functions []Fn, iter int) {
	fmt.Printf("Comparing %d functions for %d iterations each...\n", len(functions), iter)
	results := make([]BenchmarkResult, len(functions))
	for i, Fn := range functions {
		results[i] = BenchmarkSearchFunction(Fn, iter)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].time < results[j].time
	})

	fmt.Println("Fastest to slowest:")
	for _, result := range results {
		fmt.Printf("%v: %v\n", result.name, result.time)
	}
}
