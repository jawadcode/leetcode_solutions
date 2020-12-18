package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	res := make([]int, totalLen)
	aCounter, bCounter := 0, 0
	for counter := 0; counter < totalLen && (aCounter < len(nums1) || bCounter < len(nums2)); counter++ {
		if aCounter == len(nums1) {
			for i, v := range nums2[bCounter:] {
				res[counter+i] = v
			}
			break
		} else if bCounter == len(nums2) {
			for i, v := range nums1[aCounter:] {
				res[counter+i] = v
			}
			break
		}
		if nums1[aCounter] > nums2[bCounter] {
			res[counter] = nums2[bCounter]
			bCounter++
		} else {
			res[counter] = nums1[aCounter]
			aCounter++
		}
	}
	if totalLen % 2 == 0 {
		return float64(res[totalLen/2-1]+res[totalLen/2]) / 2
	}
	return float64(res[((totalLen+1)/2)-1])
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
