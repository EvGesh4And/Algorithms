package main

import (
	"fmt"
	"math"
	"time"
)

func logarithmic(n int) {
	start := time.Now()

	k := 0
	for n > 0 {
		n = n / 2
		k++
	}
	runTime := time.Since(start)
	fmt.Printf("logarithmic: %d Time: %d \n", k, runTime.Nanoseconds())
}

func linear(n int) {
	start := time.Now()

	k := 0
	for i := 0; i < n; i++ {
		k++
	}
	runTime := time.Since(start)
	fmt.Printf("linear: %d Time: %d \n", k, runTime.Nanoseconds())
}

func quadratic(n int) {
	start := time.Now()

	k := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k++
		}
	}
	runTime := time.Since(start)
	fmt.Printf("quadratic: %d Time: %d \n", k, runTime.Nanoseconds())
}

func exponential(n int) {
	start := time.Now()

	k := 0
	for i := 0; i < int(math.Pow(2, float64(n))); i++ {
		k++
	}
	runTime := time.Since(start)
	fmt.Printf("exponential: %d Time: %d \n", k, runTime.Nanoseconds())
}

func main() {
	n := 20
	logarithmic(n)
	linear(n)
	quadratic(n)
	exponential(n)
}
