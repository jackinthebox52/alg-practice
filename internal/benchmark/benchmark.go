package benchmark

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
	"strings"
)

func getFuncName(i interface{}) string {
	qualified := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return strings.Split(qualified, ".")[len(strings.Split(qualified, "."))-1]
}

type fn func([]int, int) int

// BenchmarkSearchFunction runs a search function {iter} (1000) times and returns the average time taken as a time.Duration.
// The search function must take a slice of integers and an integer to search for, and return the index of the first instance of the integer, or -1 if not found.
func BenchmarkSearchFunction(target fn, iter int) time.Duration {
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

		//Test the function {iter} times
		start := time.Now()		//TIMER START
		target(list, list[ind])
		total_time += time.Since(start)
	}
	avg_time := total_time / time.Duration(iter)
	func_name := getFuncName(target)
	fmt.Printf("Average time taken for %v: %v\n", func_name, avg_time)
	return avg_time
}


func CompareSearchFunctions(primary fn, secondary fn, iter int) {
	fmt.Printf("Comparing %v and %v for %d iterations each...\n", getFuncName(primary), getFuncName(secondary), iter)
	primary_time := BenchmarkSearchFunction(primary, iter)
	secondary_time := BenchmarkSearchFunction(secondary, iter)
	if primary_time > secondary_time {
		s := primary_time
		primary_time = secondary_time
		secondary_time = s
	}
	fmt.Printf("%v is faster than %v by %v\n", getFuncName(primary), getFuncName(secondary), secondary_time - primary_time)
}
