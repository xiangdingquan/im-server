package util

import (
	"sort"
)

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type Int32Slice []int32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Int32Intersect(nums1 Int32Slice, nums2 Int32Slice) []int32 {
	if nums1 == nil || nums2 == nil {
		return []int32{}
	}
	sort.Sort(nums1)
	sort.Sort(nums2)
	var x = 0
	var y = 0
	var result Int32Slice
	for {
		if x < len(nums1) && y < len(nums2) {
			if nums1[x] == nums2[y] {
				result = append(result, nums1[x])
				x++
				y++
			} else if nums1[x] > nums2[y] {
				y++
			} else {
				x++
			}
		} else {
			break
		}

	}
	return result
}

func Int32Subtract(nums1 Int32Slice, nums2 Int32Slice) []int32 {
	if nums2 == nil {
		return nums1
	}
	sort.Sort(nums1)
	sort.Sort(nums2)
	var x = 0
	var y = 0
	var result Int32Slice
	for x < len(nums1) && y < len(nums2) {
		if nums1[x] == nums2[y] {
			x++
			y++
		} else {
			if nums1[x] < nums2[y] {
				result = append(result, nums1[x])
				x++
			} else {
				y++
			}
		}
	}
	if x < len(nums1) {
		result = append(result, nums1[x:]...)
	}
	return result
}
