// Credits to Michael Muinos on Youtube for helping me understand how to solve this problem:
// https://www.youtube.com/watch?v=4i6-9IzQHwo
package main

import "fmt"

func fnmax(p1 int, p2 int) int {
	if p1 > p2 {
		return p1
	}
	return p2
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	i, j, max := 0, 0, 0
	// map[rune]bool is the closest we can get to a hashset of runes
	set := make(map[rune]bool)
	// Loop through s using i
	for i < len(s) {
		// c is the value at i in s
		c := s[i]
		// if c has already been seen in the array, recursively delete everything in the set and move j upto where there are no duplicates
		for set[rune(c)] {
			delete(set, rune(s[j]))
			j++
		}
		// add c to the set
		set[rune(c)] = true
		// assign max the value of either itself or (i-j+1) depending on which is highest
		max = fnmax(max, i-j+1)
		// increment
		i++
	}
	return max
}

func main() {
	testr := "bc789r0q27839qpur8"
	fmt.Println(lengthOfLongestSubstring(testr))
}
