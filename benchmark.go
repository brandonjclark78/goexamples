package main

import (
	"fmt"
	"time"
	"math/rand"
)

// sumUniqueSlow - Inefficient O(n²) approach
func sumUniqueSlow(nums []int) int {
	unique := []int{}

	for _, num := range nums {
		isUnique := true
		for _, u := range unique {
			if num == u {
				isUnique = false
				break
			}
		}
		if isUnique {
			unique = append(unique, num)
		}
	}

	sum := 0
	for _, num := range unique {
		sum += num
	}
	return sum
}

// sumUniqueFast - Efficient O(n) approach
func sumUniqueFast(nums []int) int {
	unique := make(map[int]bool)
	sum := 0

	for _, num := range nums {
		if !unique[num] {
			sum += num
			unique[num] = true
		}
	}
	return sum
}

// Benchmark function
func benchmark(function func([]int) int, nums []int, label string) {
	start := time.Now()
	result := function(nums)
	elapsed := time.Since(start)

	fmt.Printf("%s - Result: %d, Time: %s\n", label, result, elapsed)
}

// Generate large test data
func generateTestData(size int, maxValue int) []int {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = rand.Intn(maxValue)
	}
	return nums
}

func main() {
	// Generate a large dataset
	nums := generateTestData(5000, 1000) // 5000 numbers ranging from 0 to 999

	// Run benchmarks
	benchmark(sumUniqueSlow, nums, "Slow Version (O(n²))")
	benchmark(sumUniqueFast, nums, "Fast Version (O(n))")
}
