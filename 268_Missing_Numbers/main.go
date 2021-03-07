package main

import (
	"fmt"
)

func main() {

	nums := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(nums)
	fmt.Println(missingNumber(nums))

}

// XOR implementation
func missingNumber_(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		xor = xor ^ i ^ nums[i]
	}
	return xor ^ i
}

// Gauss' Formula, n(n+1)/2
func missingNumber(nums []int) int {
	expected := len(nums) * (len(nums) + 1) / 2
	actual := sum(nums...)
	return expected - actual
}
func sum(input ...int) int {
	var res int
	for i := range input {
		res += input[i]
	}
	return res
}
