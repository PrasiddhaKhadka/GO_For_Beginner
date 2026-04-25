package main

import (
	"fmt"
	"time"
)

func main() {

	nameMap := map[string]uint8{"Adam": 28, "Sarah": 25, "John": 35}

	countryList := []string{"USA", "Australia", "Newzealand"}

	for name := range nameMap {
		fmt.Printf("Name: %v\n", name)
	}

	for idx, values := range countryList {
		fmt.Printf("Country Names: %v\n\nCountry Values: %v\n", idx, values)
	}

	// TESTING THE EXECUTION TIME BETWEEN PREALLOCATED AND NOT ALLOCATED SLICE
	n := 10000
	sliceFirst := []int32{}
	sliceSecond := make([]int32, 0, n)

	fmt.Printf("The total time without preallocation: %v\n", timeLoop(sliceFirst, n))
	fmt.Printf("The total time without preallocation: %v\n", timeLoop(sliceSecond, n))

}

func timeLoop(slice []int32, n int) time.Duration {
	var t0 = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
