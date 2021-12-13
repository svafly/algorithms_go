package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j, tail := m-1, n-1, m+n-1; i >= 0 || j >= 0; tail-- {
		var cur int
		if i == -1 {
			cur = nums2[j]
			j--
		} else if j == -1 {
			cur = nums1[i]
			i--
		} else if nums1[i] < nums2[j] {
			cur = nums2[j]
			j--
		} else {
			cur = nums1[i]
			i--
		}
		nums1[tail] = cur
	}
}

//在go里面没有while语法
