package main

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = (num ^ ones) &^ twos
		twos = (num ^ twos) &^ ones
	}
	return ones
}
