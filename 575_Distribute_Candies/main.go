package main

import (
	"fmt"
	"sort"
)

func main() {
	// candies := []int{1, 1, 2, 2, 3, 3, 4}
	candies := []int{1, 11}
	fmt.Println(distributeCandies(candies))
}

// Time Complexity: O(N)
// Space Complexity: O(N)
func distributeCandies_(candyType []int) int {
	var cnt int
	max := len(candyType) / 2
	m := make(map[int]bool)
	for _, c := range candyType {
		if _, ok := m[c]; !ok {
			m[c] = true
			cnt += 1
		}
	}
	if cnt > max {
		return max
	} else {
		return cnt
	}
}

// Time Complexity: O(NlogN)
// Space Complexity: O(1)
func distributeCandies(candyType []int) int {
	sort.Ints(candyType)
	var result int
	maxCandy := len(candyType) / 2
	for i := 0; i < len(candyType); i++ {
		if i == 0 {
			result += 1
		}
		if i > 0 && result < maxCandy {
			if candyType[i] != candyType[i-1] {
				result += 1
			}
		}
	}
	return result
}
