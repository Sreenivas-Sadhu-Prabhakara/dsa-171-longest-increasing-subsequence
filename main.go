package main

import (
	"fmt"
	"sort"
)

// LIS returns the length of the longest strictly increasing subsequence.
// tails[k] holds the smallest possible tail of an increasing subsequence of
// length k+1; each element either extends the longest or improves a tail.
func LIS(nums []int) int {
	tails := []int{}
	for _, v := range nums {
		i := sort.SearchInts(tails, v) // first tail >= v
		if i == len(tails) {
			tails = append(tails, v) // v extends the longest subsequence
		} else {
			tails[i] = v // v is a better (smaller) tail for length i+1
		}
	}
	return len(tails)
}

func main() {
	fmt.Println(LIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
