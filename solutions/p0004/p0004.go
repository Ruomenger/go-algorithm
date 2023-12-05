package p0101

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen == 0 {
		return 0
	}
	len1 := len(nums1)
	len2 := len(nums2)
	two := totalLen%2 == 0

	left1 := 0
	right1 := len1 - 1
	left2 := 0
	right2 := len2 - 1
	i := left1
	j := left2
	// 0 1   2   3
	// 1 4   8i
	// 2 10j 20  24
	for i+j+1 < totalLen/2 {
		if i < len1 && j < len2 {
			if nums1[i] < nums2[j] {
				left1 = i
				i = (i+right1)/2 + 1
			} else {
				left2 = j
				j = (j+right2)/2 + 1
			}
		} else if i == len1 {
			if nums1[i] < nums2[j] {
				right2 = j
				j = (left2+j)/2 + 1
			} else {
				left2 = j
				j = (j+right2)/2 + 1
			}
		}

	}
	if !two {
		if nums1[i] > nums2[j] {
			return float64(nums2[i])
		}
		return float64(nums1[j])
	}
	return float64(nums1[i]+nums2[j]) / 2.0
}
