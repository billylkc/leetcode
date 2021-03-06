package main

import (
	"fmt"
)

func main() {
	var nums []int

	nums = []int{3, 2, 3, 4, 6, 5}

	fmt.Println(nums)
	fmt.Println(findErrorNums(nums))

}

func findErrorNums_(nums []int) []int {
	var (
		missing   int
		duplicate int
	)
	m := make(map[int]bool)

	// tranverse all list to create hash table
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = true
		} else {
			duplicate = n
		}
	}
	// Find out missing key
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			missing = i
		}
	}

	return []int{duplicate, missing}
}

func findErrorNums(nums []int) []int {
	sum := 0
	numSum := 0
	numsMap := make(map[int]struct{})
	dupVal := -1
	for idx, num := range nums {
		sum += idx + 1
		numSum += num

		if dupVal != -1 {
			continue
		}
		if _, ok := numsMap[num]; ok {
			dupVal = num
		} else {
			numsMap[num] = struct{}{}
		}
	}

	div := sum - numSum

	return []int{dupVal, dupVal + div}
}
