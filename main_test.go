package main

import "testing"

func TestLIS(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4}, // 2,3,7,101 or 2,3,7,18
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7}, 1}, // strictly increasing: duplicates don't chain
		{[]int{1, 2, 3}, 3},
		{[]int{3, 2, 1}, 1},
		{[]int{}, 0},
		{[]int{5}, 1},
	}
	for _, c := range cases {
		if got := LIS(c.in); got != c.want {
			t.Errorf("LIS(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
